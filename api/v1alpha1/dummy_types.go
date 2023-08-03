// Copyright 2023 Jimmi Dyson.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DummySpec defines the desired state of Dummy.
type DummySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Dummy. Edit dummy_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DummyStatus defines the observed state of Dummy.
type DummyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Dummy is the Schema for the dummies API.
type Dummy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DummySpec   `json:"spec,omitempty"`
	Status DummyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DummyList contains a list of Dummy.
type DummyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dummy `json:"items"`
}

//nolint:gochecknoinits // Idiomatic use of init in kubebuilder projects.
func init() {
	SchemeBuilder.Register(&Dummy{}, &DummyList{})
}
