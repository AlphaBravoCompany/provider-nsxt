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

type AlgorithmTypeNsServiceObservation struct {

	// Algorithm
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A boolean flag which reflects whether this is a default NSServices which can't be modified/deleted
	DefaultService *bool `json:"defaultService,omitempty" tf:"default_service,omitempty"`

	// Description of this resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A single destination port
	DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// The display name of this resource. Defaults to ID if not set
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Set of source ports or ranges
	SourcePorts []*string `json:"sourcePorts,omitempty" tf:"source_ports,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type AlgorithmTypeNsServiceParameters struct {

	// Algorithm
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Description of this resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A single destination port
	// +kubebuilder:validation:Optional
	DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// The display name of this resource. Defaults to ID if not set
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Set of source ports or ranges
	// +kubebuilder:validation:Optional
	SourcePorts []*string `json:"sourcePorts,omitempty" tf:"source_ports,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type TagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type TagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// AlgorithmTypeNsServiceSpec defines the desired state of AlgorithmTypeNsService
type AlgorithmTypeNsServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlgorithmTypeNsServiceParameters `json:"forProvider"`
}

// AlgorithmTypeNsServiceStatus defines the observed state of AlgorithmTypeNsService.
type AlgorithmTypeNsServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlgorithmTypeNsServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlgorithmTypeNsService is the Schema for the AlgorithmTypeNsServices API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type AlgorithmTypeNsService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.algorithm)",message="algorithm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.destinationPort)",message="destinationPort is a required parameter"
	Spec   AlgorithmTypeNsServiceSpec   `json:"spec"`
	Status AlgorithmTypeNsServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlgorithmTypeNsServiceList contains a list of AlgorithmTypeNsServices
type AlgorithmTypeNsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlgorithmTypeNsService `json:"items"`
}

// Repository type metadata.
var (
	AlgorithmTypeNsService_Kind             = "AlgorithmTypeNsService"
	AlgorithmTypeNsService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlgorithmTypeNsService_Kind}.String()
	AlgorithmTypeNsService_KindAPIVersion   = AlgorithmTypeNsService_Kind + "." + CRDGroupVersion.String()
	AlgorithmTypeNsService_GroupVersionKind = CRDGroupVersion.WithKind(AlgorithmTypeNsService_Kind)
)

func init() {
	SchemeBuilder.Register(&AlgorithmTypeNsService{}, &AlgorithmTypeNsServiceList{})
}
