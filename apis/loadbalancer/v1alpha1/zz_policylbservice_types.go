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

type PolicyLbServiceObservation struct {

	// Policy path for connected policy object
	ConnectivityPath *string `json:"connectivityPath,omitempty" tf:"connectivity_path,omitempty"`

	// Description for this resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for this resource
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Flag to enable the Service
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Log level for Load Balancer Service messages
	ErrorLogLevel *string `json:"errorLogLevel,omitempty" tf:"error_log_level,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// NSX ID for this resource
	NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

	// Policy path for this resource
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Load Balancer Service size
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []PolicyLbServiceTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyLbServiceParameters struct {

	// Policy path for connected policy object
	// +kubebuilder:validation:Optional
	ConnectivityPath *string `json:"connectivityPath,omitempty" tf:"connectivity_path,omitempty"`

	// Description for this resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for this resource
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Flag to enable the Service
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Log level for Load Balancer Service messages
	// +kubebuilder:validation:Optional
	ErrorLogLevel *string `json:"errorLogLevel,omitempty" tf:"error_log_level,omitempty"`

	// NSX ID for this resource
	// +kubebuilder:validation:Optional
	NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

	// Load Balancer Service size
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []PolicyLbServiceTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyLbServiceTagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyLbServiceTagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyLbServiceSpec defines the desired state of PolicyLbService
type PolicyLbServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyLbServiceParameters `json:"forProvider"`
}

// PolicyLbServiceStatus defines the observed state of PolicyLbService.
type PolicyLbServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyLbServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyLbService is the Schema for the PolicyLbServices API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyLbService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   PolicyLbServiceSpec   `json:"spec"`
	Status PolicyLbServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyLbServiceList contains a list of PolicyLbServices
type PolicyLbServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyLbService `json:"items"`
}

// Repository type metadata.
var (
	PolicyLbService_Kind             = "PolicyLbService"
	PolicyLbService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyLbService_Kind}.String()
	PolicyLbService_KindAPIVersion   = PolicyLbService_Kind + "." + CRDGroupVersion.String()
	PolicyLbService_GroupVersionKind = CRDGroupVersion.WithKind(PolicyLbService_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyLbService{}, &PolicyLbServiceList{})
}
