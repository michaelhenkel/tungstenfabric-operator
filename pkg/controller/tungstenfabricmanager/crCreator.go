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
			},{
				Name: "devicemanager",
				LogVolumePath: "/var/log/contrail",
			},{
				Name: "schematransformer",
				LogVolumePath: "/var/log/contrail",
			},{
				Name: "servicemonitor",
				LogVolumePath: "/var/log/contrail",
			},{
				Name: "analyticsapi",
				LogVolumePath: "/var/log/contrail",
			},{
				Name: "collector",
				LogVolumePath: "/var/log/contrail",
			},{
				Name: "redis",
				LogVolumePath: "/var/log/contrail",
				DataVolumePath: "/var/lib/redis",
			},{
				Name: "nodemanagerconfig",
				LogVolumePath: "/var/log/contrail",
				UnixSocketVolume: true,
				Env: map[string]string{
					"NODE_TYPE": "config",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
			},{
				Name: "nodemanageranalytics",
				LogVolumePath: "/var/log/contrail",
				Env: map[string]string{
					"NODE_TYPE": "analytics",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
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
			},{
				Name: "dns",
				LogVolumePath: "/var/log/contrail",
				EtcContrailVolume: true,
			},{
				Name: "named",
				LogVolumePath: "/var/log/contrail",
				Privileged: true,
				EtcContrailVolume: true,
			},{
				Name: "nodemanager",
				LogVolumePath: "/var/log/contrail",
				UnixSocketVolume: true,
				Env: map[string]string{
					"NODE_TYPE": "control",
					"DOCKER_HOST": "unix://mnt/docker.sock",
				},
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
			},
			{
				Name: "webuijob",
				LogVolumePath: "/var/log/contrail",
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
				Name: "agent",
				LogVolumePath: "/var/log/contrail",
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