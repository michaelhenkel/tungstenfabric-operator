package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TungstenfabricConfigSpec defines the desired state of TungstenfabricConfig
// +k8s:openapi-gen=true

type TungstenfabricConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Images map[string]string `json:"images,inline"`
	General *General `json:"general,inline"`

	KubemanagerConfig map[string]string `json:"kubemanagerConfig,inline"`
	CassandraConfig map[string]string `json:"cassandraConfig,inline"`
	ZookeeperConfig map[string]string `json:"zookeeperConfig,inline"`
	RabbitmqConfig map[string]string `json:"rabbitmqConfig,inline"`
	ConfigConfig map[string]string `json:"configConfig,inline"`
	ControlConfig map[string]string `json:"controlConfig,inline"`
	WebuiConfig map[string]string `json:"webuiConfig,inline"`
	VrouterConfig map[string]string `json:"vrouterConfig,inline"`

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
