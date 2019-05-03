package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TungstenfabricConfigSpec defines the desired state of TungstenfabricConfig
// +k8s:openapi-gen=true



type General struct {
	Size string `json:"size,omitempty"`
	HostNetwork string `json:"hostNetwork,omitempty"`
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
	AaaMode string `json:"aaaMode,omitempty"`
	AuthMode string `json:"authMode,omitempty"`
	CloudAdminRole string `json:"cloudAdminRole,omitempty"`
	GlobalReadOnlyRole string `json:"globalReadOnlyRole,omitempty"`
}

/*
type Images struct {
	NodeManger string `json:"nodeManager,omitempty"`
	NodeInit string `json:"nodeInit,omitempty"`
	Status string `json:"status,omitempty"`
	Api string `json:"api,omitempty"`
	DeviceManager string `json:"deviceManager,omitempty"`
	SchemaTransfomer string `json:"schemaTransformer,omitempty"`
	ServiceMonitor string `json:"serviceMonitor,omitempty"`
	AnalyticsApi string `json:"analyticsApi,omitempty"`
	Collector string `json:"collector,omitempty"`
	Redis string `json:"redis,omitempty"`
	Control string `json:"control,omitempty"`
	Named string `json:"named,omitempty"`
	Dns string `json:"dns,omitempty"`
	KubeManager string `json:"kubeManager,omitempty"`
	Cassandra string `json:"cassandra,omitempty"`
	Zookeeper string `json:"zookeeper,omitempty"`
	RabbitMq string `json:"rabbitMq,omitempty"`
}

type KubeManagerCluster struct {
	Size int32 `json:"size,omitempty"`
	KubernetesApiServer string `json:"kubernetesApiServer,omitempty"`
	KubernetesApiPort string `json:"kubernetesApiPort,omitempty"`
	PodSubnets string `json:"podSubnets,omitempty"`
	ServiceSubnets string `json:"serviceSubnets,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	IpFabricForwarding string `json:"ipFabricForwarding,omitempty"`
	IpFabricSnat string `json:"ipFabricSnat,omitempty"`
	TokenFile string `json:"tokenFile,omitempty"`
}

type CassandraCluster struct {
	Size int32 `json:"size,omitempty"`
	KubernetesApiServer string `json:"kubernetesApiServer,omitempty"`
	KubernetesApiPort string `json:"kubernetesApiPort,omitempty"`
	PodSubnets string `json:"podSubnets,omitempty"`
	ServiceSubnets string `json:"serviceSubnets,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	IpFabricForwarding string `json:"ipFabricForwarding,omitempty"`
	IpFabricSnat string `json:"ipFabricSnat,omitempty"`
	TokenFile string `json:"tokenFile,omitempty"`
}
*/
type TungstenfabricConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Images map[string]string `json:"images,inline"`
	General map[string]string `json:"general,inline"`
	KubeManagerCluster map[string]string `json:"kubeManager,inline"`
	CassandraCluster map[string]string `json:"cassandraCluster,inline"`
	ZookeeperCluster map[string]string `json:"zookeeperCluster,inline"`
	RabbitmqCluster map[string]string `json:"rabbitmqCluster,inline"`
	ConfigCluster map[string]string `json:"configCluster,inline"`
	ControlCluster map[string]string `json:"controlCluster,inline"`
}

// TungstenfabricConfigStatus defines the observed state of TungstenfabricConfig
// +k8s:openapi-gen=true
type TungstenfabricConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TungstenfabricConfig is the Schema for the tungstenfabricconfigs API
// +k8s:openapi-gen=true
type TungstenfabricConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TungstenfabricConfigSpec   `json:"spec,omitempty"`
	Status TungstenfabricConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TungstenfabricConfigList contains a list of TungstenfabricConfig
type TungstenfabricConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TungstenfabricConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TungstenfabricConfig{}, &TungstenfabricConfigList{})
}
