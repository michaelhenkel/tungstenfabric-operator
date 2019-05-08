package webuicluster

import (
	"context"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
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

var log = logf.Log.WithName("controller_webuicluster")

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileWebuiCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("webuicluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.WebuiCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.WebuiCluster{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileWebuiCluster{}

type ReconcileWebuiCluster struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileWebuiCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling WebuiCluster")

	instance := &tfv1alpha1.WebuiCluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	baseInstance := &tfv1alpha1.TungstenfabricConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, baseInstance)
	if err != nil && errors.IsNotFound(err){
		reqLogger.Info("baseconfig instance not found")
	}

	var configMap = make(map[string]string)
	for k,v := range(baseInstance.Spec.WebuiConfig){
		configMap[k] = v
	}

	cassandraInstance := &tfv1alpha1.CassandraCluster{}
	var resource tfv1alpha1.TungstenFabricResource
	resource = &tfv1alpha1.ClusterResource{
		Name: "webui",
		InstanceName: instance.Name,
		InstanceNamespace: instance.Namespace,
		Containers: instance.Spec.Containers,
		General: instance.Spec.General,
		ResourceConfig: configMap,
		BaseInstance: baseInstance,
		//WaitFor: []string{"cassandracluster"},
		CassandraInstance: cassandraInstance,
		StatusVolume: true,
		LogVolume: true,
		InitContainer: true,
	}


	// Create ConfigMap
	cm, err := resource.CreateConfigMap(r.client)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}
	controllerutil.SetControllerReference(instance, cm, r.scheme)
	err = r.client.Create(context.TODO(), cm)
	if err != nil && errors.IsAlreadyExists(err){
		err = r.client.Update(context.TODO(), cm)
	} else if err != nil {
		return reconcile.Result{}, err		
	}
	reqLogger.Info("Webui configmap created")
	
	// Create Deployment
	dep, err := resource.CreateDeployment(r.client)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}
	controllerutil.SetControllerReference(instance, dep, r.scheme)
	err = r.client.Create(context.TODO(), dep)
	if err != nil && errors.IsAlreadyExists(err){
		err = r.client.Update(context.TODO(), dep)
	} else if err != nil {
		return reconcile.Result{}, err		
	}
	reqLogger.Info("Webui deployment created")

	return reconcile.Result{}, nil
}
