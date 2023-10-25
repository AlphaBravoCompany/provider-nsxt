/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyVmTagsContextObservation struct {

	// Id of the project which the resource belongs to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type PolicyVmTagsContextParameters struct {

	// Id of the project which the resource belongs to.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

type PolicyVmTagsObservation struct {

	// Resource context
	Context []PolicyVmTagsContextObservation `json:"context,omitempty" tf:"context,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance id
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Tag specificiation for corresponding segment port
	Port []PortObservation `json:"port,omitempty" tf:"port,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []PolicyVmTagsTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyVmTagsParameters struct {

	// Resource context
	// +kubebuilder:validation:Optional
	Context []PolicyVmTagsContextParameters `json:"context,omitempty" tf:"context,omitempty"`

	// Instance id
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Tag specificiation for corresponding segment port
	// +kubebuilder:validation:Optional
	Port []PortParameters `json:"port,omitempty" tf:"port,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []PolicyVmTagsTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyVmTagsTagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyVmTagsTagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PortObservation struct {

	// Segment path where VM port should be tagged
	SegmentPath *string `json:"segmentPath,omitempty" tf:"segment_path,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []PortTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PortParameters struct {

	// Segment path where VM port should be tagged
	// +kubebuilder:validation:Required
	SegmentPath *string `json:"segmentPath" tf:"segment_path,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []PortTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PortTagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PortTagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyVmTagsSpec defines the desired state of PolicyVmTags
type PolicyVmTagsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyVmTagsParameters `json:"forProvider"`
}

// PolicyVmTagsStatus defines the observed state of PolicyVmTags.
type PolicyVmTagsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyVmTagsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyVmTags is the Schema for the PolicyVmTagss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyVmTags struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceId)",message="instanceId is a required parameter"
	Spec   PolicyVmTagsSpec   `json:"spec"`
	Status PolicyVmTagsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyVmTagsList contains a list of PolicyVmTagss
type PolicyVmTagsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyVmTags `json:"items"`
}

// Repository type metadata.
var (
	PolicyVmTags_Kind             = "PolicyVmTags"
	PolicyVmTags_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyVmTags_Kind}.String()
	PolicyVmTags_KindAPIVersion   = PolicyVmTags_Kind + "." + CRDGroupVersion.String()
	PolicyVmTags_GroupVersionKind = CRDGroupVersion.WithKind(PolicyVmTags_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyVmTags{}, &PolicyVmTagsList{})
}
