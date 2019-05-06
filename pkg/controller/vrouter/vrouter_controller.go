package vrouter

import (
	"context"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

var log = logf.Log.WithName("controller_vrouter")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Vrouter Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileVrouter{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("vrouter-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Vrouter
	err = c.Watch(&source.Kind{Type: &tfv1alpha1.Vrouter{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Vrouter
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.Vrouter{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileVrouter{}

// ReconcileVrouter reconciles a Vrouter object
type ReconcileVrouter struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Vrouter object and makes changes based on the state read
// and what is in the Vrouter.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileVrouter) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Vrouter")

	// Fetch the Vrouter instance
	instance := &tfv1alpha1.Vrouter{}
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

	baseInstance := &tfv1alpha1.TungstenfabricConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, baseInstance)
	if err != nil && errors.IsNotFound(err){
		reqLogger.Info("baseconfig instance not found")
	} else {
		for k,v := range(baseInstance.Spec.General){
			configMap[k] = v
		}
		for k,v := range(baseInstance.Spec.ControlCluster){
			configMap[k] = v
		}
		if instance.Spec.ImagePullPolicy == "" {
			instance.Spec.ImagePullPolicy = baseInstance.Spec.General["imagePullPolicy"]
		}
		if instance.Spec.HostNetwork == ""{
			instance.Spec.HostNetwork  = baseInstance.Spec.General["hostNetwork"]
		}
		if instance.Spec.VrouterNicInit == "" {
			instance.Spec.VrouterNicInit = baseInstance.Spec.Images["vrouterNicInit"]
		}
		if instance.Spec.VrouterKernelInit == "" {
			instance.Spec.VrouterKernelInit = baseInstance.Spec.Images["vrouterKernelInit"]
		}
		if instance.Spec.VrouterCni == "" {
			instance.Spec.VrouterCni = baseInstance.Spec.Images["vrouterCni"]
		}
		if instance.Spec.VrouterAgent == "" {
			instance.Spec.VrouterAgent = baseInstance.Spec.Images["vrouterNodeAgent"]
		}
		if instance.Spec.NodeInitImage == "" {
			instance.Spec.NodeInitImage = baseInstance.Spec.Images["nodeInit"]
		}
		if instance.Spec.StatusImage == "" {
			instance.Spec.StatusImage = baseInstance.Spec.Images["status"]
		}
	}

	reqLogger.Info("control instance")
	controlInstance := &tfv1alpha1.ControlCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, controlInstance)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("control instance not found")
		return reconcile.Result{Requeue: true}, nil
	}
	controlConfigMap := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfcontrolcmv1", Namespace: instance.Namespace}, controlConfigMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("control configmap not found")
		return reconcile.Result{Requeue: true}, nil
	}

	configMap["CONTROL_NODES"] = controlConfigMap.Data["CONTROL_NODES"]

	configInstance := &tfv1alpha1.ConfigCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, configInstance)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("config instance not found")
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("config instance")
	configConfigMap := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfconfigcmv1", Namespace: instance.Namespace}, configConfigMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("config configmap not found")
		return reconcile.Result{Requeue: true}, nil
	}

	configMap["CONFIG_NODES"] = configConfigMap.Data["CONTROLLER_NODES"]

	foundVroutermap := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfvroutercmv1", Namespace: instance.Namespace}, foundVroutermap)
	if err != nil && errors.IsNotFound(err) {
		cm := r.configmapForVrouter(instance, configMap)
		reqLogger.Info("Creating a new VrouterConfigmap.", "VrouterConfigmap.Namespace", cm.Namespace, "VrouterConfigmap.Name", cm.Name)
		err = r.client.Create(context.TODO(), cm)
		if err != nil {
			reqLogger.Error(err, "Failed to create new VrouterConfigmap.", "VrouterConfigmap.Namespace", cm.Namespace, "VrouterConfigmap.Name", cm.Name)
			return reconcile.Result{}, err
		}
	} else if err != nil {
		reqLogger.Error(err, "Failed to get ConfigMap.")
		return reconcile.Result{}, err
	} else {
		cm := r.configmapForVrouter(instance, configMap)
		err = r.client.Update(context.TODO(), cm)
		if err != nil {
			reqLogger.Error(err, "Failed to update VrouterConfigmap.", "VrouterConfigmap.Namespace", cm.Namespace, "VrouterConfigmap.Name", cm.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("Updated VrouterConfigmap.", "VrouterConfigmap.Namespace", cm.Namespace, "VrouterConfigmap.Name", cm.Name)
	}

	foundDaemonset := &appsv1.DaemonSet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "vrouter-" + instance.Name, Namespace: instance.Namespace}, foundDaemonset)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		ds := r.daemonSetForVrouter(instance)
		reqLogger.Info("Creating a new Daemonset.", "Daemonset.Namespace", ds.Namespace, "Daemonset.Name", ds.Name)
		err = r.client.Create(context.TODO(), ds)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Daemonset.", "Daemonset.Namespace", ds.Namespace, "Daemonset.Name", ds.Name)
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Daemonset.")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileVrouter) configmapForVrouter(m *tfv1alpha1.Vrouter, configMap map[string]string) *corev1.ConfigMap {

	configMap["DOCKER_HOST"] = "unix://mnt/docker.sock"
	configMap["CONTRAIL_STATUS_IMAGE"] = m.Spec.StatusImage
	configMap["ANALYTICS_NODES"] = configMap["CONFIG_NODES"]

	newConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfvroutercmv1",
			Namespace: m.Namespace,
		},
		Data: configMap,
	}
	controllerutil.SetControllerReference(m, newConfigMap, r.scheme)
	return newConfigMap
}

func (r *ReconcileVrouter) daemonSetForVrouter(m *tfv1alpha1.Vrouter) *appsv1.DaemonSet {
	ls := labelsForVrouter(m.Name)
	var hostNetworkBool bool
	if m.Spec.HostNetwork == "true" {
		hostNetworkBool = true
	} else {
		hostNetworkBool = false
	}
	pullPolicy := corev1.PullAlways
	if m.Spec.ImagePullPolicy == "Never" {
		pullPolicy = corev1.PullNever
	}
	if m.Spec.ImagePullPolicy == "IfNotPresent" {
		pullPolicy = corev1.PullNever
	}
	privileged := true
	ds := &appsv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "DaemonSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "vrouter-" + m.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
				        Labels: ls,
				},
				Spec: corev1.PodSpec{
				HostNetwork: hostNetworkBool,
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
						Image:   m.Spec.NodeInitImage,
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
									Name: "tfkubemanagercmv1",
								},
							},
						}},
					},{
						Image:   m.Spec.VrouterKernelInit,
						Name:    "contrail-vrouter-kernel-init",
						ImagePullPolicy: pullPolicy,
						SecurityContext: &corev1.SecurityContext{
							Privileged: &privileged,
						},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "host-usr-bin",
							MountPath: "/host/usr/bin",
						},{
							Name: "usr-src",
							MountPath: "/usr/src",
						},{
							Name: "lib-modules",
							MountPath: "/lib/modules",
						},{
							Name: "network-scripts",
							MountPath: "/etc/sysconfig/network-scripts",
						},{
							Name: "host-bin",
							MountPath: "/host/bin",
						}},
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfvroutercmv1",
								},
							},
						}},
					},{
						Image:   m.Spec.VrouterNicInit,
						Name:    "contrail-vrouter-nic-init",
						ImagePullPolicy: pullPolicy,
						SecurityContext: &corev1.SecurityContext{
							Privileged: &privileged,
						},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "host-usr-bin",
							MountPath: "/host/usr/bin",
						},{
							Name: "usr-src",
							MountPath: "/usr/src",
						},{
							Name: "lib-modules",
							MountPath: "/lib/modules",
						},{
							Name: "network-scripts",
							MountPath: "/etc/sysconfig/network-scripts",
						},{
							Name: "host-bin",
							MountPath: "/host/bin",
						}},
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfvroutercmv1",
								},
							},
						}},
					},{
						Image:   m.Spec.VrouterCni,
						Name:    "contrail-kubernetes-cni-init",
						ImagePullPolicy: pullPolicy,
						SecurityContext: &corev1.SecurityContext{
							Privileged: &privileged,
						},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "var-lib-contrail",
							MountPath: "/var/lib/contrail",
						},{
							Name: "etc-cni",
							MountPath: "/host/etc_cni",
						},{
							Name: "opt-cni-bin",
							MountPath: "/host/opt_cni_bin",
						},{
							Name: "var-log-contrail-cni",
							MountPath: "/host/log_cni",
						},{
							Name: "agent-logs",
							MountPath: "/var/log/contrail",
						}},
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfvroutercmv1",
								},
							},
						}},
					},
				},
				Containers: []corev1.Container{{
					Image:   m.Spec.VrouterAgent,
					Name:    "vrouter-agent",
					ImagePullPolicy: pullPolicy,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &privileged,
					},
					EnvFrom: []corev1.EnvFromSource{{
						ConfigMapRef: &corev1.ConfigMapEnvSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "tfkubemanagercmv1",
							},
						},
					}},
					VolumeMounts: []corev1.VolumeMount{{
						Name: "agent-logs",
						MountPath: "/var/log/contrail",
					},{
						Name: "dev",
						MountPath: "/dev",
					},{
						Name: "network-scripts",
						MountPath: "/etc/sysconfig/network-scripts",
					},{
						Name: "host-bin",
						MountPath: "/host/bin",
					},{
						Name: "usr-src",
						MountPath: "/usr/src",
					},{
						Name: "lib-modules",
						MountPath: "/lib/modules",
					},{
						Name: "var-lib-contrail",
						MountPath: "/var/lib/contrail",
					},{
						Name: "var-crashes",
						MountPath: "/var/crashes",
					}},
				},{
					Image:   m.Spec.NodeManagerImage,
					Name:    "control-nodemgr",
					ImagePullPolicy: pullPolicy,
					Env: []corev1.EnvVar{{
						Name: "NODE_TYPE",
						Value: "vrouter",
					}},
					EnvFrom: []corev1.EnvFromSource{{
						ConfigMapRef: &corev1.ConfigMapEnvSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "tfvroutermv1",
							},
						},
					}},
					VolumeMounts: []corev1.VolumeMount{{
						Name: "agent-logs",
						MountPath: "/var/log/contrail",
					},{
						Name: "docker-unix-socket",
						MountPath: "/mnt",
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
						Name: "dev",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/dev",
							},
						},
					},
					{
						Name: "network-scripts",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/etc/sysconfig/network-scripts",
							},
						},
					},
					{
						Name: "host-bin",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/bin",
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
						Name: "usr-src",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/usr/src",
							},
						},
					},
					{
						Name: "lib-modules",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/lib/modules",
							},
						},
					},
					{
						Name: "var-lib-contrail",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/lib/contrail",
							},
						},
					},
					{
						Name: "var-crashes",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/contrail/crashes",
							},
						},
					},
					{
						Name: "etc-cni",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/etc/cni",
							},
						},
					},
					{
						Name: "opt-cni-bin",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/opt/cni/bin",
							},
						},
					},
					{
						Name: "var-log-contrail-cni",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/log/contrail/cni",
							},
						},
					},
					{
						Name: "agent-logs",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/log/contrail/agent",
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
        // Set ControlCluster instance as the owner and controller
        controllerutil.SetControllerReference(m, ds, r.scheme)
        return ds
}

func labelsForVrouter(name string) map[string]string {
        return map[string]string{"app": "vrouter", "vrouter_cr": name}
}
