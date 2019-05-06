package kubemanagercluster

import (
	"context"
	"reflect"
	"strings"
	"strconv"
	"net"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
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
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"gopkg.in/yaml.v2"
)

var log = logf.Log.WithName("controller_kubemanagercluster")

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKubemanagerCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("kubemanagercluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.KubemanagerCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.KubemanagerCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileKubemanagerCluster{}

type ReconcileKubemanagerCluster struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileKubemanagerCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling KubemanagerCluster")

	instance := &tfv1alpha1.KubemanagerCluster{}
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
		for k,v := range(baseInstance.Spec.KubemanagerCluster){
			configMap[k] = v
		}
		if instance.Spec.Size == "" {
			if baseInstance.Spec.KubemanagerCluster["size"] != "" {
				instance.Spec.Size = baseInstance.Spec.KubemanagerCluster["size"]
			}
		}
		if instance.Spec.ImagePullPolicy == "" {
			instance.Spec.ImagePullPolicy = baseInstance.Spec.General["imagePullPolicy"]
		}
		if instance.Spec.HostNetwork == ""{
			instance.Spec.HostNetwork  = baseInstance.Spec.General["hostNetwork"]
		}
		if instance.Spec.KubeManagerImage == "" {
			instance.Spec.KubeManagerImage = baseInstance.Spec.Images["kubemanager"]
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


	zookeeperInstance := &tfv1alpha1.ZookeeperCluster{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, zookeeperInstance)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("zookeeper instance not found")
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("zookeeper instance")
	zookeeperConfigMap := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfzookeepercmv1", Namespace: instance.Namespace}, zookeeperConfigMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("zookeeper configmap not found")
		return reconcile.Result{Requeue: true}, nil
	}
	configMap["ZOOKEEPER_NODES"] = zookeeperConfigMap.Data["ZOOKEEPER_NODES"]
	configMap["ZOOKEEPER_NODE_PORT"] = zookeeperConfigMap.Data["ZOOKEEPER_PORT"]

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

	// create rbac
	existingServiceAccount := &corev1.ServiceAccount{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "contrail-service-account-" + instance.Name, Namespace: instance.Namespace}, existingServiceAccount)
	if err != nil && errors.IsNotFound(err) {
		serviceAccount := r.serviceAccountForKubemanagerCluster(instance)
		reqLogger.Info("Creating Service Account", "serviceAccount.Namespace", serviceAccount.Namespace, "serviceAccount.Name", serviceAccount.Name)
		err = r.client.Create(context.TODO(), serviceAccount)
		if err != nil {
			reqLogger.Error(err, "Failed to create serviceAccount.", "serviceAccount.Namespace", serviceAccount.Namespace, "serviceAccount.Name", serviceAccount.Name)
			return reconcile.Result{}, err
		}
	}

	existingClusterRole := &rbacv1.ClusterRole{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "contrail-cluster-role-" + instance.Name, Namespace: instance.Namespace}, existingClusterRole)
	if err != nil && errors.IsNotFound(err) {
		clusterRole := r.clusterRoleForKubemanagerCluster(instance)
		reqLogger.Info("Creating Cluster Role", "clusterRole.Namespace", clusterRole.Namespace, "clusterRole.Name", clusterRole.Name)
		err = r.client.Create(context.TODO(), clusterRole)
		if err != nil && !errors.IsAlreadyExists(err) {
			reqLogger.Error(err, "Failed to create clusterRole.", "clusterRole.Namespace", clusterRole.Namespace, "clusterRole.Name", clusterRole.Name)
			return reconcile.Result{}, err
		}
	}

	existingClusterRoleBinding := &rbacv1.ClusterRoleBinding{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "contrail-cluster-role-binding-" + instance.Name, Namespace: instance.Namespace}, existingClusterRoleBinding)
	if err != nil && errors.IsNotFound(err) {
		clusterRoleBinding := r.clusterRoleBindingForKubemanagerCluster(instance)
		reqLogger.Info("Creating Cluster Role Binding", "clusterRoleBinding.Namespace", clusterRoleBinding.Namespace, "clusterRoleBinding.Name", clusterRoleBinding.Name)
		err = r.client.Create(context.TODO(), clusterRoleBinding)
		if err != nil && !errors.IsAlreadyExists(err){
			reqLogger.Error(err, "Failed to create clusterRoleBinding.", "clusterRoleBinding.Namespace", clusterRoleBinding.Namespace, "clusterRoleBinding.Name", clusterRoleBinding.Name)
			return reconcile.Result{}, err
		}
	}

	existingSecret := &corev1.Secret{}
	reqLogger.Info("Trying to get secret")
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "contrail-kube-manager-token-" + instance.Name, Namespace: instance.Namespace}, existingSecret)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("secret contrail-kube-manager-token doesn't exist. Creating it")
		secret := r.secretForKubemanagerCluster(instance)
		reqLogger.Info("Creating secret", "secret.Namespace", secret.Namespace, "secret.Name", secret.Name)
		err = r.client.Create(context.TODO(), secret)
		if err != nil && !errors.IsAlreadyExists(err){
			reqLogger.Error(err, "Failed to create secret.", "secret.Namespace", secret.Namespace, "secret.Name", secret.Name)
			return reconcile.Result{}, err
		}
	} else {
		reqLogger.Info("secret contrail-kube-manager-token exists")
	}


	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "kubemanager-" + instance.Name, Namespace: instance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.deploymentForKubemanagerCluster(instance, size)
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
	labelSelector := labels.SelectorFromSet(labelsForKubemanagerCluster(instance.Name))

	listOps := &client.ListOptions{
		Namespace:     instance.Namespace,
		LabelSelector: labelSelector,
	}
	err = r.client.List(context.TODO(), listOps, podList)
	if err != nil {
		reqLogger.Error(err, "Failed to list pods.", "KubemanagerCluster.Namespace", instance.Namespace, "KubemanagerCluster.Name", instance.Name)
		return reconcile.Result{}, err
	}

	podNames := getPodNames(podList.Items)
	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
		instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update KubemanagerCluster status.")
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
		foundKubemanagerClustermap := &corev1.ConfigMap{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: "tfkubemanagercmv1", Namespace: foundKubemanagerClustermap.Namespace}, foundKubemanagerClustermap)
		if err != nil && errors.IsNotFound(err) {
			cm := r.configmapForKubemanagerCluster(instance, podIpList, configMap)
			reqLogger.Info("Creating a new KubemanagerClustermap.", "KubemanagerClustermap.Namespace", cm.Namespace, "KubemanagerClustermap.Name", cm.Name)
			err = r.client.Create(context.TODO(), cm)
			if err != nil {
				reqLogger.Error(err, "Failed to create new KubemanagerClustermap.", "KubemanagerClustermap.Namespace", cm.Namespace, "KubemanagerClustermap.Name", cm.Name)
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


func (r *ReconcileKubemanagerCluster) configmapForKubemanagerCluster(m *tfv1alpha1.KubemanagerCluster, podIpList []string, configMap map[string]string) *corev1.ConfigMap {
	
	
	reqLogger := log.WithValues("Request.Namespace", m.Namespace, "Request.Name", m.Name)
	nodeListString := strings.Join(podIpList,",")

	// check if kubeadm config map exists

	controlPlaneEndpoint := ""
	clusterName := ""
	podSubnet := ""
	serviceSubnet := ""
	controlPlaneEndpointHost := ""
	controlPlaneEndpointPort := ""

	reqLogger.Info("checking in cluster config")
	config, err := rest.InClusterConfig()
	if err == nil {
		clientset, err := kubernetes.NewForConfig(config)
		if err == nil {
			kubeadmConfigMapClient := clientset.CoreV1().ConfigMaps("kube-system")
			kcm, err := kubeadmConfigMapClient.Get("kubeadm-config", metav1.GetOptions{})
			if err != nil {
				reqLogger.Error(err, "Failed getting kubeadm config")
			}
			clusterConfig := kcm.Data["ClusterConfiguration"]
			clusterConfigByte := []byte(clusterConfig)
			clusterConfigMap := make(map[interface{}]interface{})
			err = yaml.Unmarshal(clusterConfigByte, &clusterConfigMap)
			if err != nil {
				reqLogger.Error(err, "Failed to unmarshal.")
				panic(err)
			}
			controlPlaneEndpoint = clusterConfigMap["controlPlaneEndpoint"].(string)
			controlPlaneEndpointHost, controlPlaneEndpointPort, _ = net.SplitHostPort(controlPlaneEndpoint)
			clusterName = clusterConfigMap["clusterName"].(string)
			networkConfig := make(map[interface{}]interface{})
			networkConfig = clusterConfigMap["networking"].(map[interface{}]interface{})
			podSubnet = networkConfig["podSubnet"].(string)
			serviceSubnet = networkConfig["serviceSubnet"].(string)
		}
	}

	if configMap["KUBERNETES_API_SERVER"] == "" {
		configMap["KUBERNETES_API_SERVER"] = controlPlaneEndpointHost
	}

	if configMap["KUBERNETES_API_SECURE_PORT"] == "" {
		configMap["KUBERNETES_API_SECURE_PORT"] = controlPlaneEndpointPort
	}

	if configMap["KUBERNETES_POD_SUBNETS"] == "" {
		configMap["KUBERNETES_POD_SUBNETS"] = podSubnet
	}

	if configMap["KUBERNETES_SERVICE_SUBNETS"] == "" {
		configMap["KUBERNETES_SERVICE_SUBNETS"] = serviceSubnet
	}

	if configMap["KUBERNETES_CLUSTER_NAME"] == "" {
		configMap["KUBERNETES_CLUSTER_NAME"] = clusterName
	}
	configMap["CONTROLLER_NODES"] = nodeListString
	configMap["ANALYTICS_NODES"] = configMap["CONFIG_NODES"]

	newConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tfkubemanagercmv1",
			Namespace: m.Namespace,
		},
		Data: configMap,
	}
	controllerutil.SetControllerReference(m, newConfigMap, r.scheme)
	return newConfigMap
}
func (r *ReconcileKubemanagerCluster) clusterRoleForKubemanagerCluster(m *tfv1alpha1.KubemanagerCluster) *rbacv1.ClusterRole {
	cr := &rbacv1.ClusterRole {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-cluster-role-" + m.Name,
			Namespace: m.Namespace,
		},
		Rules: []rbacv1.PolicyRule{{
			Verbs: []string{
				"get",
				"list",
				"update",
				"watch",
				"patch",
			},
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"pods",
				"services",
				"endpoints",
				"events",
			},
		}},
	}
	controllerutil.SetControllerReference(m, cr, r.scheme)
	return cr
}

func (r *ReconcileKubemanagerCluster) serviceAccountForKubemanagerCluster(m *tfv1alpha1.KubemanagerCluster) *corev1.ServiceAccount {
	sa := &corev1.ServiceAccount {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-service-account-" + m.Name,
			Namespace: m.Namespace,
		},
	}
	controllerutil.SetControllerReference(m, sa, r.scheme)
	return sa
}

func (r *ReconcileKubemanagerCluster) clusterRoleBindingForKubemanagerCluster(m *tfv1alpha1.KubemanagerCluster) *rbacv1.ClusterRoleBinding {
	crb := &rbacv1.ClusterRoleBinding {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac/v1",
			Kind:       "ClusterRoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-cluster-role-binding-" + m.Name,
			Namespace: m.Namespace,
		},
		Subjects: []rbacv1.Subject{{
			Kind: "ServiceAccount",
			Name: "contrail-service-account-" + m.Name,
			Namespace: m.Namespace,
			}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind: "ClusterRole",
			Name: "contrail-cluster-role-" + m.Name,
			},
	}
	controllerutil.SetControllerReference(m, crb, r.scheme)
	return crb
}

func (r *ReconcileKubemanagerCluster) secretForKubemanagerCluster(m *tfv1alpha1.KubemanagerCluster) *corev1.Secret {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-kube-manager-token-" + m.Name,
			Namespace: m.Namespace,
			Annotations: map[string]string{
				"kubernetes.io/service-account.name": "contrail-service-account",
			},
		},
		Type: "kubernetes.io/service-account-token",
	}
	controllerutil.SetControllerReference(m, secret, r.scheme)
	return secret
}
// deploymentForKubemanagerCluster returns a config Deployment object
func (r *ReconcileKubemanagerCluster) deploymentForKubemanagerCluster(m *tfv1alpha1.KubemanagerCluster, size int32) *appsv1.Deployment {
	ls := labelsForKubemanagerCluster(m.Name)
	replicas := size
	kubeManagerImage := m.Spec.KubeManagerImage
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
			Name:      "kubemanager-" + m.Name,
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
				ServiceAccountName: "contrail-service-account-" + m.Name,
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
								Name: "tfkubemanagercmv1",
							},
						},
					}},
				}},
				Containers: []corev1.Container{{
					Image:   kubeManagerImage,
					Name:    "kube-manager",
					ImagePullPolicy: pullPolicy,
					EnvFrom: []corev1.EnvFromSource{{
						ConfigMapRef: &corev1.ConfigMapEnvSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "tfkubemanagercmv1",
							},
						},
					}},
					VolumeMounts: []corev1.VolumeMount{{
						Name: "kubemanager-logs",
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
						Name: "pod-secret",
						VolumeSource: corev1.VolumeSource{
							Secret: &corev1.SecretVolumeSource{
								SecretName: "contrail-kube-manager-token-" + m.Name,
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
						Name: "kubemanager-logs",
						VolumeSource: corev1.VolumeSource{
							HostPath: &corev1.HostPathVolumeSource{
								Path: "/var/log/contrail/kubemanager",
							},
						},
					},
				},
			},
		},
		},
	}
	// Set KubemanagerCluster instance as the owner and controller
	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}

// labelsForKubemanagerCluster returns the labels for selecting the resources
// belonging to the given config CR name.
func labelsForKubemanagerCluster(name string) map[string]string {
        return map[string]string{"app": "kubemanager", "kubemanager_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
        var podNames []string
        for _, pod := range pods {
                podNames = append(podNames, pod.Name)
        }
        return podNames
}
