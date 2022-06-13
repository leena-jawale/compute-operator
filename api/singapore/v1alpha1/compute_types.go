// Copyright Red Hat

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ComputeSpec defines the desired state of ClusterRegistrar
type ComputeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make generate" to regenerate code after modifying this file

}

// ComputeStatus defines the observed state of ClusterRegistrar
type ComputeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make generate" to regenerate code after modifying this file

	// Conditions contains the different condition statuses for this ClusterRegistrar.
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Compute is the Schema for the clusterregistrars API. Compute is a cluster scoped resource.
type Compute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSpec   `json:"spec,omitempty"`
	Status ComputeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeList contains a list of ClusterRegistrar
type ComputeList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of ClusterRegistrar.
	// +listType=set
	Items []Compute `json:"items"`
}