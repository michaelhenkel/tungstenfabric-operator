package v1alpha1

import (
	"context"
	"strconv"
	"strings"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"k8s.io/apimachinery/pkg/runtime"
)
var log = logf.Log.WithName("TungstenFabricResource")
var err error

type TungstenFabricResource interface{
	CreateConfigMap(client.Client, metav1.Object, *runtime.Scheme) error
	CreateDeployment(client.Client, metav1.Object, *runtime.Scheme) error
	UpdateDeployment(client.Client, *appsv1.Deployment) error
	GetPodNames(client.Client) ([]string, error)
	WaitForInitContainer(client.Client) (bool, error)
	LabelPod(client.Client, string) (*corev1.Pod, error)
	GetNodeIpList() string
	CreateRbac(client.Client, metav1.Object, *runtime.Scheme) error
}

type ClusterResource struct {
	Name string
	InstanceName string
	InstanceNamespace string
	Containers []*Container
	InitContainers []*Container
	General *General
	ResourceConfig map[string]string
	BaseInstance *TungstenfabricManager
	WaitFor []string
	CassandraInstance *CassandraCluster
	ZookeeperInstance *ZookeeperCluster
	RabbitmqInstance *RabbitmqCluster
	ConfigInstance *ConfigCluster
	ControlInstance *ControlCluster
	KubemanagerInstance *KubemanagerCluster
	WebuiInstance *WebuiCluster
	VrouterInstance *Vrouter
	StatusVolume bool
	LogVolume bool
	DataVolume bool
	UnixSocketVolume bool
	HostUserBinVolume bool
	NodeIpList string
	ServiceAccount bool
	Type string
}

func (c *ClusterResource) CreateConfigMap(cl client.Client, instance metav1.Object, scheme *runtime.Scheme) error {
	for _, waitResource := range(c.WaitFor){
		err = getResourceConfig(c, cl, waitResource)
		if err != nil {
			return err
		}
	}

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tf" + c.Name + "cmv1",
			Namespace: c.InstanceNamespace,
		},
		Data: c.ResourceConfig,
	}

	existingConfigMap := &corev1.ConfigMap{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: "tf" + c.Name + "cmv1", Namespace: c.InstanceNamespace}, existingConfigMap)
	if err != nil && errors.IsNotFound(err) {
		controllerutil.SetControllerReference(instance, configMap, scheme)
		err = cl.Create(context.TODO(), configMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClusterResource) CreateDeployment(cl client.Client, instance metav1.Object, scheme *runtime.Scheme) error {

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
		return err
	}
	size := int32(size64)
	

	var hostNetworkBool bool
	if hostNetworkString == "true" {
		hostNetworkBool = true
	} else {
		hostNetworkBool = false
	}


	createStatusVolume := false
	createVarLogCniVolume := false
	createEtcCniVolume := false
	createVarCrashesVolume := false
	createVarLibContrailVolume := false
	createLibModulesVolume := false
	createUsrSrcVolume := false
	createHostBinVolume := false
	createNetworkScriptsVolume := false
	createDevVolume := false
	createEtcContrailVolume := false
	createHostUserBinVolume := false
	createUnixSocketVolume := false
	createDataVolume := false
	createLogVolume := false
	createOptBinCniVolume := false

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
		if container.DevVolume{
			devVolume := corev1.VolumeMount{
				Name: "dev",
				MountPath: "/dev",
			}
			volumeMountList = append(volumeMountList, devVolume)
			createDevVolume = true
		}
		if container.NetworkScriptsVolume{
			networkScriptsVolume := corev1.VolumeMount{
				Name: "network-scripts",
				MountPath: "/etc/sysconfig/network-scripts",
			}
			volumeMountList = append(volumeMountList, networkScriptsVolume)
			createNetworkScriptsVolume = true
		}
		if container.HostBinVolume{
			hostBinVolume := corev1.VolumeMount{
				Name: "host-bin",
				MountPath: "/bin",
			}
			volumeMountList = append(volumeMountList, hostBinVolume)
			createHostBinVolume = true
		}
		if container.UsrSrcVolume{
			usrSrcVolume := corev1.VolumeMount{
				Name: "usr-src",
				MountPath: "/usr/src",
			}
			volumeMountList = append(volumeMountList, usrSrcVolume)
			createUsrSrcVolume = true
		}
		if container.LibModulesVolume{
			libModulesVolume := corev1.VolumeMount{
				Name: "lib-modules",
				MountPath: "/lib/modules",
			}
			volumeMountList = append(volumeMountList, libModulesVolume)
			createLibModulesVolume = true
		}
		if container.VarLibContrailVolume{
			varLibContrailVolume := corev1.VolumeMount{
				Name: "var-lib-contrail",
				MountPath: "/var/lib/contrail",
			}
			volumeMountList = append(volumeMountList, varLibContrailVolume)
			createVarLibContrailVolume = true
		}
		if container.VarCrashesVolume{
			varCrashesVolume := corev1.VolumeMount{
				Name: "var-crashes",
				MountPath: "/var/contrail/crashes",
			}
			volumeMountList = append(volumeMountList, varCrashesVolume)
			createVarCrashesVolume = true
		}
		if container.EtcCniVolume{
			etcCniVolume := corev1.VolumeMount{
				Name: "var-crashes",
				MountPath: "/var/contrail/crashes",
			}
			volumeMountList = append(volumeMountList, etcCniVolume)
			createEtcCniVolume = true
		}
		if container.OptBinCniVolume{
			optBinCniVolume := corev1.VolumeMount{
				Name: "opt-bin-cni",
				MountPath: "/opt/bin/cni",
			}
			volumeMountList = append(volumeMountList, optBinCniVolume)
			createOptBinCniVolume = true
		}
		if container.VarLogCniVolume{
			varLogCniVolume := corev1.VolumeMount{
				Name: "var-crashes",
				MountPath: "/var/contrail/crashes",
			}
			volumeMountList = append(volumeMountList, varLogCniVolume)
			createVarLogCniVolume = true
		}
		if container.StatusVolume{
			statusVolume := corev1.VolumeMount{
				Name: "status",
				MountPath: "/tmp/podinfo",
			}
			volumeMountList = append(volumeMountList, statusVolume)
			createStatusVolume = true
		}


		var lifeCycle corev1.Lifecycle
		if container.LifeCycleScript != nil {
			lifeCycle = corev1.Lifecycle{
				PreStop: &corev1.Handler{
					Exec: &corev1.ExecAction{
						Command: container.LifeCycleScript,
					},
				},				
			}
		}
		deploymentContainer := corev1.Container{
			Image: container.Image,
			Name: strings.ToLower(container.Name),
			SecurityContext: &corev1.SecurityContext{
				Privileged: &container.Privileged,
			},
			Lifecycle: &lifeCycle,
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

	var initContainerList []corev1.Container
	for _, container := range(c.InitContainers){

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
		if container.DevVolume{
			devVolume := corev1.VolumeMount{
				Name: "dev",
				MountPath: "/dev",
			}
			volumeMountList = append(volumeMountList, devVolume)
			createDevVolume = true
		}
		if container.NetworkScriptsVolume{
			networkScriptsVolume := corev1.VolumeMount{
				Name: "network-scripts",
				MountPath: "/etc/sysconfig/network-scripts",
			}
			volumeMountList = append(volumeMountList, networkScriptsVolume)
			createNetworkScriptsVolume = true
		}
		if container.HostBinVolume{
			hostBinVolume := corev1.VolumeMount{
				Name: "host-bin",
				MountPath: "/bin",
			}
			volumeMountList = append(volumeMountList, hostBinVolume)
			createHostBinVolume = true
		}
		if container.UsrSrcVolume{
			usrSrcVolume := corev1.VolumeMount{
				Name: "usr-src",
				MountPath: "/usr/src",
			}
			volumeMountList = append(volumeMountList, usrSrcVolume)
			createUsrSrcVolume = true
		}
		if container.LibModulesVolume{
			libModulesVolume := corev1.VolumeMount{
				Name: "lib-modules",
				MountPath: "/lib/modules",
			}
			volumeMountList = append(volumeMountList, libModulesVolume)
			createLibModulesVolume = true
		}
		if container.VarLibContrailVolume{
			varLibContrailVolume := corev1.VolumeMount{
				Name: "var-lib-contrail",
				MountPath: "/var/lib/contrail",
			}
			volumeMountList = append(volumeMountList, varLibContrailVolume)
			createVarLibContrailVolume = true
		}
		if container.VarCrashesVolume{
			varCrashesVolume := corev1.VolumeMount{
				Name: "var-crashes",
				MountPath: "/var/contrail/crashes",
			}
			volumeMountList = append(volumeMountList, varCrashesVolume)
			createVarCrashesVolume = true
		}
		if container.EtcCniVolume{
			etcCniVolume := corev1.VolumeMount{
				Name: "var-crashes",
				MountPath: "/var/contrail/crashes",
			}
			volumeMountList = append(volumeMountList, etcCniVolume)
			createEtcCniVolume = true
		}
		if container.VarLogCniVolume{
			varLogCniVolume := corev1.VolumeMount{
				Name: "var-crashes",
				MountPath: "/var/contrail/crashes",
			}
			volumeMountList = append(volumeMountList, varLogCniVolume)
			createVarLogCniVolume = true
		}
		if container.OptBinCniVolume{
			optBinCniVolume := corev1.VolumeMount{
				Name: "opt-bin-cni",
				MountPath: "/opt/bin/cni",
			}
			volumeMountList = append(volumeMountList, optBinCniVolume)
			createOptBinCniVolume = true
		}
		if container.StatusVolume{
			statusVolume := corev1.VolumeMount{
				Name: "status",
				MountPath: "/tmp/podinfo",
			}
			volumeMountList = append(volumeMountList, statusVolume)
			createStatusVolume = true
		}

		statusImageEnv := corev1.EnvVar{
			Name: "CONTRAIL_STATUS_IMAGE",
			Value: c.BaseInstance.Spec.Images["status"],
		}
		envList = append(envList, statusImageEnv)

		var lifeCycle corev1.Lifecycle
		if container.LifeCycleScript != nil {
			lifeCycle = corev1.Lifecycle{
				PreStop: &corev1.Handler{
					Exec: &corev1.ExecAction{
						Command: container.LifeCycleScript,
					},
				},				
			}
		}

		initContainer := corev1.Container{
			Image: container.Image,
			Name: strings.ToLower(container.Name),
			SecurityContext: &corev1.SecurityContext{
				Privileged: &container.Privileged,
			},
			ImagePullPolicy: corev1.PullPolicy(container.PullPolicy),
			Lifecycle: &lifeCycle,
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
		initContainerList = append(initContainerList, initContainer)
	}

	for _, waitResource := range(c.WaitFor){
		err = getResourceConfig(c, cl, waitResource)
		if err != nil {
			return err
		}
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


	devVolume := corev1.Volume{
		Name: "dev",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/dev",
			},
		},
	}

	varLibContrailVolume := corev1.Volume{
		Name: "var-lib-contrail",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/contrail",
			},
		},
	}

	varCrashesVolume := corev1.Volume{
		Name: "var-crashes",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/contrail/crashes",
			},
		},
	}
	
	etcCniVolume := corev1.Volume{
		Name: "etc-cni",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc/cni",
			},
		},
	}

	optBinCniVolume := corev1.Volume {
		Name: "opt-cni-bin",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/opt/cni/bin",
			},
		},
	}

	varLogCniVolume := corev1.Volume{
		Name: "var-log-contrail-cni",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/log/contrail/cni",
			},
		},
	}



	usrSrcVolume := corev1.Volume{
		Name: "usr-src",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/usr/src",
			},
		},
	}

	libModulesVolume := corev1.Volume{
		Name: "lib-modules",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/lib/modules",
			},
		},
	}

	networkScriptsVolume := corev1.Volume{
		Name: "network-scripts",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc/sysconfig/network-scripts",
			},
		},
	}

	hostBinVolume := corev1.Volume{
		Name: "host-bin",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/bin",
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

	if createStatusVolume {
		volumeList = append(volumeList, statusVolume)
	}

	if createVarLogCniVolume {
		volumeList = append(volumeList, varLogCniVolume)
	}

	if createEtcCniVolume {
		volumeList = append(volumeList, etcCniVolume)
	}

	if createVarCrashesVolume {
		volumeList = append(volumeList, varCrashesVolume)
	}

	if createVarLibContrailVolume {
		volumeList = append(volumeList, varLibContrailVolume)
	}

	if createLibModulesVolume {
		volumeList = append(volumeList, libModulesVolume)
	}

	if createUsrSrcVolume {
		volumeList = append(volumeList, usrSrcVolume)
	}

	if createHostBinVolume {
		volumeList = append(volumeList, hostBinVolume)
	}

	if createNetworkScriptsVolume {
		volumeList = append(volumeList, networkScriptsVolume)
	}

	if createDevVolume {
		volumeList = append(volumeList, devVolume)
	}

	if createOptBinCniVolume {
		volumeList = append(volumeList, optBinCniVolume)
	}

	var serviceAccountName string

	if c.ServiceAccount{
		serviceAccountName = "contrail-service-account-" + c.Name
	}

	objectMeta := metav1.ObjectMeta{
		Name: "tf" + c.Name + "-" + c.InstanceName,
		Namespace: c.InstanceNamespace,
	}

	selector := metav1.LabelSelector{
		MatchLabels: map[string]string{"app": c.Name, c.Name + "_cr": c.Name},
	}

	podTemplateSpec := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{"app": c.Name, c.Name + "_cr": c.Name},
		},
		Spec: corev1.PodSpec{
			ServiceAccountName: serviceAccountName,
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
	}

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: objectMeta,
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &selector,
			Template: podTemplateSpec,		
		},
	}

	daemonset := &appsv1.DaemonSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "DaemonSet",
		},
		ObjectMeta: objectMeta,
		Spec: appsv1.DaemonSetSpec{
			Selector: &selector,
			Template: podTemplateSpec,		
		},
	}

	if c.Type == "deployment"{
		existingDeployment := &appsv1.Deployment{}
		err = cl.Get(context.TODO(), types.NamespacedName{Name: "tf" + c.Name + "-" + c.InstanceName, Namespace: c.InstanceNamespace}, existingDeployment)
		if err != nil && errors.IsNotFound(err) {
			controllerutil.SetControllerReference(instance, deployment, scheme)
			err = cl.Create(context.TODO(), deployment)
			if err != nil {
				return err
			}
		}
	}
	if c.Type == "daemonset"{
		existingDaemonset := &appsv1.DaemonSet{}
		err = cl.Get(context.TODO(), types.NamespacedName{Name: "tf" + c.Name + "-" + c.InstanceName, Namespace: c.InstanceNamespace}, existingDaemonset)
		if err != nil && errors.IsNotFound(err) {
			controllerutil.SetControllerReference(instance, daemonset, scheme)
			err = cl.Create(context.TODO(), daemonset)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getResourceConfig(c *ClusterResource, cl client.Client, resourceType string) error {
	reqLogger := log.WithValues("Request.Namespace", c.InstanceNamespace, "Request.Name", c.InstanceName)
	reqLogger.Info("getting " + resourceType + " config")

	switch resourceType{
	case "cassandra":
		err = cl.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.CassandraInstance)

	case "zookeeper":
		err = cl.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.ZookeeperInstance)

	case "rabbitmq":
		err = cl.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.RabbitmqInstance)

	case "config":
		err = cl.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.ConfigInstance)

	case "control":
		err = cl.Get(context.TODO(), types.NamespacedName{Name: c.InstanceName, Namespace: c.InstanceNamespace}, c.ControlInstance)

	}

	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info(resourceType + " instance not found")
		return err
	}
	reqLogger.Info(resourceType + " instance")
	configMap := &corev1.ConfigMap{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: "tf" + resourceType + "cmv1", Namespace: c.InstanceNamespace}, configMap)
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

	case "control":
		c.ResourceConfig["CONTROL_NODES"] = configMap.Data["CONTROLLER_NODES"]
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

func (c ClusterResource) CreateRbac(cl client.Client, instance metav1.Object, scheme *runtime.Scheme) error {


	//controllerutil.SetControllerReference(instance, dep, r.scheme)

	existingServiceAccount := &corev1.ServiceAccount{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: "contrail-service-account-" + c.Name, Namespace: c.InstanceNamespace}, existingServiceAccount)
	if err != nil && errors.IsNotFound(err) {
		serviceAccount := createServiceAccount(c.Name, c.InstanceNamespace)
		controllerutil.SetControllerReference(instance, serviceAccount, scheme)
		err = cl.Create(context.TODO(), serviceAccount)
		if err != nil {
			return err
		}
	}
	
	existingClusterRole := &rbacv1.ClusterRole{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: "contrail-cluster-role-" + c.Name}, existingClusterRole)
	if err != nil && errors.IsNotFound(err) {
		clusterRole := createClusterRole(c.Name, c.InstanceNamespace)
		controllerutil.SetControllerReference(instance, clusterRole, scheme)
		err = cl.Create(context.TODO(), clusterRole)
		if err != nil {
			return err
		}
	}

	existingClusterRoleBinding := &rbacv1.ClusterRoleBinding{}
	err = cl.Get(context.TODO(), types.NamespacedName{Name: "contrail-cluster-role-binding-" + c.Name}, existingClusterRoleBinding)
	if err != nil && errors.IsNotFound(err) {
		clusterRoleBinding := createClusterRoleBinding(c.Name, c.InstanceNamespace)
		controllerutil.SetControllerReference(instance, clusterRoleBinding, scheme)
		err = cl.Create(context.TODO(), clusterRoleBinding)
		if err != nil {
			return err
		}
	}
	return nil
}

func createServiceAccount(Name string, Namespace string) *corev1.ServiceAccount {
	sa := &corev1.ServiceAccount {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-service-account-" + Name,
			Namespace: Namespace,
		},
	}
	return sa
}

func createClusterRoleBinding(Name string, Namespace string) *rbacv1.ClusterRoleBinding {
	crb := &rbacv1.ClusterRoleBinding {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac/v1",
			Kind:       "ClusterRoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-cluster-role-binding-" + Name,
			Namespace: Namespace,
		},
		Subjects: []rbacv1.Subject{{
			Kind: "ServiceAccount",
			Name: "contrail-service-account-" + Name,
			Namespace: Namespace,
			}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind: "ClusterRole",
			Name: "contrail-cluster-role-" + Name,
			},
	}
	return crb
}

func createClusterRole(Name string, Namespace string) *rbacv1.ClusterRole {
	cr := &rbacv1.ClusterRole {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "contrail-cluster-role-" + Name,
			Namespace: Namespace,
		},
		Rules: []rbacv1.PolicyRule{{
			Verbs: []string{
				"*",
			},
			APIGroups: []string{
				"*",
			},
			Resources: []string{
				"*",
			},
		}},
	}
	return cr
}
