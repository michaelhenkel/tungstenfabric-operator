package config

import (
	"context"
	"reflect"
	"strings"

	configv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	cassandrav1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/cassandra/v1alpha1"
	zookeeperv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/zookeeper/v1alpha1"
	rabbitmqv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/rabbitmq/v1alpha1"

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

var log = logf.Log.WithName("controller_config")

// Add creates a new Config Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileConfig{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("config-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Config
	err = c.Watch(&source.Kind{Type: &configv1alpha1.Config{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Config
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &configv1alpha1.Config{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileConfig{}

// ReconcileConfig reconciles a Config object
type ReconcileConfig struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Config object and makes changes based on the state read
// and what is in the Config.Spec
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileConfig) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Config")

	instance := &configv1alpha1.Config{}
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

	var configMap map[string]string
	configMap = make(map[string]string)
        // Fetch the Cassandra instance
        cassandraInstance := &cassandrav1alpha1.Cassandra{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: "cassandra", Namespace: instance.Namespace}, cassandraInstance)
        //err := r.client.Get(context.TODO(), request.NamespacedName, cassandraInstance)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("cassandra instance not found")
                return reconcile.Result{Requeue: true}, nil
        }
        reqLogger.Info("cassandra instance")
        cassandraConfigMap := &corev1.ConfigMap{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfcassandracmv1", Namespace: cassandraInstance.Namespace}, cassandraConfigMap)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("cassandra configmap not found")
                return reconcile.Result{Requeue: true}, nil
        }

	configMap["CONFIGDB_PORT"] = cassandraConfigMap.Data["CASSANDRA_PORT"]
	configMap["CONFIGDB_CQL_PORT"] = cassandraConfigMap.Data["CASSANDRA_CQL_PORT"]
	configMap["CONFIGDB_NODES"] = cassandraConfigMap.Data["CASSANDRA_SEEDS"]

	// Fetch the Zookeeper instance
        zookeeperInstance := &zookeeperv1alpha1.Zookeeper{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: "zookeeper", Namespace: instance.Namespace}, zookeeperInstance)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("zookeeper instance not found")
                return reconcile.Result{Requeue: true}, nil
        }
        reqLogger.Info("zookeeper instance")
        zookeeperConfigMap := &corev1.ConfigMap{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfzookeepercmv1", Namespace: zookeeperInstance.Namespace}, zookeeperConfigMap)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("zookeeper configmap not found")
                return reconcile.Result{Requeue: true}, nil
        }
	configMap["ZOOKEEPER_NODES"] = zookeeperConfigMap.Data["ZOOKEEPER_NODES"]
	configMap["ZOOKEEPER_NODE_PORT"] = zookeeperConfigMap.Data["ZOOKEEPER_PORT"]

	// Fetch the Rabbitmq instance
        rabbitmqInstance := &rabbitmqv1alpha1.Rabbitmq{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: "rabbitmq", Namespace: instance.Namespace}, rabbitmqInstance)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("rabbitmq instance not found")
                return reconcile.Result{Requeue: true}, nil
        }
        reqLogger.Info("rabbitmq instance")
        rabbitmqConfigMap := &corev1.ConfigMap{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfrabbitmqcmv1", Namespace: rabbitmqInstance.Namespace}, rabbitmqConfigMap)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("rabbitmq configmap not found")
                return reconcile.Result{Requeue: true}, nil
        }
	configMap["RABBITMQ_NODES"] = rabbitmqConfigMap.Data["RABBITMQ_NODES"]
	configMap["RABBITMQ_NODE_PORT"] = rabbitmqConfigMap.Data["RABBITMQ_NODE_PORT"]

	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
                // Define a new deployment
                dep := r.deploymentForConfig(instance)
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

	// Update the Config status with the pod names
        // List the pods for this config's deployment
        podList := &corev1.PodList{}
        labelSelector := labels.SelectorFromSet(labelsForConfig(instance.Name))

        listOps := &client.ListOptions{
                Namespace:     instance.Namespace,
                LabelSelector: labelSelector,
        }
        err = r.client.List(context.TODO(), listOps, podList)
        if err != nil {
                reqLogger.Error(err, "Failed to list pods.", "Config.Namespace", instance.Namespace, "Config.Name", instance.Name)
                return reconcile.Result{}, err
        }

	podNames := getPodNames(podList.Items)
	// Update status.Nodes if needed
        if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
                instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
                if err != nil {
                        reqLogger.Error(err, "Failed to update Config status.")
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
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfconfigcmv1", Namespace: foundConfigmap.Namespace}, foundConfigmap)
		if err != nil && errors.IsNotFound(err) {
			cm := r.configmapForConfig(instance, podIpList, configMap)
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
func newPodForCR(cr *configv1alpha1.Config) *corev1.Pod {
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

func (r *ReconcileConfig) configmapForConfig(m *configv1alpha1.Config, podIpList []string, configMap map[string]string) *corev1.ConfigMap {
	nodeListString := strings.Join(podIpList,",")
	configDbNodes := configMap["CONFIGDB_NODES"]
	configDbPort := configMap["CONFIGDB_PORT"]
	configDbCqlPort := configMap["CONFIGDB_CQL_PORT"]
	zookeeperNodes := configMap["ZOOKEEPER_NODES"] 
	zookeeperPort := configMap["ZOOKEEPER_NODE_PORT"] 
	rabbitmqNodes := configMap["RABBITMQ_NODES"] 
	rabbitmqPort := configMap["RABBITMQ_NODE_PORT"] 

	newConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfconfigcmv1",
			Namespace: m.Namespace,
		},
		Data: map[string]string{
			"CONTROLLER_NODES": nodeListString,
			"CONFIGDB_NODES": configDbNodes,
			"CONFIGDB_PORT": configDbPort,
			"CONFIGDB_CQL_PORT": configDbCqlPort,
			"NODE_TYPE": "config",
			"DOCKER_HOST": "unix://mnt/docker.sock",
			"CONTRAIL_STATUS_IMAGE": m.Spec.StatusImage,
			"ZOOKEEPER_NODES": zookeeperNodes,
			"ZOOKEEPER_NODE_PORT": zookeeperPort,
			"RABBITMQ_NODES": rabbitmqNodes,
			"RABBITMQ_NODE_PORT": rabbitmqPort,
			//"COLLECTOR_NODES":,
		},
	}
        controllerutil.SetControllerReference(m, newConfigMap, r.scheme)
        return newConfigMap
}
// deploymentForConfig returns a config Deployment object
func (r *ReconcileConfig) deploymentForConfig(m *configv1alpha1.Config) *appsv1.Deployment {
        ls := labelsForConfig(m.Name)
        replicas := m.Spec.Size
	apiImage := m.Spec.ApiImage
	deviceManagerImage := m.Spec.DeviceManagerImage
	schemaTransformerImage := m.Spec.SchemaTransformerImage
	serviceMonitorImage := m.Spec.ServiceMonitorImage
	nodeManagerImage := m.Spec.NodeManagerImage
	nodeInitImage := m.Spec.NodeInitImage
	pullPolicy := corev1.PullAlways
	if m.Spec.ImagePullPolicy == "Never" {
		pullPolicy = corev1.PullNever
	}
	if m.Spec.ImagePullPolicy == "IfNotPresent" {
		pullPolicy = corev1.PullNever
	}

	privileged := true
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
						}},
					},{
                                                Image:   nodeInitImage,
                                                Name:    "node-init",
						SecurityContext: &corev1.SecurityContext{
							Privileged: &privileged,
						},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "host-usr-bin",
							MountPath: "/host/usr/bin",
						}},
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfconfigcmv1",
								},
							},
						}},
                                        }},
                                        Containers: []corev1.Container{{
						Image:   apiImage,
						Name:    "api",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfconfigcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "config-logs",
							MountPath: "/var/log/contrail",
						}},
					},{
						Image:   schemaTransformerImage,
						Name:    "schema",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
								Name: "tfconfigcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "config-logs",
							MountPath: "/var/log/contrail",
						}},
					},{
						Image:   deviceManagerImage,
						Name:    "devicemgr",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfconfigcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "config-logs",
							MountPath: "/var/log/contrail",
						}},
					},{
						Image:   serviceMonitorImage,
						Name:    "svcmonitor",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfconfigcmv1",
								},
							},
						}},
					},{
						Image:   nodeManagerImage,
						Name:    "nodemgr",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfconfigcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "config-logs",
							MountPath: "/var/log/contrail",
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
										corev1.DownwardAPIVolumeFile{
											Path: "pod_labelsx",
											FieldRef: &corev1.ObjectFieldSelector{
												FieldPath: "metadata.labels",
											},
										},
									},
								},
							},
						},
						{
							Name: "config-logs",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/log/contrail/config",
								},
							},
						},
						{
							Name: "docker-unix-socket",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/run",
								},
							},
						},
						{
							Name: "host-usr-bin",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/usr/bin",
								},
							},
						},
					},
				},
                        },
                },
        }
        // Set Config instance as the owner and controller
        controllerutil.SetControllerReference(m, dep, r.scheme)
        return dep
}

// labelsForConfig returns the labels for selecting the resources
// belonging to the given config CR name.
func labelsForConfig(name string) map[string]string {
        return map[string]string{"app": "config", "config_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
        var podNames []string
        for _, pod := range pods {
                podNames = append(podNames, pod.Name)
        }
        return podNames
}
