package v1alpha1

import (
	"context"
	"strconv"
	"strings"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)
var log = logf.Log.WithName("TungstenFabricResource")
var err error

type TungstenFabricResource interface{
	CreateConfigMap(client.Client) (*corev1.ConfigMap, error)
	CreateDeployment(client.Client) (*appsv1.Deployment, error)
	UpdateDeployment(client.Client, *appsv1.Deployment) error
	GetPodNames(client.Client) ([]string, error)
	WaitForInitContainer(client.Client) (bool, error)
	LabelPod(client.Client, string) (*corev1.Pod, error)
	GetNodeIpList() string
}

type ClusterResource struct {
	Name string
	InstanceName string
	InstanceNamespace string
	Containers []*Container
	General *General
	ResourceConfig map[string]string
	BaseInstance *TungstenfabricConfig
	WaitFor []string
	CassandraInstance *CassandraCluster
	ZookeeperInstance *ZookeeperCluster
	RabbitmqInstance *RabbitmqCluster
	ConfigInstance *ConfigCluster
	ControlInstance *ControlCluster
	StatusVolume bool
	LogVolume bool
	DataVolume bool
	UnixSocketVolume bool
	HostUserBinVolume bool
	InitContainer bool
	NodeInitContainer bool
	NodeIpList string
}

func (c *ClusterResource) CreateConfigMap(client client.Client) (*corev1.ConfigMap, error) {
	for _, waitResource := range(c.WaitFor){
		err = getResourceConfig(c, client, waitResource)
		if err != nil {
			return nil, err
		}
	}
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tf" + c.Name + "cmv1",
			Namespace: c.InstanceNamespace,
		},
		Data: c.ResourceConfig,
	}
	return cm, nil
}

func (c *ClusterResource) CreateDeployment(client client.Client) (*appsv1.Deployment, error) {

	for _, container := range(c.Containers){
		if container.Image == "" {
			container.Image = c.BaseInstance.Spec.Images[container.Name]
		}

		if container.PullPolicy == "" {
			container.PullPolicy = c.BaseInstance.Spec.General.PullPolicy
		}
		
	}
	var sizeString string
	var hostNetworkString string
	if c.General != nil {
		if c.General.Size == "" {
			sizeString = c.BaseInstance.Spec.General.Size
		} else {
			sizeString = c.General.Size
		}
		if c.General.HostNetwork == ""{
			hostNetworkString = c.BaseInstance.Spec.General.HostNetwork		
		} else {
			hostNetworkString = c.General.HostNetwork
		}
	} else {
		sizeString = c.BaseInstance.Spec.General.Size
		hostNetworkString = c.BaseInstance.Spec.General.HostNetwork
	}

	size64, err := strconv.ParseInt(sizeString, 10, 64)
	if err != nil {
		return nil, err
	}
	size := int32(size64)
	

	var hostNetworkBool bool
	if hostNetworkString == "true" {
		hostNetworkBool = true
	} else {
		hostNetworkBool = false
	}
	createDataVolume := false
	createLogVolume := false
	createUnixSocketVolume := false
	createHostUserBinVolume := false
	createEtcContrailVolume := false

	var containerList []corev1.Container
	for _, container := range(c.Containers){

		var envList []corev1.EnvVar
		if len(container.Env) > 0 {
			for k, v := range(container.Env){
				env := corev1.EnvVar{
					Name: k,
					Value: v,
				}
				envList = append(envList, env)
			}
		}
		var volumeMountList []corev1.VolumeMount
		if container.LogVolumePath != ""{
			logVolumeMount := corev1.VolumeMount{
				Name: c.Name + "-logs",
				MountPath: container.LogVolumePath,
			}
			volumeMountList = append(volumeMountList, logVolumeMount)
			createLogVolume = true
		}
		if container.DataVolumePath != ""{
			dataVolumeMount := corev1.VolumeMount{
				Name: c.Name + "-data",
				MountPath: container.DataVolumePath,
			}
			volumeMountList = append(volumeMountList, dataVolumeMount)
			createDataVolume = true
		}
		if container.UnixSocketVolume{
			unixSocketVolume := corev1.VolumeMount{
				Name: "docker-unix-socket",
				MountPath: "/mnt",
			}
			volumeMountList = append(volumeMountList, unixSocketVolume)
			createUnixSocketVolume = true
		}
		if container.HostUserBinVolume{
			hostUserBinVolume := corev1.VolumeMount{
				Name: "host-usr-bin",
				MountPath: "/host/usr/bin",
			}
			volumeMountList = append(volumeMountList, hostUserBinVolume)
			createHostUserBinVolume = true
		}
		if container.EtcContrailVolume{
			etcContrailVolume := corev1.VolumeMount{
				Name: "etc-contrail",
				MountPath: "/etc/contrail",
			}
			volumeMountList = append(volumeMountList, etcContrailVolume)
			createEtcContrailVolume = true
		}

		deploymentContainer := corev1.Container{
			Image: container.Image,
			Name: strings.ToLower(container.Name),
			SecurityContext: &corev1.SecurityContext{
				Privileged: &container.Privileged,
			},
			ImagePullPolicy: corev1.PullPolicy(container.PullPolicy),
			EnvFrom: []corev1.EnvFromSource{{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "tf" + c.Name + "cmv1",
					},
				},
			}},
			Env: envList,
			VolumeMounts: volumeMountList,
		}
		containerList = append(containerList, deploymentContainer)
	}

	for _, waitResource := range(c.WaitFor){
		err = getResourceConfig(c, client, waitResource)
		if err != nil {
			return nil, err
		}
	}

	privileged := true
	var initContainerList []corev1.Container
	initContainer := corev1.Container{
		Image:   "busybox",
		Name:    "init",
		Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
		VolumeMounts: []corev1.VolumeMount{{
			Name: "status",
			MountPath: "/tmp/podinfo",
		}},
	}
	nodeInitContainer := corev1.Container{
		Image:   c.BaseInstance.Spec.Images["nodeinit"],
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
					Name: "tf" + c.Name + "cmv1",
				},
			},
		}},
		Env: []corev1.EnvVar{{
			Name: "CONTRAIL_STATUS_IMAGE",
			Value: c.BaseInstance.Spec.Images["status"],
		}},
	}

	if c.InitContainer{
		initContainerList = append(initContainerList, initContainer)
	}

	if c.NodeInitContainer{
		initContainerList = append(initContainerList, nodeInitContainer)
		createHostUserBinVolume = true
	}

	var volumeList []corev1.Volume
	statusVolume := corev1.Volume{
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
	}

	logVolume := corev1.Volume{
		Name: c.Name + "-logs",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/log/contrail/" + c.Name,
			},
		},
	}

	dataVolume := corev1.Volume{
		Name: c.Name + "-data",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/contrail/" + c.Name,
			},
		},
	}

	unixSocketVolume := corev1.Volume{
		Name: "docker-unix-socket",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/run",
			},
		},
	}

	hostUserBinVolume := corev1.Volume{
		Name: "host-usr-bin",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/usr/bin",
			},
		},
	}

	etcContrailVolume := corev1.Volume{
		Name: "etc-contrail",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{
			},
		},
	}


	volumeList = append(volumeList, statusVolume)

	if createLogVolume {
		volumeList = append(volumeList, logVolume)
	}

	if createDataVolume {
		volumeList = append(volumeList, dataVolume)
	}

	if createUnixSocketVolume {
		volumeList = append(volumeList, unixSocketVolume)
	}

	if createHostUserBinVolume {
		volumeList = append(volumeList, hostUserBinVolume)
	}

	if createEtcContrailVolume {
		volumeList = append(volumeList, etcContrailVolume)
	}

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tf" + c.Name + "-" + c.InstanceName,
			Namespace: c.InstanceNamespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": c.Name, c.Name + "_cr": c.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": c.Name, c.Name + "_cr": c.Name},
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
					InitContainers: initContainerList,
					Containers: containerList,
					Volumes: volumeList,
				},
			},
		},		
	}
	return dep, nil
}

func getResourceConfig(c *ClusterResource, client client.Client, resourceType string) error {
	reqLogger := log.WithValues("Request.Namespace", c.InstanceNamespace, "Request.Name", c.InstanceName)
	reqLogger.Info("getting " + resourceType + " config")

	switch resourceType{
	case "cassandra":
		err = client.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.CassandraInstance)

	case "zookeeper":
		err = client.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.ZookeeperInstance)

	case "rabbitmq":
		err = client.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.RabbitmqInstance)

	case "config":
		err = client.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.ConfigInstance)
	}

	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info(resourceType + " instance not found")
		return err
	}
	reqLogger.Info(resourceType + " instance")
	configMap := &corev1.ConfigMap{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "tf" + resourceType + "cmv1", Namespace: c.InstanceNamespace}, configMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info(resourceType + " configmap not found")
		return err
	}

	switch resourceType{
	case "cassandra":
		c.ResourceConfig["CONFIGDB_PORT"] = configMap.Data["CASSANDRA_PORT"]
		c.ResourceConfig["CONFIGDB_CQL_PORT"] = configMap.Data["CASSANDRA_CQL_PORT"]
		c.ResourceConfig["CONFIGDB_NODES"] = configMap.Data["CASSANDRA_SEEDS"]

	case "zookeeper":
		c.ResourceConfig["ZOOKEEPER_NODES"] = configMap.Data["ZOOKEEPER_NODES"]
		c.ResourceConfig["ZOOKEEPER_NODE_PORT"] = configMap.Data["ZOOKEEPER_PORT"]

	case "rabbitmq":
		c.ResourceConfig["RABBITMQ_NODES"] = configMap.Data["RABBITMQ_NODES"]
		c.ResourceConfig["RABBITMQ_NODE_PORT"] = configMap.Data["RABBITMQ_NODE_PORT"]

	case "config":
		c.ResourceConfig["CONFIG_NODES"] = configMap.Data["CONTROLLER_NODES"]
		c.ResourceConfig["ANALYTICS_NODES"] = configMap.Data["CONTROLLER_NODES"]
	}

	return nil
}



func (c *ClusterResource) UpdateDeployment(client client.Client, deployment *appsv1.Deployment) error {

	var sizeString string
	if c.General != nil {
		if c.General.Size == "" {
			sizeString = c.BaseInstance.Spec.General.Size
		} else {
			sizeString = c.General.Size
		}
	} else {
		sizeString = c.BaseInstance.Spec.General.Size
	}

	size64, err := strconv.ParseInt(sizeString, 10, 64)
	if err != nil {
		return err
	}
	size := int32(size64)

	foundDeployment := &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "tf" + c.Name + "-" + c.InstanceName, Namespace: c.InstanceNamespace}, foundDeployment)
	if *foundDeployment.Spec.Replicas != size {
		foundDeployment.Spec.Replicas = &size
		err = client.Update(context.TODO(), foundDeployment)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (c *ClusterResource) GetPodNames(cl client.Client) ([]string, error) {
	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(map[string]string{"app": c.Name, c.Name + "_cr": c.Name})
	var podNames []string

	listOps := &client.ListOptions{
		Namespace:     c.InstanceNamespace,
		LabelSelector: labelSelector,
	}
	err = cl.List(context.TODO(), listOps, podList)
	if err != nil {
		return podNames, err
	}
	for _, pod := range podList.Items {
		podNames = append(podNames, pod.Name)
	}
	return podNames, nil
}

func (c *ClusterResource) WaitForInitContainer(cl client.Client) (bool, error) {
	var sizeString string
	if c.General != nil {
		if c.General.Size == "" {
			sizeString = c.BaseInstance.Spec.General.Size
		} else {
			sizeString = c.General.Size
		}
	} else {
		sizeString = c.BaseInstance.Spec.General.Size
	}

	size64, err := strconv.ParseInt(sizeString, 10, 64)
	if err != nil {
		return false, err
	}
	size := int32(size64)
	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(map[string]string{"app": c.Name, c.Name + "_cr": c.Name})

	listOps := &client.ListOptions{
		Namespace:     c.InstanceNamespace,
		LabelSelector: labelSelector,
	}
	err = cl.List(context.TODO(), listOps, podList)
	if err != nil {
		return false, err
	}

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
		c.NodeIpList = strings.Join(podIpList,",")
		return true, nil
	}

	return false, nil
}

func (c ClusterResource) LabelPod(cl client.Client, podName string) (*corev1.Pod, error) {
	foundPod := &corev1.Pod{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: podName, Namespace: c.InstanceNamespace}, foundPod)
	if err != nil {
		return foundPod, err
	}
	podMetaData := foundPod.ObjectMeta
	podLabels := podMetaData.Labels
	podLabels["status"] = "ready"
	foundPod.ObjectMeta.Labels = podLabels
	return foundPod, nil
}

func (c ClusterResource) GetNodeIpList() string {
	return c.NodeIpList
}