package beta

import (
	"github.com/upbound/upjet/pkg/config"
)

const version string = "v1alpha1"
const shortGroup string = "beta"

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

	p.AddResourceConfigurator("nsxt_transport_node", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TransportNode"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_uplink_host_switch_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UplinkHostSwitchProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_edge_transport_node_rtep", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EdgeTransportNodeRtep"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_distributed_flood_protection_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyDistributedFloodProtectionProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_distributed_flood_protection_profile_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyDistributedFloodProtectionProfileBinding"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_flood_protection_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayFloodProtectionProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_flood_protection_profile_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayFloodProtectionProfileBinding"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_global_manager", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGlobalManager"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_client_ssl_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbClientSslProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_http_application_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbHttpApplicationProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_http_monitor_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbHttpMonitorProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_https_monitor_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbHttpsMonitorProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_icmp_monitor_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbIcmpMonitorProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_passive_monitor_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbPassiveMonitorProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_tcp_monitor_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbTcpMonitorProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_lb_udp_monitor_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyLbUdpMonitorProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_metadata_proxy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyMetadataProxy"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_site", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySite"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_tier0_gateway_gre_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier0GatewayGreTunnel"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_tier0_inter_vrf_routing", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier0InterVrfRouting"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_upgrade_precheck_acknowledge", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UpgradePrecheckAcknowledge"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_upgrade_prepare", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UpgradePrepare"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_upgrade_run", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UpgradeRun"
		r.Version = version
	})
}
