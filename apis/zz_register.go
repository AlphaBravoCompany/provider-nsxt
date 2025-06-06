/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/AlphaBravoCompany/provider-nsxt/apis/dhcp/v1alpha1"
	v1alpha1dns "github.com/AlphaBravoCompany/provider-nsxt/apis/dns/v1alpha1"
	v1alpha1evpn "github.com/AlphaBravoCompany/provider-nsxt/apis/evpn/v1alpha1"
	v1alpha1firewall "github.com/AlphaBravoCompany/provider-nsxt/apis/firewall/v1alpha1"
	v1alpha1gatewaysandrouting "github.com/AlphaBravoCompany/provider-nsxt/apis/gatewaysandrouting/v1alpha1"
	v1alpha1groupingandtagging "github.com/AlphaBravoCompany/provider-nsxt/apis/groupingandtagging/v1alpha1"
	v1alpha1ipam "github.com/AlphaBravoCompany/provider-nsxt/apis/ipam/v1alpha1"
	v1alpha1loadbalancer "github.com/AlphaBravoCompany/provider-nsxt/apis/loadbalancer/v1alpha1"
	v1alpha1multitenancy "github.com/AlphaBravoCompany/provider-nsxt/apis/multitenancy/v1alpha1"
	v1alpha1nsxtalgorithmtypensservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtalgorithmtypensservice/v1alpha1"
	v1alpha1nsxtclustervirtualip "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtclustervirtualip/v1alpha1"
	v1alpha1nsxtcomputemanager "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtcomputemanager/v1alpha1"
	v1alpha1nsxtdhcprelayprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtdhcprelayprofile/v1alpha1"
	v1alpha1nsxtdhcprelayservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtdhcprelayservice/v1alpha1"
	v1alpha1nsxtdhcpserverippool "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtdhcpserverippool/v1alpha1"
	v1alpha1nsxtdhcpserverprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtdhcpserverprofile/v1alpha1"
	v1alpha1nsxtedgecluster "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtedgecluster/v1alpha1"
	v1alpha1nsxtethertypensservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtethertypensservice/v1alpha1"
	v1alpha1nsxtfailuredomain "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtfailuredomain/v1alpha1"
	v1alpha1nsxtfirewallsection "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtfirewallsection/v1alpha1"
	v1alpha1nsxticmptypensservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxticmptypensservice/v1alpha1"
	v1alpha1nsxtigmptypensservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtigmptypensservice/v1alpha1"
	v1alpha1nsxtipblock "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtipblock/v1alpha1"
	v1alpha1nsxtipblocksubnet "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtipblocksubnet/v1alpha1"
	v1alpha1nsxtipdiscoveryswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtipdiscoveryswitchingprofile/v1alpha1"
	v1alpha1nsxtippool "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtippool/v1alpha1"
	v1alpha1nsxtippoolallocationipaddress "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtippoolallocationipaddress/v1alpha1"
	v1alpha1nsxtipprotocolnsservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtipprotocolnsservice/v1alpha1"
	v1alpha1nsxtipset "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtipset/v1alpha1"
	v1alpha1nsxtl4portsetnsservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtl4portsetnsservice/v1alpha1"
	v1alpha1nsxtlbclientsslprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbclientsslprofile/v1alpha1"
	v1alpha1nsxtlbcookiepersistenceprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbcookiepersistenceprofile/v1alpha1"
	v1alpha1nsxtlbfasttcpapplicationprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbfasttcpapplicationprofile/v1alpha1"
	v1alpha1nsxtlbfastudpapplicationprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbfastudpapplicationprofile/v1alpha1"
	v1alpha1nsxtlbhttpapplicationprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttpapplicationprofile/v1alpha1"
	v1alpha1nsxtlbhttpforwardingrule "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttpforwardingrule/v1alpha1"
	v1alpha1nsxtlbhttpmonitor "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttpmonitor/v1alpha1"
	v1alpha1nsxtlbhttprequestrewriterule "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttprequestrewriterule/v1alpha1"
	v1alpha1nsxtlbhttpresponserewriterule "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttpresponserewriterule/v1alpha1"
	v1alpha1nsxtlbhttpsmonitor "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttpsmonitor/v1alpha1"
	v1alpha1nsxtlbhttpvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbhttpvirtualserver/v1alpha1"
	v1alpha1nsxtlbicmpmonitor "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbicmpmonitor/v1alpha1"
	v1alpha1nsxtlbpassivemonitor "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbpassivemonitor/v1alpha1"
	v1alpha1nsxtlbpool "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbpool/v1alpha1"
	v1alpha1nsxtlbserversslprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbserversslprofile/v1alpha1"
	v1alpha1nsxtlbservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbservice/v1alpha1"
	v1alpha1nsxtlbsourceippersistenceprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbsourceippersistenceprofile/v1alpha1"
	v1alpha1nsxtlbtcpmonitor "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbtcpmonitor/v1alpha1"
	v1alpha1nsxtlbtcpvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbtcpvirtualserver/v1alpha1"
	v1alpha1nsxtlbudpmonitor "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbudpmonitor/v1alpha1"
	v1alpha1nsxtlbudpvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlbudpvirtualserver/v1alpha1"
	v1alpha1nsxtlogicaldhcpport "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicaldhcpport/v1alpha1"
	v1alpha1nsxtlogicaldhcpserver "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicaldhcpserver/v1alpha1"
	v1alpha1nsxtlogicalport "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicalport/v1alpha1"
	v1alpha1nsxtlogicalroutercentralizedserviceport "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicalroutercentralizedserviceport/v1alpha1"
	v1alpha1nsxtlogicalrouterdownlinkport "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicalrouterdownlinkport/v1alpha1"
	v1alpha1nsxtlogicalrouterlinkportontier0 "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicalrouterlinkportontier0/v1alpha1"
	v1alpha1nsxtlogicalrouterlinkportontier1 "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicalrouterlinkportontier1/v1alpha1"
	v1alpha1nsxtlogicalswitch "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicalswitch/v1alpha1"
	v1alpha1nsxtlogicaltier0router "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicaltier0router/v1alpha1"
	v1alpha1nsxtlogicaltier1router "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtlogicaltier1router/v1alpha1"
	v1alpha1nsxtmacmanagementswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtmacmanagementswitchingprofile/v1alpha1"
	v1alpha1nsxtmanagercluster "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtmanagercluster/v1alpha1"
	v1alpha1nsxtnatrule "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtnatrule/v1alpha1"
	v1alpha1nsxtnsgroup "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtnsgroup/v1alpha1"
	v1alpha1nsxtnsservicegroup "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtnsservicegroup/v1alpha1"
	v1alpha1nsxtpolicyhosttransportnodeprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyhosttransportnodeprofile/v1alpha1"
	v1alpha1nsxtpolicyipsecvpndpdprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyipsecvpndpdprofile/v1alpha1"
	v1alpha1nsxtpolicyipsecvpnikeprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyipsecvpnikeprofile/v1alpha1"
	v1alpha1nsxtpolicyipsecvpnlocalendpoint "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyipsecvpnlocalendpoint/v1alpha1"
	v1alpha1nsxtpolicyipsecvpnservice "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyipsecvpnservice/v1alpha1"
	v1alpha1nsxtpolicyipsecvpnsession "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyipsecvpnsession/v1alpha1"
	v1alpha1nsxtpolicyl2vpnsession "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicyl2vpnsession/v1alpha1"
	v1alpha1nsxtpolicytransportzone "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtpolicytransportzone/v1alpha1"
	v1alpha1nsxtqosswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtqosswitchingprofile/v1alpha1"
	v1alpha1nsxtspoofguardswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtspoofguardswitchingprofile/v1alpha1"
	v1alpha1nsxtstaticroute "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtstaticroute/v1alpha1"
	v1alpha1nsxtswitchsecurityswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtswitchsecurityswitchingprofile/v1alpha1"
	v1alpha1nsxttransportnode "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxttransportnode/v1alpha1"
	v1alpha1nsxtuplinkhostswitchprofile "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtuplinkhostswitchprofile/v1alpha1"
	v1alpha1nsxtvlanlogicalswitch "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtvlanlogicalswitch/v1alpha1"
	v1alpha1nsxtvmtags "github.com/AlphaBravoCompany/provider-nsxt/apis/nsxtvmtags/v1alpha1"
	v1alpha1segments "github.com/AlphaBravoCompany/provider-nsxt/apis/segments/v1alpha1"
	v1alpha1apis "github.com/AlphaBravoCompany/provider-nsxt/apis/v1alpha1"
	v1beta1 "github.com/AlphaBravoCompany/provider-nsxt/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1evpn.SchemeBuilder.AddToScheme,
		v1alpha1firewall.SchemeBuilder.AddToScheme,
		v1alpha1gatewaysandrouting.SchemeBuilder.AddToScheme,
		v1alpha1groupingandtagging.SchemeBuilder.AddToScheme,
		v1alpha1ipam.SchemeBuilder.AddToScheme,
		v1alpha1loadbalancer.SchemeBuilder.AddToScheme,
		v1alpha1multitenancy.SchemeBuilder.AddToScheme,
		v1alpha1nsxtalgorithmtypensservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtclustervirtualip.SchemeBuilder.AddToScheme,
		v1alpha1nsxtcomputemanager.SchemeBuilder.AddToScheme,
		v1alpha1nsxtdhcprelayprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtdhcprelayservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtdhcpserverippool.SchemeBuilder.AddToScheme,
		v1alpha1nsxtdhcpserverprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtedgecluster.SchemeBuilder.AddToScheme,
		v1alpha1nsxtethertypensservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtfailuredomain.SchemeBuilder.AddToScheme,
		v1alpha1nsxtfirewallsection.SchemeBuilder.AddToScheme,
		v1alpha1nsxticmptypensservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtigmptypensservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtipblock.SchemeBuilder.AddToScheme,
		v1alpha1nsxtipblocksubnet.SchemeBuilder.AddToScheme,
		v1alpha1nsxtipdiscoveryswitchingprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtippool.SchemeBuilder.AddToScheme,
		v1alpha1nsxtippoolallocationipaddress.SchemeBuilder.AddToScheme,
		v1alpha1nsxtipprotocolnsservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtipset.SchemeBuilder.AddToScheme,
		v1alpha1nsxtl4portsetnsservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbclientsslprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbcookiepersistenceprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbfasttcpapplicationprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbfastudpapplicationprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttpapplicationprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttpforwardingrule.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttpmonitor.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttprequestrewriterule.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttpresponserewriterule.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttpsmonitor.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbhttpvirtualserver.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbicmpmonitor.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbpassivemonitor.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbpool.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbserversslprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbsourceippersistenceprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbtcpmonitor.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbtcpvirtualserver.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbudpmonitor.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlbudpvirtualserver.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicaldhcpport.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicaldhcpserver.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicalport.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicalroutercentralizedserviceport.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicalrouterdownlinkport.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicalrouterlinkportontier0.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicalrouterlinkportontier1.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicalswitch.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicaltier0router.SchemeBuilder.AddToScheme,
		v1alpha1nsxtlogicaltier1router.SchemeBuilder.AddToScheme,
		v1alpha1nsxtmacmanagementswitchingprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtmanagercluster.SchemeBuilder.AddToScheme,
		v1alpha1nsxtnatrule.SchemeBuilder.AddToScheme,
		v1alpha1nsxtnsgroup.SchemeBuilder.AddToScheme,
		v1alpha1nsxtnsservicegroup.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyhosttransportnodeprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyipsecvpndpdprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyipsecvpnikeprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyipsecvpnlocalendpoint.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyipsecvpnservice.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyipsecvpnsession.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicyl2vpnsession.SchemeBuilder.AddToScheme,
		v1alpha1nsxtpolicytransportzone.SchemeBuilder.AddToScheme,
		v1alpha1nsxtqosswitchingprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtspoofguardswitchingprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtstaticroute.SchemeBuilder.AddToScheme,
		v1alpha1nsxtswitchsecurityswitchingprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxttransportnode.SchemeBuilder.AddToScheme,
		v1alpha1nsxtuplinkhostswitchprofile.SchemeBuilder.AddToScheme,
		v1alpha1nsxtvlanlogicalswitch.SchemeBuilder.AddToScheme,
		v1alpha1nsxtvmtags.SchemeBuilder.AddToScheme,
		v1alpha1segments.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
