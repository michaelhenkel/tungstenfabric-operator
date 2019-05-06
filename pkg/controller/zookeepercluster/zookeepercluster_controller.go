package zookeepercluster

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

var log = logf.Log.WithName("controller_zookeepercluster")


func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileZookeeperCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("zookeepercluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.ZookeeperCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.ZookeeperCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileZookeeperCluster{}

type ReconcileZookeeperCluster struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileZookeeperCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ZookeeperCluster")

	instance := &tfv1alpha1.ZookeeperCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	var size int32
	var baseConfigMap map[string]string
	baseConfigMap = make(map[string]string)

	baseInstance := &tfv1alpha1.TungstenfabricConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, baseInstance)
	if err != nil && errors.IsNotFound(err){
		reqLogger.Info("baseconfig instance not found")
	} else {
		for k,v := range(baseInstance.Spec.General){
			baseConfigMap[k] = v
		}
		for k,v := range(baseInstance.Spec.ZookeeperCluster){
			baseConfigMap[k] = v
		}
		if instance.Spec.Size == "" {
			if baseInstance.Spec.ZookeeperCluster["size"] != "" {
				instance.Spec.Size = baseInstance.Spec.ZookeeperCluster["size"]
			}
		}
		if instance.Spec.ImagePullPolicy == "" {
			instance.Spec.ImagePullPolicy = baseInstance.Spec.General["imagePullPolicy"]
		}
		if instance.Spec.HostNetwork == ""{
			instance.Spec.HostNetwork  = baseInstance.Spec.General["hostNetwork"]
		}
		if instance.Spec.Image == "" {
			instance.Spec.Image = baseInstance.Spec.Images["zookeeper"]
		}
	}

	size64, err := strconv.ParseInt(instance.Spec.Size, 10, 64)
	if err != nil {
		return reconcile.Result{}, err
	}
	size = int32(size64)

	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "zookeeper-" + instance.Name, Namespace: instance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deploymentForZookeeperCluster(instance, size)
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
		// Spec updated - return and requeue
		return reconcile.Result{Requeue: true}, nil
	}

	// Update the ZookeeperCluster status with the pod names
	// List the pods for this zookeeper's deployment
	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(labelsForZookeeperCluster(instance.Name))
	listOps := &client.ListOptions{
		Namespace:     instance.Namespace,
		LabelSelector: labelSelector,
	}
	err = r.client.List(context.TODO(), listOps, podList)
	if err != nil {
		reqLogger.Error(err, "Failed to list pods.", "ZookeeperCluster.Namespace", instance.Namespace, "ZookeeperCluster.Name", instance.Name)
		return reconcile.Result{}, err
	}

	podNames := getPodNames(podList.Items)
	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
		instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update ZookeeperCluster status.")
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
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfzookeepercmv1", Namespace: instance.Namespace}, foundConfigmap)
		if err != nil && errors.IsNotFound(err) {
			cm := r.configmapForZookeeperCluster(instance, podIpList, baseConfigMap)
			reqLogger.Info("Creating a new Configmap.", "Configmap.Namespace", cm.Namespace, "Configmap.Name", cm.Name)
			err = r.client.Create(context.TODO(), cm)
			if err != nil {
				reqLogger.Error(err, "Failed to create new Configmap.", "Configmap.Namespace", cm.Namespace, "Configmap.Name", cm.Name)
				return reconcile.Result{}, err
			}
		} else if err != nil {
			reqLogger.Error(err, "Failed to get ConfigMap.")
			return reconcile.Result{}, err
		} else {
			cm := r.configmapForZookeeperCluster(instance, podIpList, baseConfigMap)
			err = r.client.Update(context.TODO(), cm)
			if err != nil {
				reqLogger.Error(err, "Failed to update Configmap.", "Configmap.Namespace", cm.Namespace, "Configmap.Name", cm.Name)
				return reconcile.Result{}, err
			}
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

func (r *ReconcileZookeeperCluster) configmapForZookeeperCluster(m *tfv1alpha1.ZookeeperCluster, podIpList []string, baseConfigMap map[string]string) *corev1.ConfigMap {
	nodeListString := strings.Join(podIpList,",")

	baseConfigMap["CONTROLLER_NODES"] = nodeListString
	baseConfigMap["ZOOKEEPER_NODES"] = nodeListString
	baseConfigMap["NODE_TYPE"] = "config-database"

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfzookeepercmv1",
			Namespace: m.Namespace,
		},
		Data: baseConfigMap,
	}
	controllerutil.SetControllerReference(m, configMap, r.scheme)
	return configMap
}
// deploymentForZookeeperCluster returns a zookeeper Deployment object
func (r *ReconcileZookeeperCluster) deploymentForZookeeperCluster(m *tfv1alpha1.ZookeeperCluster, size int32) *appsv1.Deployment {
	ls := labelsForZookeeperCluster(m.Name)
	replicas := size
	zookeeperImage := m.Spec.Image
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
	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "zookeeper-" + m.Name,
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
						}},
						Containers: []corev1.Container{{
							Image:   zookeeperImage,
							Name:    "zookeeper",
							ImagePullPolicy: pullPolicy,
							EnvFrom: []corev1.EnvFromSource{{
								ConfigMapRef: &corev1.ConfigMapEnvSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "tfzookeepercmv1",
									},
								},
							}},
							VolumeMounts: []corev1.VolumeMount{{
								Name: "zookeeper-data",
								MountPath: "/var/lib/zookeeper",
							},{
								Name: "zookeeper-logs",
								MountPath: "/var/log/zookeeper",
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
        // Set ZookeeperCluster instance as the owner and controller
        controllerutil.SetControllerReference(m, dep, r.scheme)
        return dep
}

// labelsForZookeeperCluster returns the labels for selecting the resources
// belonging to the given zookeeper CR name.
func labelsForZookeeperCluster(name string) map[string]string {
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
