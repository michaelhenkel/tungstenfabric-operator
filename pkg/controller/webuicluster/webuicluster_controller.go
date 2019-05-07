package webuicluster

import (
	"context"

	tfv1alpha1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/tf/v1alpha1"
//	basev1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/base/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	"fmt"
)

var log = logf.Log.WithName("controller_webuicluster")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new WebuiCluster Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileWebuiCluster{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("webuicluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource WebuiCluster
	err = c.Watch(&source.Kind{Type: &tfv1alpha1.WebuiCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner WebuiCluster
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

// ReconcileWebuiCluster reconciles a WebuiCluster object
type ReconcileWebuiCluster struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}


// Reconcile reads that state of the cluster for a WebuiCluster object and makes changes based on the state read
// and what is in the WebuiCluster.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.

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
	for k,v := range(baseInstance.Spec.WebuiCluster){
		configMap[k] = v
	}

	cassandraInstance := &tfv1alpha1.CassandraCluster{}
	var webuiResource tfv1alpha1.TungstenFabricResource
	webuiResource = tfv1alpha1.ClusterResource{
		Name: instance.Name,
		Namespace: instance.Namespace,
		Containers: instance.Spec.Containers,
		BaseParameter: instance.Spec.BaseParameter,
		ResourceConfig: configMap,
		BaseInstance: baseInstance,
		WaitForCassandra: true,
		CassandraInstance: cassandraInstance,
	}

	cm, err := webuiResource.CreateConfigMap(r.client)
	if err != nil {
		return reconcile.Result{}, err
	}
	controllerutil.SetControllerReference(instance, cm, r.scheme)
	for _, container := range(instance.Spec.Containers){
		fmt.Println(container.Name, container.Image)
	}


/*
	webuiCustomDeployment 

	configMap, baseInstance := tfv1alpha1.GetBaseParameter(r.client, instance.Name, instance.Namespace)

	for k,v := range(baseInstance.Spec.WebuiCluster){
		configMap[k] = v
	}

	if instance.Spec.Size == "" {
		if baseInstance.Spec.WebuiCluster["size"] != "" {
			instance.Spec.Size = baseInstance.Spec.WebuiCluster["size"]
		}
	}

	if instance.Spec.WebuiWebImage == "" {
		instance.Spec.WebuiWebImage = baseInstance.Spec.Images["webuiWeb"]
	}
	if instance.Spec.WebuiJobImage == "" {
		instance.Spec.WebuiJobImage = baseInstance.Spec.Images["webuiJob"]
	}

	pullPolicy := corev1.PullAlways
	if instance.Spec.ImagePullPolicy == "Never" {
		pullPolicy = corev1.PullNever
	}
	if instance.Spec.ImagePullPolicy == "IfNotPresent" {
		pullPolicy = corev1.PullNever
	}

	webuiContainerList := []corev1.Container{
		{
			Image:   instance.Spec.WebuiWebImage,
			Name:    "webui-web",
			ImagePullPolicy: pullPolicy,
			EnvFrom: []corev1.EnvFromSource{{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "tfwebuicmv1",
					},
				},
			}},
			VolumeMounts: []corev1.VolumeMount{{
				Name: "webui-logs",
				MountPath: "/var/log/contrail",
			}},
		},
		{
			Image:   instance.Spec.WebuiJobImage,
			Name:    "webui-job",
			ImagePullPolicy: pullPolicy,
			EnvFrom: []corev1.EnvFromSource{{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "tfwebuicmv1",
					},
				},
			}},
			VolumeMounts: []corev1.VolumeMount{{
				Name: "webui-logs",
				MountPath: "/var/log/contrail",

			}},
		},
	}

	webuiCustomDeployment := CustomDeployment{
		Size: instance.Spec.Size,
		InstanceName: instance.Name,
		Name: "webui",
		InitContainerList: []*corev1.Container,
		ContainerList: webuiContainerList,
		VolumeList: []*corev1.Volume,
		Namespace: instance.Namespace,
		BaseInitContainer: true,
		NodeInitContainer: true,
		StatusVolume: true,
		LogVolume: true,
		HostUserBinVolume: false,
		HostNetwork: "true",
		Labels: map[string]string{"app": "webui", "webui_cr": "webui"},
		ServiceAccountName: "",
		PullPolicy: m.Spec.ImagePullPolicy,
	}

*/
	return reconcile.Result{}, nil
}
