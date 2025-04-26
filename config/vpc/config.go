package vpc

import "github.com/upbound/upjet/pkg/config"

const version = "v1alpha1"
const shortGroup = "vpc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_vpc_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VpcGatewayPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_vpc_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VpcGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_vpc_security_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VpcSecurityPolicy"
		r.Version = version
	})
}
