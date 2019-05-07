package controlcluster

import (
	"context"
	"reflect"
	"strings"
	"strconv"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"

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

var log = logf.Log.WithName("controller_controlcluster")

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileControlCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("controlcluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.ControlCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.ControlCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileControlCluster{}

type ReconcileControlCluster struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileControlCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ControlCluster")

	instance := &tfv1alpha1.ControlCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	var size int32
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
		if instance.Spec.Size == "" {
			if baseInstance.Spec.ControlCluster["size"] != "" {
				instance.Spec.Size = baseInstance.Spec.ControlCluster["size"]
			}
		}
		if instance.Spec.ImagePullPolicy == "" {
			instance.Spec.ImagePullPolicy = baseInstance.Spec.General["imagePullPolicy"]
		}
		if instance.Spec.HostNetwork == ""{
			instance.Spec.HostNetwork  = baseInstance.Spec.General["hostNetwork"]
		}
		if instance.Spec.ControlImage == "" {
			instance.Spec.ControlImage = baseInstance.Spec.Images["control"]
		}
		if instance.Spec.NamedImage == "" {
			instance.Spec.NamedImage = baseInstance.Spec.Images["named"]
		}
		if instance.Spec.DnsImage == "" {
			instance.Spec.DnsImage = baseInstance.Spec.Images["dns"]
		}
		if instance.Spec.NodeManagerImage == "" {
			instance.Spec.NodeManagerImage = baseInstance.Spec.Images["nodeManager"]
		}
		if instance.Spec.NodeInitImage == "" {
			instance.Spec.NodeInitImage = baseInstance.Spec.Images["nodeInit"]
		}
		if instance.Spec.StatusImage == "" {
			instance.Spec.StatusImage = baseInstance.Spec.Images["status"]
		}
	}

	size64, err := strconv.ParseInt(instance.Spec.Size, 10, 64)
	if err != nil {
		return reconcile.Result{}, err
	}
	size = int32(size64)
	cassandraInstance := &tfv1alpha1.CassandraCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, cassandraInstance)

	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("cassandra instance not found")
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("cassandra instance")
	cassandraConfigMap := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfcassandraclustercmv1", Namespace: instance.Namespace}, cassandraConfigMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("cassandra configmap not found")
		return reconcile.Result{Requeue: true}, nil
	}

	configMap["CONFIGDB_PORT"] = cassandraConfigMap.Data["CASSANDRA_PORT"]
	configMap["CONFIGDB_CQL_PORT"] = cassandraConfigMap.Data["CASSANDRA_CQL_PORT"]
	configMap["CONFIGDB_NODES"] = cassandraConfigMap.Data["CASSANDRA_SEEDS"]

	// Fetch the Rabbitmq instance
	rabbitmqInstance := &tfv1alpha1.RabbitmqCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, rabbitmqInstance)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("rabbitmq instance not found")
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("rabbitmq instance")
	rabbitmqConfigMap := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfrabbitmqcmv1", Namespace: instance.Namespace}, rabbitmqConfigMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("rabbitmq configmap not found")
		return reconcile.Result{Requeue: true}, nil
	}
	configMap["RABBITMQ_NODES"] = rabbitmqConfigMap.Data["RABBITMQ_NODES"]
	configMap["RABBITMQ_NODE_PORT"] = rabbitmqConfigMap.Data["RABBITMQ_NODE_PORT"]

	// Fetch the Config instance
	configInstance := &tfv1alpha1.ConfigCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, configInstance)
	//err := r.client.Get(context.TODO(), request.NamespacedName, cassandraInstance)
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

	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "control-" + instance.Name, Namespace: instance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deploymentForControlCluster(instance, size)
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

	if *foundDeployment.Spec.Replicas != size {
		foundDeployment.Spec.Replicas = &size
		err = r.client.Update(context.TODO(), foundDeployment)
		if err != nil {
			reqLogger.Error(err, "Failed to update Deployment.", "Deployment.Namespace", foundDeployment.Namespace, "Deployment.Name", foundDeployment.Name)
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(labelsForControlCluster(instance.Name))

	listOps := &client.ListOptions{
		Namespace:     instance.Namespace,
		LabelSelector: labelSelector,
	}
	err = r.client.List(context.TODO(), listOps, podList)
	if err != nil {
		reqLogger.Error(err, "Failed to list pods.", "ControlCluster.Namespace", instance.Namespace, "ControlCluster.Name", instance.Name)
		return reconcile.Result{}, err
	}

	podNames := getPodNames(podList.Items)
	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
		instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update ControlCluster status.")
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
		foundControlClustermap := &corev1.ConfigMap{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfcontrolcmv1", Namespace: instance.Namespace}, foundControlClustermap)
		if err != nil && errors.IsNotFound(err) {
			cm := r.configmapForControlCluster(instance, podIpList, configMap)
			reqLogger.Info("Creating a new ControlClustermap.", "ControlClustermap.Namespace", cm.Namespace, "ControlClustermap.Name", cm.Name)
			err = r.client.Create(context.TODO(), cm)
			if err != nil {
				reqLogger.Error(err, "Failed to create new ControlClustermap.", "ControlClustermap.Namespace", cm.Namespace, "ControlClustermap.Name", cm.Name)
				return reconcile.Result{}, err
			}
		} else if err != nil {
			reqLogger.Error(err, "Failed to get ConfigMap.")
			return reconcile.Result{}, err
		} else {

			cm := r.configmapForControlCluster(instance, podIpList, configMap)
			err = r.client.Update(context.TODO(), cm)
			if err != nil {
				reqLogger.Error(err, "Failed to update Configmap.", "Configmap.Namespace", cm.Namespace, "Configmap.Name", cm.Name)
				return reconcile.Result{}, err
			}
			reqLogger.Info("Updated ControlClustermap.", "ControlClustermap.Namespace", cm.Namespace, "ControlClustermap.Name", cm.Name)
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

func (r *ReconcileControlCluster) configmapForControlCluster(m *tfv1alpha1.ControlCluster, podIpList []string, configMap map[string]string) *corev1.ConfigMap {
	nodeListString := strings.Join(podIpList,",")

	configMap["CONTROLLER_NODES"] = nodeListString
	configMap["DOCKER_HOST"] = "unix://mnt/docker.sock"
	configMap["CONTRAIL_STATUS_IMAGE"] = m.Spec.StatusImage
	configMap["ANALYTICS_NODES"] = configMap["CONFIG_NODES"]

	newConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfcontrolcmv1",
			Namespace: m.Namespace,
		},
		Data: configMap,
	}
	controllerutil.SetControllerReference(m, newConfigMap, r.scheme)
	return newConfigMap
}
// deploymentForControlCluster returns a config Deployment object
func (r *ReconcileControlCluster) deploymentForControlCluster(m *tfv1alpha1.ControlCluster, size int32) *appsv1.Deployment {
	ls := labelsForControlCluster(m.Name)
	replicas :=size
	controlImage := m.Spec.ControlImage
	dnsImage := m.Spec.DnsImage
	namedImage := m.Spec.NamedImage
	nodeManagerImage := m.Spec.NodeManagerImage
	nodeInitImage := m.Spec.NodeInitImage
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
        dep := &appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "control-" + m.Name,
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
									Name: "tfcontrolcmv1",
								},
							},
						}},
					}},
					Containers: []corev1.Container{{
						Image:   controlImage,
						Name:    "control",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfcontrolcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "control-logs",
							MountPath: "/var/log/contrail",
						}},
					},{
						Image:   namedImage,
						Name:    "named",
						ImagePullPolicy: pullPolicy,
						SecurityContext: &corev1.SecurityContext{
							Privileged: &privileged,
						},
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
								Name: "tfcontrolcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name: "control-logs",
								MountPath: "/var/log/contrail",
							},{
								Name: "dns-config",
								MountPath: "/etc/contrail",								
						}},
					},{
						Image:   dnsImage,
						Name:    "dns",
						ImagePullPolicy: pullPolicy,
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfcontrolcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name: "control-logs",
								MountPath: "/var/log/contrail",
							},{
								Name: "dns-config",
								MountPath: "/etc/contrail",								
						}},
					},{
						Image:   nodeManagerImage,
						Name:    "control-nodemgr",
						ImagePullPolicy: pullPolicy,
						Env: []corev1.EnvVar{{
							Name: "NODE_TYPE",
							Value: "control",
						}},
						EnvFrom: []corev1.EnvFromSource{{
							ConfigMapRef: &corev1.ConfigMapEnvSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "tfcontrolcmv1",
								},
							},
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name: "control-logs",
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
							Name: "control-logs",
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
						{
							Name: "dns-config",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{
								},
							},
						},
					},
				},
 			},
		},
        }
        // Set ControlCluster instance as the owner and controller
        controllerutil.SetControllerReference(m, dep, r.scheme)
        return dep
}

// labelsForControlCluster returns the labels for selecting the resources
// belonging to the given config CR name.
func labelsForControlCluster(name string) map[string]string {
        return map[string]string{"app": "control", "control_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
        var podNames []string
        for _, pod := range pods {
                podNames = append(podNames, pod.Name)
        }
        return podNames
}

