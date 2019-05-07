package v1alpha1

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	basev1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/base/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)
var log = logf.Log.WithName("TungstenFabricResource")

type TungstenFabricResource interface{
	CreateConfigMap(client.Client) (*corev1.ConfigMap, error)
//	CreateDeployment() *appsv1.Deployment
}

type ClusterResource struct {
	Name string
	Namespace string
	Containers []*basev1.Container
	BaseParameter *basev1.BaseParameter
	ResourceConfig map[string]string
	BaseInstance *TungstenfabricConfig
	WaitForCassandra bool
	WaitForZookeeper bool
	WaitForRabbitmq bool
	WaitForConfig bool
	WaitForControl bool
	CassandraInstance *CassandraCluster
}

func (c ClusterResource) CreateConfigMap(client client.Client) (*corev1.ConfigMap, error) {
	var err error
	for _, container := range(c.Containers){
		if container.Image == "" {
			container.Image = c.BaseInstance.Spec.Images[container.Name]
		}
	}
	if c.WaitForCassandra {
		err = getCassandraConfig(&c, client)
		if err != nil {
			return nil, err
		}
	}

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: c.Name,
			Namespace: c.Namespace,
		},
		Data: c.ResourceConfig,
	}

	err = client.Create(context.TODO(), cm)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func getCassandraConfig(c *ClusterResource, client client.Client) error {
	reqLogger := log.WithValues("Request.Namespace", c.Namespace, "Request.Name", c.Name)
	reqLogger.Info("getting cassandra config")

	err := client.Get(context.TODO(), types.NamespacedName{Name: c.Name, Namespace: c.Namespace}, c.CassandraInstance)

	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("cassandra instance not found")
		return err
	}
	reqLogger.Info("cassandra instance")
	cassandraConfigMap := &corev1.ConfigMap{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "tfcassandraclustercmv1", Namespace: c.Namespace}, cassandraConfigMap)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("cassandra configmap not found")
		return err
	}

	c.ResourceConfig["CONFIGDB_PORT"] = cassandraConfigMap.Data["CASSANDRA_PORT"]
	c.ResourceConfig["CONFIGDB_CQL_PORT"] = cassandraConfigMap.Data["CASSANDRA_CQL_PORT"]
	c.ResourceConfig["CONFIGDB_NODES"] = cassandraConfigMap.Data["CASSANDRA_SEEDS"]

	return nil
}