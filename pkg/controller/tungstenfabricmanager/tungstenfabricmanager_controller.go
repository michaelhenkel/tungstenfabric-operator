package tungstenfabricmanager

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
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
/*



	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/vrouter"
*/
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/kubemanagercluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/controlcluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/configcluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/rabbitmqcluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/zookeepercluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/cassandracluster"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/webuicluster"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
)

var log = logf.Log.WithName("controller_tungstenfabricmanager")

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	apiextensionsv1beta1.AddToScheme(scheme.Scheme)
	return &ReconcileTungstenfabricManager{client: mgr.GetClient(), scheme: mgr.GetScheme(), manager: mgr}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("tungstenfabricmanager-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &tfv1alpha1.TungstenfabricManager{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.TungstenfabricManager{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTungstenfabricManager{}
type ReconcileTungstenfabricManager struct {

	client client.Client
	scheme *runtime.Scheme
	manager manager.Manager
}


func (r *ReconcileTungstenfabricManager) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling TungstenfabricManager")

	instance := &tfv1alpha1.TungstenfabricManager{}
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
/*












	functionMap["vrouter"] = vrouter.Add
	crdMap["Vrouter"] = functionMap
*/

	functionMap["kubemanagercluster"] = kubemanagercluster.Add
	crdMap["KubemanagerCluster"] = functionMap

	functionMap["controlcluster"] = controlcluster.Add
	crdMap["ControlCluster"] = functionMap

	functionMap["configcluster"] = configcluster.Add
	crdMap["ConfigCluster"] = functionMap

	functionMap["rabbitmqcluster"] = rabbitmqcluster.Add
	crdMap["RabbitmqCluster"] = functionMap

	functionMap["zookeepercluster"] = zookeepercluster.Add
	crdMap["ZookeeperCluster"] = functionMap

	functionMap["cassandracluster"] = cassandracluster.Add
	crdMap["CassandraCluster"] = functionMap
	
	functionMap["webuicluster"] = webuicluster.Add
	crdMap["WebuiCluster"] = functionMap

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
/*
    m := map[string]func(string) string{
        "foo": func(s string) string { return s + "nurf" },
    }

    m["foo"]("baz") // "baznurf"
*/
	


	for _, resource := range(instance.Spec.StartResources){
		switch resource{
		case "CassandraCluster":
			err = r.CassandraCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "ZookeeperCluster":
			err = r.ZookeeperCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "RabbitmqCluster":
			err = r.RabbitmqCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "ConfigCluster":
			err = r.ConfigCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "ControlCluster":
			err = r.ControlCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "KubemanagerCluster":
			err = r.KubemanagerCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "WebuiCluster":
			err = r.WebuiCluster(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		case "Vrouter":
			err = r.Vrouter(instance.Name, instance.Namespace)
			if err != nil {
				reqLogger.Error(err, "Failed to create resource " + resource)
				return reconcile.Result{}, err
			}
		}
	}

/*
	var instanceMap = make(map[string]metav1.Object)

	instanceMap["ConfigCluster"] = &tfv1alpha1.ConfigCluster{}
	instanceMap["ControlCluster"] = &tfv1alpha1.ControlCluster{}
	instanceMap["ZookeeperCluster"] = &tfv1alpha1.ZookeeperCluster{}
	instanceMap["CassandraCluster"] = &tfv1alpha1.CassandraCluster{}
	instanceMap["RabbitmqCluster"] = &tfv1alpha1.RabbitmqCluster{}
	instanceMap["KubemanagerCluster"] = &tfv1alpha1.KubemanagerCluster{}
	instanceMap["WebuiCluster"] = &tfv1alpha1.WebuiCluster{}


	//instanceMap["Vrouter"] = tfv1alpha1.Vrouter{}
	//bla := tfv1alpha1.ConfigCluster{}


	var cr metav1.Object
	var clusterResource metav1.Object
	for _, resource := range(instance.Spec.StartResources){
		cr = instanceMap[resource]
		err := r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, cr)
		
			clusterResource = r.CreateConfigResource(instance.Name, instance.Namespace, resource)
			err = r.client.Create(context.TODO(), &clusterResource)
			if err != nil {
				return reconcile.Result{}, err
			}
		} else if err != nil {
			return  reconcile.Result{}, err
		}
	}
*/

	return  reconcile.Result{},nil
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}