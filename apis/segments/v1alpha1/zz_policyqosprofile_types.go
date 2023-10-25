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

type EgressRateShaperObservation struct {

	// Average Bandwidth in mbps
	AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

	// Burst size in bytes
	BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	// Whether this rate shaper is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Peak Bandwidth in mbps
	PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}

type EgressRateShaperParameters struct {

	// Average Bandwidth in mbps
	// +kubebuilder:validation:Optional
	AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

	// Burst size in bytes
	// +kubebuilder:validation:Optional
	BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	// Whether this rate shaper is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Peak Bandwidth in mbps
	// +kubebuilder:validation:Optional
	PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}

type IngressBroadcastRateShaperObservation struct {

	// Average Bandwidth in kbps
	AverageBwKbps *float64 `json:"averageBwKbps,omitempty" tf:"average_bw_kbps,omitempty"`

	// Burst size in bytes
	BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	// Whether this rate shaper is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Peak Bandwidth in kbps
	PeakBwKbps *float64 `json:"peakBwKbps,omitempty" tf:"peak_bw_kbps,omitempty"`
}

type IngressBroadcastRateShaperParameters struct {

	// Average Bandwidth in kbps
	// +kubebuilder:validation:Optional
	AverageBwKbps *float64 `json:"averageBwKbps,omitempty" tf:"average_bw_kbps,omitempty"`

	// Burst size in bytes
	// +kubebuilder:validation:Optional
	BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	// Whether this rate shaper is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Peak Bandwidth in kbps
	// +kubebuilder:validation:Optional
	PeakBwKbps *float64 `json:"peakBwKbps,omitempty" tf:"peak_bw_kbps,omitempty"`
}

type IngressRateShaperObservation struct {

	// Average Bandwidth in mbps
	AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

	// Burst size in bytes
	BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	// Whether this rate shaper is enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Peak Bandwidth in mbps
	PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}

type IngressRateShaperParameters struct {

	// Average Bandwidth in mbps
	// +kubebuilder:validation:Optional
	AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

	// Burst size in bytes
	// +kubebuilder:validation:Optional
	BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

	// Whether this rate shaper is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Peak Bandwidth in mbps
	// +kubebuilder:validation:Optional
	PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}

type PolicyQosProfileContextObservation struct {

	// Id of the project which the resource belongs to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type PolicyQosProfileContextParameters struct {

	// Id of the project which the resource belongs to.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

type PolicyQosProfileObservation struct {

	// Class of service
	ClassOfService *float64 `json:"classOfService,omitempty" tf:"class_of_service,omitempty"`

	// Resource context
	Context []PolicyQosProfileContextObservation `json:"context,omitempty" tf:"context,omitempty"`

	// Description for this resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for this resource
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// DSCP Priority
	DscpPriority *float64 `json:"dscpPriority,omitempty" tf:"dscp_priority,omitempty"`

	// Trust mode for DSCP
	DscpTrusted *bool `json:"dscpTrusted,omitempty" tf:"dscp_trusted,omitempty"`

	EgressRateShaper []EgressRateShaperObservation `json:"egressRateShaper,omitempty" tf:"egress_rate_shaper,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IngressBroadcastRateShaper []IngressBroadcastRateShaperObservation `json:"ingressBroadcastRateShaper,omitempty" tf:"ingress_broadcast_rate_shaper,omitempty"`

	IngressRateShaper []IngressRateShaperObservation `json:"ingressRateShaper,omitempty" tf:"ingress_rate_shaper,omitempty"`

	// NSX ID for this resource
	NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

	// Policy path for this resource
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []PolicyQosProfileTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyQosProfileParameters struct {

	// Class of service
	// +kubebuilder:validation:Optional
	ClassOfService *float64 `json:"classOfService,omitempty" tf:"class_of_service,omitempty"`

	// Resource context
	// +kubebuilder:validation:Optional
	Context []PolicyQosProfileContextParameters `json:"context,omitempty" tf:"context,omitempty"`

	// Description for this resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for this resource
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// DSCP Priority
	// +kubebuilder:validation:Optional
	DscpPriority *float64 `json:"dscpPriority,omitempty" tf:"dscp_priority,omitempty"`

	// Trust mode for DSCP
	// +kubebuilder:validation:Optional
	DscpTrusted *bool `json:"dscpTrusted,omitempty" tf:"dscp_trusted,omitempty"`

	// +kubebuilder:validation:Optional
	EgressRateShaper []EgressRateShaperParameters `json:"egressRateShaper,omitempty" tf:"egress_rate_shaper,omitempty"`

	// +kubebuilder:validation:Optional
	IngressBroadcastRateShaper []IngressBroadcastRateShaperParameters `json:"ingressBroadcastRateShaper,omitempty" tf:"ingress_broadcast_rate_shaper,omitempty"`

	// +kubebuilder:validation:Optional
	IngressRateShaper []IngressRateShaperParameters `json:"ingressRateShaper,omitempty" tf:"ingress_rate_shaper,omitempty"`

	// NSX ID for this resource
	// +kubebuilder:validation:Optional
	NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []PolicyQosProfileTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyQosProfileTagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyQosProfileTagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyQosProfileSpec defines the desired state of PolicyQosProfile
type PolicyQosProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyQosProfileParameters `json:"forProvider"`
}

// PolicyQosProfileStatus defines the observed state of PolicyQosProfile.
type PolicyQosProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyQosProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyQosProfile is the Schema for the PolicyQosProfiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyQosProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   PolicyQosProfileSpec   `json:"spec"`
	Status PolicyQosProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyQosProfileList contains a list of PolicyQosProfiles
type PolicyQosProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyQosProfile `json:"items"`
}

// Repository type metadata.
var (
	PolicyQosProfile_Kind             = "PolicyQosProfile"
	PolicyQosProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyQosProfile_Kind}.String()
	PolicyQosProfile_KindAPIVersion   = PolicyQosProfile_Kind + "." + CRDGroupVersion.String()
	PolicyQosProfile_GroupVersionKind = CRDGroupVersion.WithKind(PolicyQosProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyQosProfile{}, &PolicyQosProfileList{})
}
