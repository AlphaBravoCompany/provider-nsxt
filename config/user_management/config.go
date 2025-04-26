package usermanagement

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroup = "usermanagement"
	version    = "v1alpha1"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_node_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NodeUser"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ldap_identity_source", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLdapIdentitySource"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_user_management_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyUserManagementRole"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_user_management_role_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyUserManagementRoleBinding"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_principal_identity", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PrincipalIdentity"
		r.Version = version
	})
}
