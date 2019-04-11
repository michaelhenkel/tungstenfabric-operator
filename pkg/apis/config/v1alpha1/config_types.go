package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        basev1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/base/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigSpec defines the desired state of Config
// +k8s:openapi-gen=true
type ConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	basev1.BaseParameter `json:",inline"`
	ApiImage string `json:"apiImage,omitempty"`
	DeviceManagerImage string `json:"deviceManagerImage,omitempty"`
	SchemaTransformerImage string `json:"schemaTransformerImage,omitempty"`
	ServiceMonitorImage string `json:"serviceMonitorImage,omitempty"`
	NodeManagerImage string `json:"nodeManagerImage,omitempty"`
	NodeInitImage string `json:"nodeInitImage,omitempty"`
	StatusImage string `json:"statusImage,omitempty"`
	ApiPort string `json:"apiport,omitempty"`
	ApiIntrospectionPort string `json:"apiintrospectionport,omitempty"`
	AuthMode string `json:"authmode,omitempty"`
	AaaMode string `json:"aaamode,omitempty"`
	CloudAdminRole string `json:"cloudadminrole,omitempty"`
	GlobalReadOnlyRole string `json:"globalreadonlyrole,omitempty"`
}

// ConfigStatus defines the observed state of Config
// +k8s:openapi-gen=true
type ConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Config is the Schema for the configs API
// +k8s:openapi-gen=true
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
