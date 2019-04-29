package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	basev1 "github.com/michaelhenkel/tungstenfabric-operator/pkg/apis/base/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubemanagerClusterSpec defines the desired state of KubemanagerCluster
// +k8s:openapi-gen=true
type KubemanagerClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	basev1.BaseParameter `json:",inline"`
	KubeManagerImage string `json:"kubeManagerImage,omitempty"`
	NodeInitImage string `json:"nodeInitImage,omitempty"`
	StatusImage string `json:"statusImage,omitempty"`
	CassandraClusterName string `json:"cassandraClusterName,omitempty"`
	RabbitmqClusterName string `json:"rabbitmqClusterName,omitempty"`
	ConfigClusterName string `json:"configClusterName,omitempty"`
	KubernetesApiServer string `json:"kubernetesApiServer,omitempty"`
	KubernetesApiPort string `json:"kubernetesApiPort,omitempty"`
	KubernetesPodSubnets string `json:"kubernetesPodSubnets,omitempty"`
	KubernetesServicesSubnets string `json:"kubernetesServiceSubnets,omitempty"`
	KubernetesClusterName string `json:"kubernetesClusterName,omitempty"`
	KubernetesIpFabricForwarding string `json:"kubernetesIpFabricForwarding,omitempty"`
	KubernetesIpFabricSnat string `json:"kubernetesIpFabricSnat,omitempty"`
	KubernetesTokenFile string `json:"kubernetesTokenFile,omitempty"`
	ZookeeperClusterName string `json:"zookeeperClusterName,omitempty"`
}

// KubemanagerClusterStatus defines the observed state of KubemanagerCluster
// +k8s:openapi-gen=true
type KubemanagerClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubemanagerCluster is the Schema for the kubemanagerclusters API
// +k8s:openapi-gen=true
type KubemanagerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   KubemanagerClusterSpec   `json:"spec,omitempty"`
	Status KubemanagerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubemanagerClusterList contains a list of KubemanagerCluster
type KubemanagerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubemanagerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubemanagerCluster{}, &KubemanagerClusterList{})
}
