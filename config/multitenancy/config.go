package multitenancy

import "github.com/upbound/upjet/pkg/config"

const (
	shortGroup = "multitenancy"
	version    = "v1alpha1"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyProject"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_share", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyShare"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_shared_resource", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySharedResource"
		r.Version = version
	})
}
