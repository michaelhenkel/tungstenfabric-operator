package vrouter

import (
	"context"
	"reflect"
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
)

var log = logf.Log.WithName("controller_vrouter")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Vrouter Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileVrouter{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("vrouter-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Vrouter
	err = c.Watch(&source.Kind{Type: &tfv1alpha1.Vrouter{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Vrouter
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &tfv1alpha1.Vrouter{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileVrouter{}

// ReconcileVrouter reconciles a Vrouter object
type ReconcileVrouter struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Vrouter object and makes changes based on the state read
// and what is in the Vrouter.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileVrouter) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Vrouter")

	// Fetch the Vrouter instance
	instance := &tfv1alpha1.Vrouter{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	baseInstance := &tfv1alpha1.TungstenfabricManager{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, baseInstance)
	if err != nil && errors.IsNotFound(err){
		reqLogger.Info("baseconfig instance not found")
	}

	var configMap = make(map[string]string)
	for k,v := range(baseInstance.Spec.ControlConfig){
		configMap[k] = v
	}
	configInstance := &tfv1alpha1.ConfigCluster{}
	controlInstance := &tfv1alpha1.ControlCluster{}
	var resource tfv1alpha1.TungstenFabricResource
	clusterResource := &tfv1alpha1.ClusterResource{
		Name: "vrouter",
		InstanceName: instance.Name,
		InstanceNamespace: instance.Namespace,
		Containers: instance.Spec.Containers,
		General: instance.Spec.General,
		ResourceConfig: configMap,
		BaseInstance: baseInstance,
		InitContainers: instance.Spec.InitContainers,
		WaitFor: []string{"control","config"},
		ControlInstance: controlInstance,
		ConfigInstance: configInstance,
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
	clusterResource.ResourceConfig["CONTROL_NODES"] = resource.GetNodeIpList()

	// Create ConfigMap
	err = resource.CreateConfigMap(r.client, instance, r.scheme)
	if err != nil {
		return reconcile.Result{Requeue: true}, nil
	}
	reqLogger.Info("Vrouter configmap created")

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
