package zookeeper

import (
	"context"
	"reflect"
	"fmt"
	"strings"

	zookeeperv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/zookeeper/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_zookeeper")

// Add creates a new Zookeeper Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileZookeeper{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("zookeeper-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Zookeeper
	err = c.Watch(&source.Kind{Type: &zookeeperv1alpha1.Zookeeper{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Zookeeper
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &zookeeperv1alpha1.Zookeeper{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileZookeeper{}

// ReconcileZookeeper reconciles a Zookeeper object
type ReconcileZookeeper struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Zookeeper object and makes changes based on the state read
// and what is in the Zookeeper.Spec
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileZookeeper) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Zookeeper")

	// Fetch the Zookeeper instance
	instance := &zookeeperv1alpha1.Zookeeper{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
                // Define a new deployment
                dep := r.deploymentForZookeeper(instance)
                reqLogger.Info("Creating a new Deployment.", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
                err = r.client.Create(context.TODO(), dep)
                if err != nil {
                        reqLogger.Error(err, "Failed to create new Deployment.", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
                        return reconcile.Result{}, err
                }
                // Deployment created successfully - return and requeue
                return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
                reqLogger.Error(err, "Failed to get Deployment.")
                return reconcile.Result{}, err
        }
	size := instance.Spec.Size
        if *foundDeployment.Spec.Replicas != size {
                foundDeployment.Spec.Replicas = &size
                err = r.client.Update(context.TODO(), foundDeployment)
                if err != nil {
                        reqLogger.Error(err, "Failed to update Deployment.", "Deployment.Namespace", foundDeployment.Namespace, "Deployment.Name", foundDeployment.Name)
                        return reconcile.Result{}, err
                }
                // Spec updated - return and requeue
                return reconcile.Result{Requeue: true}, nil
        }

	// Update the Zookeeper status with the pod names
        // List the pods for this zookeeper's deployment
        podList := &corev1.PodList{}
        labelSelector := labels.SelectorFromSet(labelsForZookeeper(instance.Name))

        listOps := &client.ListOptions{
                Namespace:     instance.Namespace,
                LabelSelector: labelSelector,
        }
        err = r.client.List(context.TODO(), listOps, podList)
        if err != nil {
                reqLogger.Error(err, "Failed to list pods.", "Zookeeper.Namespace", instance.Namespace, "Zookeeper.Name", instance.Name)
                return reconcile.Result{}, err
        }

	podNames := getPodNames(podList.Items)
	// Update status.Nodes if needed
        if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
                instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
                if err != nil {
                        reqLogger.Error(err, "Failed to update Zookeeper status.")
                        return reconcile.Result{}, err
                }
        }

	// Get state for init PODs
	initContainerRunning := true
	var podIpList []string
	var podNodeNameList []string
	for _, pod := range(podList.Items){
		if pod.Status.PodIP != "" {
			podIpList = append(podIpList, pod.Status.PodIP)
			podNodeNameList = append(podNodeNameList, pod.Spec.NodeName)
			for _, initContainerStatus := range(pod.Status.InitContainerStatuses){
				if initContainerStatus.Name == "init" && initContainerStatus.State.Running == nil {
					initContainerRunning = false
				}
			}
		}
	}
	if int32(len(podIpList)) == size && initContainerRunning {
		foundConfigmap := &corev1.ConfigMap{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfzookeepercmv1", Namespace: foundConfigmap.Namespace}, foundConfigmap)
		if err != nil && errors.IsNotFound(err) {
			// Define a new configmap
			/*
			if instance.Spec.HostNetwork{
				podIpList = podNodeNameList
			}
			*/
			cm := r.configmapForZookeeper(instance, podIpList)
			reqLogger.Info("Creating a new Configmap.", "Configmap.Namespace", cm.Namespace, "Configmap.Name", cm.Name)
			err = r.client.Create(context.TODO(), cm)
			if err != nil {
				reqLogger.Error(err, "Failed to create new Configmap.", "Configmap.Namespace", cm.Namespace, "Configmap.Name", cm.Name)
				return reconcile.Result{}, err
			}
		} else if err != nil {
			reqLogger.Error(err, "Failed to get ConfigMap.")
			return reconcile.Result{}, err
		}
		for _, pod := range(podList.Items){
			foundPod := &corev1.Pod{}
			err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, foundPod)
			if err != nil {
				reqLogger.Error(err, "Failed to update Pod label.")
				return reconcile.Result{}, err
			}
			podMetaData := pod.ObjectMeta
			podLabels := podMetaData.Labels
			podLabels["status"] = "ready"
			foundPod.ObjectMeta.Labels = podLabels
			err = r.client.Update(context.TODO(), foundPod)
			if err != nil {
				reqLogger.Error(err, "Failed to update Pod label.")
				return reconcile.Result{}, err
			}
		}
	} else {
                return reconcile.Result{Requeue: true}, nil
	}
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *zookeeperv1alpha1.Zookeeper) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}

func (r *ReconcileZookeeper) configmapForZookeeper(m *zookeeperv1alpha1.Zookeeper, podIpList []string) *corev1.ConfigMap {
	nodeListString := strings.Join(podIpList,",")

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfzookeepercmv1",
			Namespace: m.Namespace,
		},
		Data: map[string]string{
			"CONTROLLER_NODES": nodeListString,
			"ZOOKEEPER_NODES": nodeListString,
			"ZOOKEEPER_PORT": "2181",
			"ZOOKEEPER_PORTS": "2888:3888",
			"NODE_TYPE": "config-database",
		},
	}
        controllerutil.SetControllerReference(m, configMap, r.scheme)
        return configMap
}
// deploymentForZookeeper returns a zookeeper Deployment object
func (r *ReconcileZookeeper) deploymentForZookeeper(m *zookeeperv1alpha1.Zookeeper) *appsv1.Deployment {
        ls := labelsForZookeeper(m.Name)
        replicas := m.Spec.Size
	zookeeperImage := fmt.Sprintf("%s/contrail-external-zookeeper:%s", m.Spec.Registry, m.Spec.Version)

        dep := &appsv1.Deployment{
                TypeMeta: metav1.TypeMeta{
                        APIVersion: "apps/v1",
                        Kind:       "Deployment",
                },
                ObjectMeta: metav1.ObjectMeta{
                        Name:      m.Name,
                        Namespace: m.Namespace,
                },
                Spec: appsv1.DeploymentSpec{
                        Replicas: &replicas,
                        Selector: &metav1.LabelSelector{
                                MatchLabels: ls,
                        },
                        Template: corev1.PodTemplateSpec{
                                ObjectMeta: metav1.ObjectMeta{
                                        Labels: ls,
                                },
                                Spec: corev1.PodSpec{
					HostNetwork: m.Spec.HostNetwork,
					NodeSelector: map[string]string{
						"node-role.kubernetes.io/master":"",
					},
					Tolerations: []corev1.Toleration{{
						Operator: corev1.TolerationOpExists,
						Effect: corev1.TaintEffectNoSchedule,
					},{
						Operator: corev1.TolerationOpExists,
						Effect: corev1.TaintEffectNoExecute,
					}},
                                        InitContainers: []corev1.Container{{
                                                Image:   "busybox",
                                                Name:    "init",
                                                Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "status",
							MountPath: "/tmp/podinfo",
						},{
							Name: "zookeeper-data",
							MountPath: "/var/lib/zookeeper",
						},{
							Name: "zookeeper-logs",
							MountPath: "/var/log/zookeeper",
						}},
                                        }},
                                        Containers: []corev1.Container{{
                                                Image:   zookeeperImage,
                                                Name:    "zookeeper",
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfzookeepercmv1",
								},
							},
						}},
                                        }},
					Volumes: []corev1.Volume{
						{
							Name: "status",
							VolumeSource: corev1.VolumeSource{
								DownwardAPI: &corev1.DownwardAPIVolumeSource{
									Items: []corev1.DownwardAPIVolumeFile{
										corev1.DownwardAPIVolumeFile{
											Path: "pod_labels",
											FieldRef: &corev1.ObjectFieldSelector{
												FieldPath: "metadata.labels",
											},
										},
									},
								},
							},
						},
						{
							Name: "zookeeper-data",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/lib/contrail/zookeeper",
								},
							},
						},
						{
							Name: "zookeeper-logs",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/log/contrail/zookeeper",
								},
							},
						},
					},
				},
                        },
                },
        }
        // Set Zookeeper instance as the owner and controller
        controllerutil.SetControllerReference(m, dep, r.scheme)
        return dep
}

// labelsForZookeeper returns the labels for selecting the resources
// belonging to the given zookeeper CR name.
func labelsForZookeeper(name string) map[string]string {
        return map[string]string{"app": "zookeeper", "zookeeper_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
        var podNames []string
        for _, pod := range pods {
                podNames = append(podNames, pod.Name)
        }
        return podNames
}
