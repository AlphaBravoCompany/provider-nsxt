package fabric

import (
	"github.com/upbound/upjet/pkg/config"
)

const version = "v1alpha1"
const shortGroup = "fabric"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_cluster_virtual_ip", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ClusterVirtualIp"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_compute_manager", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ComputeManager"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_edge_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EdgeCluster"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_edge_high_availability_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EdgeHighAvailabilityProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_edge_transport_node", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EdgeTransportNode"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_failure_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FailureDomain"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_manager_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ManagerCluster"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_compute_sub_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyComputeSubCluster"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_host_transport_node", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyHostTransportNode"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_host_transport_node_collection", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyHostTransportNodeCollection"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_host_transport_node_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyHostTransportNodeProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_transport_zone", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTransportZone"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_uplink_host_switch_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyUplinkHostSwitchProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_vtep_ha_host_switch_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyVtepHaHostSwitchProfile"
		r.Version = version
	})
}
