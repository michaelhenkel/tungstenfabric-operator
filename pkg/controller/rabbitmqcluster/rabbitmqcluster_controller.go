package rabbitmqcluster

import (
	"context"
	"reflect"
<<<<<<< HEAD
	"strings"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
=======
	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
>>>>>>> v0.0.4
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
<<<<<<< HEAD
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
=======
>>>>>>> v0.0.4
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_rabbitmqcluster")

<<<<<<< HEAD
/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new RabbitmqCluster Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
=======
>>>>>>> v0.0.4
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

<<<<<<< HEAD
// newReconciler returns a new reconcile.Reconciler
=======
>>>>>>> v0.0.4
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileRabbitmqCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

<<<<<<< HEAD
// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
=======
func add(mgr manager.Manager, r reconcile.Reconciler) error {
>>>>>>> v0.0.4
	c, err := controller.New("rabbitmqcluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

<<<<<<< HEAD
	// Watch for changes to primary resource RabbitmqCluster
=======
>>>>>>> v0.0.4
	err = c.Watch(&source.Kind{Type: &tfv1alpha1.RabbitmqCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

<<<<<<< HEAD
	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner RabbitmqCluster
=======
>>>>>>> v0.0.4
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.RabbitmqCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileRabbitmqCluster{}

<<<<<<< HEAD
// ReconcileRabbitmqCluster reconciles a RabbitmqCluster object
type ReconcileRabbitmqCluster struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
=======
type ReconcileRabbitmqCluster struct {
>>>>>>> v0.0.4
	client client.Client
	scheme *runtime.Scheme
}

<<<<<<< HEAD
// Reconcile reads that state of the cluster for a RabbitmqCluster object and makes changes based on the state read
// and what is in the RabbitmqCluster.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
=======
>>>>>>> v0.0.4
func (r *ReconcileRabbitmqCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling RabbitmqCluster")

	// Fetch the RabbitmqCluster instance
	instance := &tfv1alpha1.RabbitmqCluster{}
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

<<<<<<< HEAD
	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deploymentForRabbitmqCluster(instance)
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

	// Update the RabbitmqCluster status with the pod names
	// List the pods for this rabbitmq's deployment
	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(labelsForRabbitmqCluster(instance.Name))
	listOps := &client.ListOptions{
		Namespace:     instance.Namespace,
		LabelSelector: labelSelector,
	}
	err = r.client.List(context.TODO(), listOps, podList)
	if err != nil {
		reqLogger.Error(err, "Failed to list pods.", "RabbitmqCluster.Namespace", instance.Namespace, "RabbitmqCluster.Name", instance.Name)
		return reconcile.Result{}, err
	}

	podNames := getPodNames(podList.Items)
	// Update status.Nodes if needed
=======
	baseInstance := &tfv1alpha1.TungstenfabricManager{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, baseInstance)
	if err != nil && errors.IsNotFound(err){
		reqLogger.Info("baseconfig instance not found")
	}

	var configMap = make(map[string]string)
	for k,v := range(baseInstance.Spec.RabbitmqConfig){
		configMap[k] = v
	}

	var resource tfv1alpha1.TungstenFabricResource
	clusterResource := &tfv1alpha1.ClusterResource{
		Name: "rabbitmq",
		InstanceName: instance.Name,
		InstanceNamespace: instance.Namespace,
		Containers: instance.Spec.Containers,
		General: instance.Spec.General,
		ResourceConfig: configMap,
		BaseInstance: baseInstance,
		InitContainers: instance.Spec.InitContainers,
		Type: instance.Spec.Type,
		VolumeList: map[string]bool{},
	}
	resource = clusterResource

	// Create Deployment
	err = resource.CreateDeployment(r.client, instance, r.scheme)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info(clusterResource.Name + " deployment created")

	var podNames []string
	podNames, err = resource.GetPodNames(r.client)
	if err != nil {
		reqLogger.Error(err, "Failed to get PodNames")
		return reconcile.Result{}, err
	} else {
		reqLogger.Info("Got PodNames")
	}

>>>>>>> v0.0.4
	if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
		instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
		if err != nil {
<<<<<<< HEAD
			reqLogger.Error(err, "Failed to update RabbitmqCluster status.")
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
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfrabbitmqcmv1", Namespace: foundConfigmap.Namespace}, foundConfigmap)
		if err != nil && errors.IsNotFound(err) {
			// Define a new configmap
			/*
			if instance.Spec.HostNetwork{
				podIpList = podNodeNameList
			}
			*/
			cm := r.configmapForRabbitmqCluster(instance, podIpList)
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
func newPodForCR(cr *tfv1alpha1.RabbitmqCluster) *corev1.Pod {
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

func (r *ReconcileRabbitmqCluster) configmapForRabbitmqCluster(m *tfv1alpha1.RabbitmqCluster, podIpList []string) *corev1.ConfigMap {
	nodeListString := strings.Join(podIpList,",")

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfrabbitmqcmv1",
			Namespace: m.Namespace,
		},
		Data: map[string]string{
			"CONTROLLER_NODES": nodeListString,
			"RABBITMQ_NODES": nodeListString,
			"RABBITMQ_NODE_PORT": m.Spec.Port,
			"RABBITMQ_ERLANG_COOKIE": m.Spec.Cookie,
			"NODE_TYPE": "config-database",
		},
	}
	controllerutil.SetControllerReference(m, configMap, r.scheme)
	return configMap
}
// deploymentForRabbitmqCluster returns a rabbitmq Deployment object
func (r *ReconcileRabbitmqCluster) deploymentForRabbitmqCluster(m *tfv1alpha1.RabbitmqCluster) *appsv1.Deployment {
	ls := labelsForRabbitmqCluster(m.Name)
	replicas := m.Spec.Size
	rabbitmqImage := m.Spec.Image
	pullPolicy := corev1.PullAlways
	if m.Spec.ImagePullPolicy == "Never" {
		pullPolicy = corev1.PullNever
	}
	if m.Spec.ImagePullPolicy == "IfNotPresent" {
		pullPolicy = corev1.PullNever
	}
	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: m.Name,
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
							}},
						}},
						Containers: []corev1.Container{{
							Image:   rabbitmqImage,
							Name:    "rabbitmq",
							ImagePullPolicy: pullPolicy,
							EnvFrom: []corev1.EnvFromSource{{
								ConfigMapRef: &corev1.ConfigMapEnvSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "tfrabbitmqcmv1",
									},
								},
							}},
							VolumeMounts: []corev1.VolumeMount{{
								Name: "rabbitmq-data",
								MountPath: "/var/lib/rabbitmq",
							},{
								Name: "rabbitmq-logs",
								MountPath: "/var/log/rabbitmq",
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
							Name: "rabbitmq-data",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/lib/contrail/rabbitmq",
								},
							},
						},
						{
							Name: "rabbitmq-logs",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/log/contrail/rabbitmq",
								},
							},
						},
					},
				},
			},
			},
        }
        // Set RabbitmqCluster instance as the owner and controller
        controllerutil.SetControllerReference(m, dep, r.scheme)
        return dep
}

// labelsForRabbitmqCluster returns the labels for selecting the resources
// belonging to the given rabbitmq CR name.
func labelsForRabbitmqCluster(name string) map[string]string {
        return map[string]string{"app": "rabbitmq", "rabbitmq_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
        var podNames []string
        for _, pod := range pods {
                podNames = append(podNames, pod.Name)
        }
        return podNames
}
=======
			reqLogger.Error(err, "Failed to update Pod status.")
			return reconcile.Result{}, err
		}
	}
	reqLogger.Info("Updated Node status with PodNames")

	var initContainerRunning bool
	initContainerRunning, err = resource.WaitForInitContainer(r.client)
	if err != nil || !initContainerRunning{
		reqLogger.Info("Init container not running")
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("Init Container running")

	clusterResource.ResourceConfig["CONTROLLER_NODES"] = resource.GetNodeIpList()
	clusterResource.ResourceConfig["RABBITMQ_NODES"] = resource.GetNodeIpList()

	// Create ConfigMap
	err = resource.CreateConfigMap(r.client, instance, r.scheme)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("Rabbitmq configmap created")

	var labeledPod *corev1.Pod
	for _, pod := range(podNames){
		labeledPod, err = resource.LabelPod(r.client, pod)
		if err != nil {
			return reconcile.Result{}, err
		}
			err = r.client.Update(context.TODO(), labeledPod)
		if err != nil {
			reqLogger.Error(err, "Failed to update Pod label.")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Labeled Pod")
	}

	return reconcile.Result{}, nil
}
>>>>>>> v0.0.4
