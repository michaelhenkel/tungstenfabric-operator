package tungstenfabricconfig

import (
	"context"
	"strings"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"

	corev1 "k8s.io/api/core/v1"
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
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/cassandracluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/zookeepercluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/rabbitmqcluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/configcluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/controlcluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/kubemanagercluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/vrouter"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
)

var log = logf.Log.WithName("controller_tungstenfabricconfig")

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	apiextensionsv1beta1.AddToScheme(scheme.Scheme)
	return &ReconcileTungstenfabricConfig{client: mgr.GetClient(), scheme: mgr.GetScheme(), manager: mgr}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("tungstenfabricconfig-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.TungstenfabricConfig{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.TungstenfabricConfig{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTungstenfabricConfig{}
type ReconcileTungstenfabricConfig struct {

	client client.Client
	scheme *runtime.Scheme
	manager manager.Manager
}


func (r *ReconcileTungstenfabricConfig) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling TungstenfabricConfig")

	instance := &tfv1alpha1.TungstenfabricConfig{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}


	type AddFunction func(manager.Manager) error

	var functionMap map[string]AddFunction
	functionMap = make(map[string]AddFunction)



	var crdMap map[string]map[string]AddFunction
	crdMap = make(map[string]map[string]AddFunction)
	
	crdVersion := "v1alpha1"
	crdGroup := "tf.tungstenfabric.io"

	functionMap["cassandracluster"] = cassandracluster.Add
	crdMap["CassandraCluster"] = functionMap

	functionMap["zookeepercluster"] = zookeepercluster.Add
	crdMap["ZookeeperCluster"] = functionMap

	functionMap["rabbitmqcluster"] = rabbitmqcluster.Add
	crdMap["RabbitmqCluster"] = functionMap

	functionMap["configcluster"] = configcluster.Add
	crdMap["ConfigCluster"] = functionMap

	functionMap["controlcluster"] = controlcluster.Add
	crdMap["ControlCluster"] = functionMap

	functionMap["kubemanagercluster"] = kubemanagercluster.Add
	crdMap["KubemanagerCluster"] = functionMap

	functionMap["vrouter"] = vrouter.Add
	crdMap["Vrouter"] = functionMap

	for crdName, crdFunction := range crdMap{
		singular := strings.ToLower(crdName)
		plural := singular + "s"
		newCrd := apiextensionsv1beta1.CustomResourceDefinition{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: plural + "." + crdGroup, Namespace: newCrd.Namespace}, &newCrd)
		if err != nil && errors.IsNotFound(err) {
			crd := r.CreateCrd(crdName,crdVersion,crdGroup, newCrd.Namespace)
			reqLogger.Info("Creating a new crd.", "crd.Namespace", crd.Namespace, "crd.Name", crd.Name)
			err = r.client.Create(context.TODO(), crd)
			if err != nil {
				reqLogger.Error(err, "Failed to create new crd.", "crd.Namespace", crd.Namespace, "crd.Name", crd.Name)
				return reconcile.Result{}, err
			}
		} else if err != nil {
			reqLogger.Error(err, "Failed to get CRD.")
			return reconcile.Result{}, err
		}
		
		err = crdFunction[strings.ToLower(crdName)](r.manager)
		if err != nil {
			reqLogger.Error(err, "Failed to run function.")
			return reconcile.Result{}, err
		}
		
	}
	return reconcile.Result{}, nil
}
