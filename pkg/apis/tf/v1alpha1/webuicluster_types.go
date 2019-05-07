package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	basev1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/base/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebuiClusterSpec defines the desired state of WebuiCluster
// +k8s:openapi-gen=true

type WebuiClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Containers []*basev1.Container `json:"containers"`
	BaseParameter *basev1.BaseParameter
}

// WebuiClusterStatus defines the observed state of WebuiCluster
// +k8s:openapi-gen=true
type WebuiClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebuiCluster is the Schema for the webuiclusters API
// +k8s:openapi-gen=true
type WebuiCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebuiClusterSpec   `json:"spec,omitempty"`
	Status WebuiClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebuiClusterList contains a list of WebuiCluster
type WebuiClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebuiCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebuiCluster{}, &WebuiClusterList{})
}
