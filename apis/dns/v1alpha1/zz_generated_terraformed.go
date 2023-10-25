/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this PolicyDnsForwarderZone
func (mg *PolicyDnsForwarderZone) GetTerraformResourceType() string {
	return "nsxt_policy_dns_forwarder_zone"
}

// GetConnectionDetailsMapping for this PolicyDnsForwarderZone
func (tr *PolicyDnsForwarderZone) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this PolicyDnsForwarderZone
func (tr *PolicyDnsForwarderZone) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PolicyDnsForwarderZone
func (tr *PolicyDnsForwarderZone) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this PolicyDnsForwarderZone
func (tr *PolicyDnsForwarderZone) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this PolicyDnsForwarderZone
func (tr *PolicyDnsForwarderZone) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PolicyDnsForwarderZone
func (tr *PolicyDnsForwarderZone) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this PolicyDnsForwarderZone using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PolicyDnsForwarderZone) LateInitialize(attrs []byte) (bool, error) {
	params := &PolicyDnsForwarderZoneParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *PolicyDnsForwarderZone) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this PolicyGatewayDnsForwarder
func (mg *PolicyGatewayDnsForwarder) GetTerraformResourceType() string {
	return "nsxt_policy_gateway_dns_forwarder"
}

// GetConnectionDetailsMapping for this PolicyGatewayDnsForwarder
func (tr *PolicyGatewayDnsForwarder) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this PolicyGatewayDnsForwarder
func (tr *PolicyGatewayDnsForwarder) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PolicyGatewayDnsForwarder
func (tr *PolicyGatewayDnsForwarder) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this PolicyGatewayDnsForwarder
func (tr *PolicyGatewayDnsForwarder) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this PolicyGatewayDnsForwarder
func (tr *PolicyGatewayDnsForwarder) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PolicyGatewayDnsForwarder
func (tr *PolicyGatewayDnsForwarder) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this PolicyGatewayDnsForwarder using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PolicyGatewayDnsForwarder) LateInitialize(attrs []byte) (bool, error) {
	params := &PolicyGatewayDnsForwarderParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *PolicyGatewayDnsForwarder) GetTerraformSchemaVersion() int {
	return 0
}
