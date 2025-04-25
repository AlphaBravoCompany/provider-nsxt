package firewall

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "firewall"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_context_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyContextProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_context_profile_custom_attribute", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyContextProfileCustomAttribute"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_firewall_exclude_list_member", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyFirewallExcludeListMember"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_intrusion_service_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIntrusionServicePolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_intrusion_service_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIntrusionServiceProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_parent_security_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyParentSecurityPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_predefined_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyPredefinedGatewayPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_predefined_security_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyPredefinedSecurityPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_security_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySecurityPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_security_policy_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySecurityPolicyRule"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyService"
		r.Version = version
	})
}
