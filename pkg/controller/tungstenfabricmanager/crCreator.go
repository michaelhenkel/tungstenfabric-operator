package tungstenfabricmanager

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *ReconcileTungstenfabricManager) ConfigCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.ConfigClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "api",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "devicemanager",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "schematransformer",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "servicemonitor",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "analyticsapi",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "collector",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "redis",
				LogVolumePath: "/var/log/contrail",
				DataVolumePath: "/var/lib/redis",
				ResourceConfigMap: true,
			},{
				Name: "nodemanagerconfig",
				LogVolumePath: "/var/log/contrail",
				UnixSocketVolume: true,
				Env: map[string]string{
					"NODE_TYPE": "config",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
				ResourceConfigMap: true,
			},{
				Name: "nodemanageranalytics",
				LogVolumePath: "/var/log/contrail",
				UnixSocketVolume: true,
				Env: map[string]string{
					"NODE_TYPE": "analytics",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "init",
				StatusVolume: true,
				Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
			},{
				Name: "nodeinit",
				Privileged: true,
				HostUserBinVolume: true,
				ResourceConfigMap: true,
			},
		},
	}

	clusterResource := &tfv1alpha1.ConfigCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.ConfigCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) ControlCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.ControlClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "control",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},{
				Name: "dns",
				LogVolumePath: "/var/log/contrail",
				EtcContrailVolume: true,
				ResourceConfigMap: true,
			},{
				Name: "named",
				LogVolumePath: "/var/log/contrail",
				Privileged: true,
				EtcContrailVolume: true,
				ResourceConfigMap: true,
			},{
				Name: "nodemanager",
				LogVolumePath: "/var/log/contrail",
				UnixSocketVolume: true,
				Env: map[string]string{
					"NODE_TYPE": "control",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "init",
				StatusVolume: true,
				Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
			},{
				Name: "nodeinit",
				Privileged: true,
				HostUserBinVolume: true,
				ResourceConfigMap: true,
			},
		},
	}

	clusterResource := &tfv1alpha1.ControlCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.ControlCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) KubemanagerCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.KubemanagerClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "kubemanager",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "init",
				StatusVolume: true,
				Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
			},{
				Name: "nodeinit",
				Privileged: true,
				HostUserBinVolume: true,
				ResourceConfigMap: true,
			},
		},
	}

	clusterResource := &tfv1alpha1.KubemanagerCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.KubemanagerCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) WebuiCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.WebuiClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "webuiweb",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},
			{
				Name: "webuijob",
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "nodeinit",
				Privileged: true,
				HostUserBinVolume: true,
				ResourceConfigMap: true,
			},
		},
	}

	clusterResource := &tfv1alpha1.WebuiCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.WebuiCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) Vrouter(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.VrouterSpec{
		Type: "daemonset",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "vrouteragent",
				Privileged: true,
				LifeCycleScript: []string{"/cleanup.sh"},
				LogVolumePath: "/var/log/contrail",
				DevVolume: true,
				NetworkScriptsVolume: true,
				HostBinVolume: true,
				UsrSrcVolume: true,
				LibModulesVolume: true,
				VarLibContrailVolume: true,
				VarCrashesVolume: true,
				ResourceConfigMap: true,
			},{
				Name: "nodemanager",
				LogVolumePath: "/var/log/contrail",
				UnixSocketVolume: true,
				Env: map[string]string{
					"NODE_TYPE": "vrouter",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "nodeinit",
				Privileged: true,
				HostUserBinVolume: true,
				ResourceConfigMap: true,
			},{
				Name: "vrouterkernelinit",
				Privileged: true,
				HostUserBinVolume: true,
				UsrSrcVolume: true,
				LibModulesVolume: true,
				NetworkScriptsVolume: true,
				HostBinVolume: true,
				ResourceConfigMap: true,
			},{
				Name: "vrouternicinit",
				Privileged: true,
				HostUserBinVolume: true,
				UsrSrcVolume: true,
				LibModulesVolume: true,
				NetworkScriptsVolume: true,
				HostBinVolume: true,
				ResourceConfigMap: true,
			},{
				Name: "vroutercni",
				Privileged: true,
				HostUserBinVolume: true,
				VarLibContrailVolume: true,
				EtcCniVolume: true,
				OptBinCniVolume: true,
				VarLogCniVolume: true,
				LogVolumePath: "/var/log/contrail",
				ResourceConfigMap: true,
			},
		},
	}

	clusterResource := &tfv1alpha1.Vrouter{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.Vrouter{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) CassandraCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.CassandraClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "cassandra",
				LogVolumePath: "/var/log/cassandra",
				DataVolumePath: "/var/lib/cassandra",
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "init",
				StatusVolume: true,
				Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
			},
		},
	}

	clusterResource := &tfv1alpha1.CassandraCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.CassandraCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) ZookeeperCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.ZookeeperClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "zookeeper",
				LogVolumePath: "/var/log/zookeeper",
				DataVolumePath: "/var/lib/zookeeper",
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "init",
				StatusVolume: true,
				Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
			},
		},
	}

	clusterResource := &tfv1alpha1.ZookeeperCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.ZookeeperCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}

func (r *ReconcileTungstenfabricManager) RabbitmqCluster(
	instanceName string,
	instanceNamespace string) error {

	resourceSpec := tfv1alpha1.RabbitmqClusterSpec{
		Type: "deployment",
		Containers: []*tfv1alpha1.Container{
			{
				Name: "rabbitmq",
				LogVolumePath: "/var/log/rabbitmq",
				DataVolumePath: "/var/lib/rabbitmq",
				ResourceConfigMap: true,
			},
		},
		InitContainers: []*tfv1alpha1.Container{
			{
				Name: "init",
				StatusVolume: true,
				Command: []string{"sh","-c","until grep ready /tmp/podinfo/pod_labels > /dev/null 2>&1; do sleep 1; done"},
			},
		},
	}

	clusterResource := &tfv1alpha1.RabbitmqCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: instanceName,
			Namespace: instanceNamespace,
		},
		Spec: resourceSpec,
	}

	resource := tfv1alpha1.RabbitmqCluster{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: instanceNamespace}, &resource)
	if err != nil && errors.IsNotFound(err) {
		err = r.client.Create(context.TODO(), clusterResource)
		if err != nil{
			return err
		}
	}
	return nil
}