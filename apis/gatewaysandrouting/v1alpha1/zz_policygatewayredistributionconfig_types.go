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

type PolicyGatewayRedistributionConfigObservation struct {

	// Flag to enable route redistribution for BGP
	BGPEnabled *bool `json:"bgpEnabled,omitempty" tf:"bgp_enabled,omitempty"`

	// Id of associated Tier0 Gateway on NSX
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// Policy path for Tier0 gateway
	GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Id of associated Gateway Locale Service on NSX
	LocaleServiceID *string `json:"localeServiceId,omitempty" tf:"locale_service_id,omitempty"`

	// Flag to enable route redistribution for OSPF
	OspfEnabled *bool `json:"ospfEnabled,omitempty" tf:"ospf_enabled,omitempty"`

	// List of routes to be aggregated
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// Path of the site the Tier0 redistribution
	SitePath *string `json:"sitePath,omitempty" tf:"site_path,omitempty"`
}

type PolicyGatewayRedistributionConfigParameters struct {

	// Flag to enable route redistribution for BGP
	// +kubebuilder:validation:Optional
	BGPEnabled *bool `json:"bgpEnabled,omitempty" tf:"bgp_enabled,omitempty"`

	// Policy path for Tier0 gateway
	// +kubebuilder:validation:Optional
	GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

	// Flag to enable route redistribution for OSPF
	// +kubebuilder:validation:Optional
	OspfEnabled *bool `json:"ospfEnabled,omitempty" tf:"ospf_enabled,omitempty"`

	// List of routes to be aggregated
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Path of the site the Tier0 redistribution
	// +kubebuilder:validation:Optional
	SitePath *string `json:"sitePath,omitempty" tf:"site_path,omitempty"`
}

type RuleObservation struct {

	// BGP destination for this rule
	BGP *bool `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// Rule name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OSPF destination for this rule
	Ospf *bool `json:"ospf,omitempty" tf:"ospf,omitempty"`

	// Route map to be associated with the redistribution rule
	RouteMapPath *string `json:"routeMapPath,omitempty" tf:"route_map_path,omitempty"`

	// List of redistribution types
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type RuleParameters struct {

	// BGP destination for this rule
	// +kubebuilder:validation:Optional
	BGP *bool `json:"bgp,omitempty" tf:"bgp,omitempty"`

	// Rule name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OSPF destination for this rule
	// +kubebuilder:validation:Optional
	Ospf *bool `json:"ospf,omitempty" tf:"ospf,omitempty"`

	// Route map to be associated with the redistribution rule
	// +kubebuilder:validation:Optional
	RouteMapPath *string `json:"routeMapPath,omitempty" tf:"route_map_path,omitempty"`

	// List of redistribution types
	// +kubebuilder:validation:Optional
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

// PolicyGatewayRedistributionConfigSpec defines the desired state of PolicyGatewayRedistributionConfig
type PolicyGatewayRedistributionConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyGatewayRedistributionConfigParameters `json:"forProvider"`
}

// PolicyGatewayRedistributionConfigStatus defines the observed state of PolicyGatewayRedistributionConfig.
type PolicyGatewayRedistributionConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyGatewayRedistributionConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGatewayRedistributionConfig is the Schema for the PolicyGatewayRedistributionConfigs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyGatewayRedistributionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.gatewayPath)",message="gatewayPath is a required parameter"
	Spec   PolicyGatewayRedistributionConfigSpec   `json:"spec"`
	Status PolicyGatewayRedistributionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGatewayRedistributionConfigList contains a list of PolicyGatewayRedistributionConfigs
type PolicyGatewayRedistributionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyGatewayRedistributionConfig `json:"items"`
}

// Repository type metadata.
var (
	PolicyGatewayRedistributionConfig_Kind             = "PolicyGatewayRedistributionConfig"
	PolicyGatewayRedistributionConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyGatewayRedistributionConfig_Kind}.String()
	PolicyGatewayRedistributionConfig_KindAPIVersion   = PolicyGatewayRedistributionConfig_Kind + "." + CRDGroupVersion.String()
	PolicyGatewayRedistributionConfig_GroupVersionKind = CRDGroupVersion.WithKind(PolicyGatewayRedistributionConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyGatewayRedistributionConfig{}, &PolicyGatewayRedistributionConfigList{})
}
