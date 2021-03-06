package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigClusterSpec defines the desired state of ConfigCluster
// +k8s:openapi-gen=true
type ConfigClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Containers []*Container `json:"containers"`
	InitContainers []*Container `json:"initcontainers"`
	ConfigParameters map[string]string `json:"configparameters,omitempty"`
	Type string `json:"type,omitempty"`
	General *General `json:"general,omitempty"`
}

// ConfigClusterStatus defines the observed state of ConfigCluster
// +k8s:openapi-gen=true
type ConfigClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigCluster is the Schema for the configclusters API
// +k8s:openapi-gen=true
type ConfigCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigClusterSpec   `json:"spec,omitempty"`
	Status ConfigClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigClusterList contains a list of ConfigCluster
type ConfigClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigCluster{}, &ConfigClusterList{})
}
