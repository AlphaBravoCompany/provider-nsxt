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

type ClusterVirtualIpObservation struct {

	// On enable it ignores duplicate address detection and DNS lookup validation check
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Virtual IPv4 address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Virtual IPv6 address
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`
}

type ClusterVirtualIpParameters struct {

	// On enable it ignores duplicate address detection and DNS lookup validation check
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// Virtual IPv4 address
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Virtual IPv6 address
	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`
}

// ClusterVirtualIpSpec defines the desired state of ClusterVirtualIp
type ClusterVirtualIpSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterVirtualIpParameters `json:"forProvider"`
}

// ClusterVirtualIpStatus defines the observed state of ClusterVirtualIp.
type ClusterVirtualIpStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterVirtualIpObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterVirtualIp is the Schema for the ClusterVirtualIps API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type ClusterVirtualIp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterVirtualIpSpec   `json:"spec"`
	Status            ClusterVirtualIpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterVirtualIpList contains a list of ClusterVirtualIps
type ClusterVirtualIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVirtualIp `json:"items"`
}

// Repository type metadata.
var (
	ClusterVirtualIp_Kind             = "ClusterVirtualIp"
	ClusterVirtualIp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterVirtualIp_Kind}.String()
	ClusterVirtualIp_KindAPIVersion   = ClusterVirtualIp_Kind + "." + CRDGroupVersion.String()
	ClusterVirtualIp_GroupVersionKind = CRDGroupVersion.WithKind(ClusterVirtualIp_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterVirtualIp{}, &ClusterVirtualIpList{})
}
