/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policydhcprelay "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/dhcp/policydhcprelay"
	policydhcpserver "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/dhcp/policydhcpserver"
	policydhcpv4staticbinding "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/dhcp/policydhcpv4staticbinding"
	policydhcpv6staticbinding "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/dhcp/policydhcpv6staticbinding"
	policydnsforwarderzone "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/dns/policydnsforwarderzone"
	policygatewaydnsforwarder "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/dns/policygatewaydnsforwarder"
	policyevpnconfig "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/evpn/policyevpnconfig"
	policyevpntenant "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/evpn/policyevpntenant"
	policyevpntunnelendpoint "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/evpn/policyevpntunnelendpoint"
	policyvnipool "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/evpn/policyvnipool"
	policycontextprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policycontextprofile"
	policycontextprofilecustomattribute "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policycontextprofilecustomattribute"
	policyfirewallexcludelistmember "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policyfirewallexcludelistmember"
	policygatewaypolicy "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policygatewaypolicy"
	policyintrusionservicepolicy "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policyintrusionservicepolicy"
	policyintrusionserviceprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policyintrusionserviceprofile"
	policyparentsecuritypolicy "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policyparentsecuritypolicy"
	policypredefinedgatewaypolicy "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policypredefinedgatewaypolicy"
	policypredefinedsecuritypolicy "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policypredefinedsecuritypolicy"
	policysecuritypolicy "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policysecuritypolicy"
	policysecuritypolicyrule "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policysecuritypolicyrule"
	policyservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/firewall/policyservice"
	policybgpconfig "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policybgpconfig"
	policybgpneighbor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policybgpneighbor"
	policygatewaycommunitylist "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policygatewaycommunitylist"
	policygatewayprefixlist "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayprefixlist"
	policygatewayqosprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayqosprofile"
	policygatewayredistributionconfig "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayredistributionconfig"
	policygatewayroutemap "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayroutemap"
	policynatrule "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policynatrule"
	policyospfconfig "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policyospfconfig"
	policystaticroute "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroute"
	policytier0gateway "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gateway"
	policytier0gatewayhavipconfig "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayhavipconfig"
	policytier0gatewayinterface "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayinterface"
	policytier1gateway "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gateway"
	policytier1gatewayinterface "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gatewayinterface"
	policydomain "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/groupingandtagging/policydomain"
	policygroup "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/groupingandtagging/policygroup"
	policyvmtags "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/groupingandtagging/policyvmtags"
	policyipaddressallocation "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/ipam/policyipaddressallocation"
	policyipblock "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/ipam/policyipblock"
	policyippool "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/ipam/policyippool"
	policyippoolblocksubnet "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/ipam/policyippoolblocksubnet"
	policyippoolstaticsubnet "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/ipam/policyippoolstaticsubnet"
	policylbpool "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/loadbalancer/policylbpool"
	policylbservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/loadbalancer/policylbservice"
	policylbvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/loadbalancer/policylbvirtualserver"
	policyproject "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/multitenancy/policyproject"
	policyshare "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/multitenancy/policyshare"
	policysharedresource "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/multitenancy/policysharedresource"
	algorithmtypensservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtalgorithmtypensservice/algorithmtypensservice"
	clustervirtualip "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtclustervirtualip/clustervirtualip"
	computemanager "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtcomputemanager/computemanager"
	dhcprelayprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtdhcprelayprofile/dhcprelayprofile"
	dhcprelayservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtdhcprelayservice/dhcprelayservice"
	dhcpserverippool "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtdhcpserverippool/dhcpserverippool"
	dhcpserverprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtdhcpserverprofile/dhcpserverprofile"
	edgecluster "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtedgecluster/edgecluster"
	ethertypensservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtethertypensservice/ethertypensservice"
	failuredomain "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtfailuredomain/failuredomain"
	firewallsection "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtfirewallsection/firewallsection"
	icmptypensservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxticmptypensservice/icmptypensservice"
	igmptypensservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtigmptypensservice/igmptypensservice"
	ipblock "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtipblock/ipblock"
	ipblocksubnet "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtipblocksubnet/ipblocksubnet"
	ipdiscoveryswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtipdiscoveryswitchingprofile/ipdiscoveryswitchingprofile"
	ippool "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtippool/ippool"
	ippoolallocationipaddress "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtippoolallocationipaddress/ippoolallocationipaddress"
	ipprotocolnsservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtipprotocolnsservice/ipprotocolnsservice"
	ipset "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtipset/ipset"
	l4portsetnsservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtl4portsetnsservice/l4portsetnsservice"
	lbclientsslprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbclientsslprofile/lbclientsslprofile"
	lbcookiepersistenceprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbcookiepersistenceprofile/lbcookiepersistenceprofile"
	lbfasttcpapplicationprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbfasttcpapplicationprofile/lbfasttcpapplicationprofile"
	lbfastudpapplicationprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbfastudpapplicationprofile/lbfastudpapplicationprofile"
	lbhttpapplicationprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttpapplicationprofile/lbhttpapplicationprofile"
	lbhttpforwardingrule "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttpforwardingrule/lbhttpforwardingrule"
	lbhttpmonitor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttpmonitor/lbhttpmonitor"
	lbhttprequestrewriterule "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttprequestrewriterule/lbhttprequestrewriterule"
	lbhttpresponserewriterule "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttpresponserewriterule/lbhttpresponserewriterule"
	lbhttpsmonitor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttpsmonitor/lbhttpsmonitor"
	lbhttpvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbhttpvirtualserver/lbhttpvirtualserver"
	lbicmpmonitor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbicmpmonitor/lbicmpmonitor"
	lbpassivemonitor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbpassivemonitor/lbpassivemonitor"
	lbpool "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbpool/lbpool"
	lbserversslprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbserversslprofile/lbserversslprofile"
	lbservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbservice/lbservice"
	lbsourceippersistenceprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbsourceippersistenceprofile/lbsourceippersistenceprofile"
	lbtcpmonitor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbtcpmonitor/lbtcpmonitor"
	lbtcpvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbtcpvirtualserver/lbtcpvirtualserver"
	lbudpmonitor "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbudpmonitor/lbudpmonitor"
	lbudpvirtualserver "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlbudpvirtualserver/lbudpvirtualserver"
	logicaldhcpport "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicaldhcpport/logicaldhcpport"
	logicaldhcpserver "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicaldhcpserver/logicaldhcpserver"
	logicalport "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicalport/logicalport"
	logicalroutercentralizedserviceport "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicalroutercentralizedserviceport/logicalroutercentralizedserviceport"
	logicalrouterdownlinkport "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicalrouterdownlinkport/logicalrouterdownlinkport"
	logicalrouterlinkportontier0 "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicalrouterlinkportontier0/logicalrouterlinkportontier0"
	logicalrouterlinkportontier1 "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicalrouterlinkportontier1/logicalrouterlinkportontier1"
	logicalswitch "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicalswitch/logicalswitch"
	logicaltier0router "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicaltier0router/logicaltier0router"
	logicaltier1router "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtlogicaltier1router/logicaltier1router"
	macmanagementswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtmacmanagementswitchingprofile/macmanagementswitchingprofile"
	managercluster "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtmanagercluster/managercluster"
	natrule "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtnatrule/natrule"
	nsgroup "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtnsgroup/nsgroup"
	nsservicegroup "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtnsservicegroup/nsservicegroup"
	policyhosttransportnodeprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyhosttransportnodeprofile/policyhosttransportnodeprofile"
	policyipsecvpndpdprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyipsecvpndpdprofile/policyipsecvpndpdprofile"
	policyipsecvpnikeprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyipsecvpnikeprofile/policyipsecvpnikeprofile"
	policyipsecvpnlocalendpoint "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyipsecvpnlocalendpoint/policyipsecvpnlocalendpoint"
	policyipsecvpnservice "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyipsecvpnservice/policyipsecvpnservice"
	policyipsecvpnsession "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyipsecvpnsession/policyipsecvpnsession"
	policyipsecvpntunnelprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyipsecvpntunnelprofile/policyipsecvpntunnelprofile"
	policyl2vpnsession "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicyl2vpnsession/policyl2vpnsession"
	policytransportzone "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtpolicytransportzone/policytransportzone"
	qosswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtqosswitchingprofile/qosswitchingprofile"
	spoofguardswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtspoofguardswitchingprofile/spoofguardswitchingprofile"
	staticroute "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtstaticroute/staticroute"
	switchsecurityswitchingprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtswitchsecurityswitchingprofile/switchsecurityswitchingprofile"
	vlanlogicalswitch "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtvlanlogicalswitch/vlanlogicalswitch"
	vmtags "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/nsxtvmtags/vmtags"
	providerconfig "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/providerconfig"
	policyfixedsegment "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policyfixedsegment"
	policyipdiscoveryprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policyipdiscoveryprofile"
	policymacdiscoveryprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policymacdiscoveryprofile"
	policyqosprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policyqosprofile"
	policysegment "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policysegment"
	policysegmentsecurityprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policysegmentsecurityprofile"
	policyspoofguardprofile "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policyspoofguardprofile"
	policyvlansegment "github.com/AlphaBravoCompany/provider-nsxt/internal/controller/segments/policyvlansegment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policydhcprelay.Setup,
		policydhcpserver.Setup,
		policydhcpv4staticbinding.Setup,
		policydhcpv6staticbinding.Setup,
		policydnsforwarderzone.Setup,
		policygatewaydnsforwarder.Setup,
		policyevpnconfig.Setup,
		policyevpntenant.Setup,
		policyevpntunnelendpoint.Setup,
		policyvnipool.Setup,
		policycontextprofile.Setup,
		policycontextprofilecustomattribute.Setup,
		policyfirewallexcludelistmember.Setup,
		policygatewaypolicy.Setup,
		policyintrusionservicepolicy.Setup,
		policyintrusionserviceprofile.Setup,
		policyparentsecuritypolicy.Setup,
		policypredefinedgatewaypolicy.Setup,
		policypredefinedsecuritypolicy.Setup,
		policysecuritypolicy.Setup,
		policysecuritypolicyrule.Setup,
		policyservice.Setup,
		policybgpconfig.Setup,
		policybgpneighbor.Setup,
		policygatewaycommunitylist.Setup,
		policygatewayprefixlist.Setup,
		policygatewayqosprofile.Setup,
		policygatewayredistributionconfig.Setup,
		policygatewayroutemap.Setup,
		policynatrule.Setup,
		policyospfconfig.Setup,
		policystaticroute.Setup,
		policytier0gateway.Setup,
		policytier0gatewayhavipconfig.Setup,
		policytier0gatewayinterface.Setup,
		policytier1gateway.Setup,
		policytier1gatewayinterface.Setup,
		policydomain.Setup,
		policygroup.Setup,
		policyvmtags.Setup,
		policyipaddressallocation.Setup,
		policyipblock.Setup,
		policyippool.Setup,
		policyippoolblocksubnet.Setup,
		policyippoolstaticsubnet.Setup,
		policylbpool.Setup,
		policylbservice.Setup,
		policylbvirtualserver.Setup,
		policyproject.Setup,
		policyshare.Setup,
		policysharedresource.Setup,
		algorithmtypensservice.Setup,
		clustervirtualip.Setup,
		computemanager.Setup,
		dhcprelayprofile.Setup,
		dhcprelayservice.Setup,
		dhcpserverippool.Setup,
		dhcpserverprofile.Setup,
		edgecluster.Setup,
		ethertypensservice.Setup,
		failuredomain.Setup,
		firewallsection.Setup,
		icmptypensservice.Setup,
		igmptypensservice.Setup,
		ipblock.Setup,
		ipblocksubnet.Setup,
		ipdiscoveryswitchingprofile.Setup,
		ippool.Setup,
		ippoolallocationipaddress.Setup,
		ipprotocolnsservice.Setup,
		ipset.Setup,
		l4portsetnsservice.Setup,
		lbclientsslprofile.Setup,
		lbcookiepersistenceprofile.Setup,
		lbfasttcpapplicationprofile.Setup,
		lbfastudpapplicationprofile.Setup,
		lbhttpapplicationprofile.Setup,
		lbhttpforwardingrule.Setup,
		lbhttpmonitor.Setup,
		lbhttprequestrewriterule.Setup,
		lbhttpresponserewriterule.Setup,
		lbhttpsmonitor.Setup,
		lbhttpvirtualserver.Setup,
		lbicmpmonitor.Setup,
		lbpassivemonitor.Setup,
		lbpool.Setup,
		lbserversslprofile.Setup,
		lbservice.Setup,
		lbsourceippersistenceprofile.Setup,
		lbtcpmonitor.Setup,
		lbtcpvirtualserver.Setup,
		lbudpmonitor.Setup,
		lbudpvirtualserver.Setup,
		logicaldhcpport.Setup,
		logicaldhcpserver.Setup,
		logicalport.Setup,
		logicalroutercentralizedserviceport.Setup,
		logicalrouterdownlinkport.Setup,
		logicalrouterlinkportontier0.Setup,
		logicalrouterlinkportontier1.Setup,
		logicalswitch.Setup,
		logicaltier0router.Setup,
		logicaltier1router.Setup,
		macmanagementswitchingprofile.Setup,
		managercluster.Setup,
		natrule.Setup,
		nsgroup.Setup,
		nsservicegroup.Setup,
		policyhosttransportnodeprofile.Setup,
		policyipsecvpndpdprofile.Setup,
		policyipsecvpnikeprofile.Setup,
		policyipsecvpnlocalendpoint.Setup,
		policyipsecvpnservice.Setup,
		policyipsecvpnsession.Setup,
		policyipsecvpntunnelprofile.Setup,
		policyl2vpnsession.Setup,
		policytransportzone.Setup,
		qosswitchingprofile.Setup,
		spoofguardswitchingprofile.Setup,
		staticroute.Setup,
		switchsecurityswitchingprofile.Setup,
		vlanlogicalswitch.Setup,
		vmtags.Setup,
		providerconfig.Setup,
		policyfixedsegment.Setup,
		policyipdiscoveryprofile.Setup,
		policymacdiscoveryprofile.Setup,
		policyqosprofile.Setup,
		policysegment.Setup,
		policysegmentsecurityprofile.Setup,
		policyspoofguardprofile.Setup,
		policyvlansegment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
