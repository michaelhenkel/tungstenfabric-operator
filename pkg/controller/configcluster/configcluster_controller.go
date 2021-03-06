package configcluster

import (
	"context"
	"reflect"
	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	//appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_configcluster")

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileConfigCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("configcluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.ConfigCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.ConfigCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileConfigCluster{}

type ReconcileConfigCluster struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileConfigCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	resourceType := "config"
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ConfigCluster")

	instance := &tfv1alpha1.ConfigCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	baseInstance := &tfv1alpha1.TungstenfabricManager{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, baseInstance)
	if err != nil && errors.IsNotFound(err){
		reqLogger.Info("baseconfig instance not found")
	}

	var configMap = make(map[string]string)
	for k,v := range(baseInstance.Spec.ConfigConfig){
		configMap[k] = v
	}
	cassandraInstance := &tfv1alpha1.CassandraCluster{}
	rabbitmqInstance := &tfv1alpha1.RabbitmqCluster{}
	zookeeperInstance := &tfv1alpha1.ZookeeperCluster{}
	var resource tfv1alpha1.TungstenFabricResource
	clusterResource := &tfv1alpha1.ClusterResource{
		Name: resourceType,
		InstanceName: instance.Name,
		InstanceNamespace: instance.Namespace,
		Containers: instance.Spec.Containers,
		General: instance.Spec.General,
		ResourceConfig: configMap,
		BaseInstance: baseInstance,
		InitContainers: instance.Spec.InitContainers,
		WaitFor: []string{"cassandra","zookeeper","rabbitmq"},
		CassandraInstance: cassandraInstance,
		ZookeeperInstance: zookeeperInstance,
		RabbitmqInstance: rabbitmqInstance,
		Type: instance.Spec.Type,
		VolumeList: map[string]bool{},
	}
	resource = clusterResource

	// Create Deployment

	err = resource.CreateDeployment(r.client, instance, r.scheme)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info(clusterResource.Name + " deployment created")

	var podNames []string
	podNames, err = resource.GetPodNames(r.client)
	if err != nil {
		reqLogger.Error(err, "Failed to get PodNames")
		return reconcile.Result{}, err
	} else {
		reqLogger.Info("Got PodNames")
	}

	if !reflect.DeepEqual(podNames, instance.Status.Nodes) {
		instance.Status.Nodes = podNames
		err = r.client.Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update Pod status.")
			return reconcile.Result{}, err
		}
	}
	reqLogger.Info("Updated Node status with PodNames")

	var initContainerRunning bool
	initContainerRunning, err = resource.WaitForInitContainer(r.client)
	if err != nil || !initContainerRunning{
		reqLogger.Info("Init container not running")
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("Init Container running")

	clusterResource.ResourceConfig["CONTROLLER_NODES"] = resource.GetNodeIpList()
	clusterResource.ResourceConfig["CONFIG_NODES"] = resource.GetNodeIpList()

	// Create ConfigMap
	err = resource.CreateConfigMap(r.client, instance, r.scheme)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}

	reqLogger.Info("Config configmap created")

	var labeledPod *corev1.Pod
	for _, pod := range(podNames){
		labeledPod, err = resource.LabelPod(r.client, pod)
		if err != nil {
			return reconcile.Result{}, err
		}
			err = r.client.Update(context.TODO(), labeledPod)
		if err != nil {
			reqLogger.Error(err, "Failed to update Pod label.")
			return reconcile.Result{}, err
		}
		reqLogger.Info("Labeled Pod")
	}

	return reconcile.Result{}, nil
}
