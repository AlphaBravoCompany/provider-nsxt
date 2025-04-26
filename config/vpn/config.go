package vpn

import "github.com/upbound/upjet/pkg/config"

const version = "v1alpha1"
const shortGroup = "vpn"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_dpd_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpsecVpnDpdProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_ike_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpsecVpnIkeProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_local_endpoint", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpsecVpnLocalEndpoint"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpsecVpnService"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_session", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpsecVpnSession"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_tunnel_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpsecVpnTunnelProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_l2_vpn_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyL2VpnService"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_l2_vpn_session", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyL2VpnSession"
		r.Version = version
	})
}
