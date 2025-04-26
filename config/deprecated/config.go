package deprecated

import (
	"github.com/upbound/upjet/pkg/config"
)

const version = "v1alpha1"
const shortGroup = "deprecated"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("nsxt_algorithm_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AlgorithmTypeNsService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_dhcp_relay_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DhcpRelayProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_dhcp_relay_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DhcpRelayService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_dhcp_server_ip_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DhcpServerIpPool"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_dhcp_server_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DhcpServerProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ether_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EtherTypeNsService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_firewall_section", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FirewallSection"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_icmp_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IcmpTypeNsService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_igmp_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IgmpTypeNsService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_block", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpBlock"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_block_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpBlockSubnet"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_discovery_switching_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpDiscoverySwitchingProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpPool"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_pool_allocation_ip_address", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpPoolAllocationIpAddress"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_protocol_ns_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpProtocolNsService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ip_set", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IpSet"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_l4_port_set_ns_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "L4PortSetNsService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_client_ssl_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbClientSslProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_cookie_persistence_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbCookiePersistenceProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_fast_tcp_application_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbFastTcpApplicationProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_fast_udp_application_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbFastUdpApplicationProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_http_application_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpApplicationProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_http_forwarding_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpForwardingRule"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_http_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpMonitor"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_http_request_rewrite_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpRequestRewriteRule"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_http_response_rewrite_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpResponseRewriteRule"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_http_virtual_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpVirtualServer"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_https_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbHttpsMonitor"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_icmp_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbIcmpMonitor"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_passive_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbPassiveMonitor"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbPool"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_server_ssl_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbServerSslProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbService"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_source_ip_persistence_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbSourceIpPersistenceProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_tcp_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbTcpMonitor"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_tcp_virtual_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbTcpVirtualServer"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_udp_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbUdpMonitor"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_lb_udp_virtual_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbUdpVirtualServer"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_dhcp_port", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalDhcpPort"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_dhcp_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalDhcpServer"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_port", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalPort"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_router_centralized_service_port", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalRouterCentralizedServicePort"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_router_downlink_port", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalRouterDownlinkPort"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_router_link_port_on_tier0", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalRouterLinkPortOnTier0"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_router_link_port_on_tier1", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalRouterLinkPortOnTier1"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_switch", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalSwitch"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_tier0_router", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalTier0Router"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_logical_tier1_router", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogicalTier1Router"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_mac_management_switching_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MacManagementSwitchingProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_nat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NatRule"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ns_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsGroup"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_ns_service_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsServiceGroup"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_qos_switching_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "QosSwitchingProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_spoofguard_switching_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SpoofguardSwitchingProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StaticRoute"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_switch_security_switching_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SwitchSecuritySwitchingProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_vlan_logical_switch", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VlanLogicalSwitch"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_vm_tags", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VmTags"
		r.Version = version
	})
}
