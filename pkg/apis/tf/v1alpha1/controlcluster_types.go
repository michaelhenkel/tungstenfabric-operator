package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ControlClusterSpec defines the desired state of ControlCluster
// +k8s:openapi-gen=true
type ControlClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Containers []*Container `json:"containers"`
	InitContainers []*Container `json:"initcontainers"`
	ConfigParameters map[string]string `json:"configparameters,omitempty"`
	Type string `json:"type,omitempty"`
	General *General `json:"general,omitempty"`
}

// ControlClusterStatus defines the observed state of ControlCluster
// +k8s:openapi-gen=true
type ControlClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControlCluster is the Schema for the controlclusters API
// +k8s:openapi-gen=true
type ControlCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControlClusterSpec   `json:"spec,omitempty"`
	Status ControlClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControlClusterList contains a list of ControlCluster
type ControlClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ControlCluster{}, &ControlClusterList{})
}
