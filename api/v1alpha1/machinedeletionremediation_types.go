/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MachineDeletionOnCloudProviderReason     = "MachineDeletionOnCloudProviderCausesNewNodeName"
	MachineDeletionOnBareMetalProviderReason = "MachineDeletionOnBareMetalProviderKeepsNodeName"
	MachineDeletionOnUndefinedProviderReason = "MachineDeletionUndefinedNodeNameExpectation"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MachineDeletionRemediationSpec defines the desired state of MachineDeletionRemediation
type MachineDeletionRemediationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

// MachineDeletionRemediationStatus defines the observed state of MachineDeletionRemediation
type MachineDeletionRemediationStatus struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="conditions",xDescriptors="urn:alm:descriptor:io.kubernetes.conditions"
	// Represents the observations of a MachineDeletionRemediation's current state.
	// Known .status.conditions.type are: "Processing", "Succeeded" and "PermanentNodeDeletionExpected"
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=mdr

// MachineDeletionRemediation is the Schema for the machinedeletionremediations API
// +operator-sdk:csv:customresourcedefinitions:resources={{"MachineDeletionRemediation","v1alpha1","machinedeletionremediations"}}
type MachineDeletionRemediation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineDeletionRemediationSpec   `json:"spec,omitempty"`
	Status MachineDeletionRemediationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MachineDeletionRemediationList contains a list of MachineDeletionRemediation
type MachineDeletionRemediationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineDeletionRemediation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MachineDeletionRemediation{}, &MachineDeletionRemediationList{})
}
