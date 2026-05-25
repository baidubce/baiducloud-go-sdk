package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vpc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VPC = "vpc"

	CONSTANT_SECURITY_GROUP = "securityGroup"

	CONSTANT_NAT = "nat"

	CONSTANT_DNAT_RULE = "dnatRule"

	CONSTANT_RESOURCE_IP = "resourceIp"

	CONSTANT_ENI = "eni"

	CONSTANT_IP_GROUP = "ipGroup"

	CONSTANT_PRIVATE_IP_ADDRESS_INFO = "privateIpAddressInfo"

	CONSTANT_SNAT_RULE = "snatRule"

	CONSTANT_HAVIP = "havip"

	CONSTANT_ET_GATEWAY = "etGateway"

	CONSTANT_IP_SET = "ipSet"

	CONSTANT_ENTERPRISE = "enterprise"

	CONSTANT_SECURITY = "security"

	CONSTANT_VPN = "vpn"

	CONSTANT_CGW = "cgw"

	CONSTANT_DELETE_IP_ADDRESS = "deleteIpAddress"

	CONSTANT_PEERCONN = "peerconn"

	CONSTANT_I_PV6_GATEWAY = "IPv6Gateway"

	CONSTANT_RATE_LIMIT_RULE = "rateLimitRule"

	CONSTANT_BATCH_CREATE = "batchCreate"

	CONSTANT_VPNCONN = "vpnconn"

	CONSTANT_STATUS = "status"

	CONSTANT_ACL = "acl"

	CONSTANT_RULE = "rule"

	CONSTANT_BIND_IP_SET = "bindIpSet"

	CONSTANT_PROBE = "probe"

	CONSTANT_GATEWAY = "gateway"

	CONSTANT_LIMITRULE = "limitrule"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_UPDATE = "update"

	CONSTANT_IPRESERVE = "ipreserve"

	CONSTANT_SSL_VPN_SERVER = "sslVpnServer"

	CONSTANT_PRIVATE_IP = "privateIp"

	CONSTANT_DELETE_PROTECT = "deleteProtect"

	CONSTANT_IP_ADDRESS = "ipAddress"

	CONSTANT_ROUTE = "route"

	CONSTANT_EGRESS_ONLY_RULE = "egressOnlyRule"

	CONSTANT_SHUTDOWN_RELAY = "shutdownRelay"

	CONSTANT_HEALTH_CHECK = "healthCheck"

	CONSTANT_SSL_VPN_USER = "sslVpnUser"

	CONSTANT_UNBIND_IP_SET = "unbindIpSet"

	CONSTANT_OPEN_RELAY = "openRelay"

	CONSTANT_BATCH_ADD = "batchAdd"

	CONSTANT_BATCH_DEL = "batchDel"
)

// Client of vpc service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getAcceptPeerToPeerConnectionApplicationsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getActiveStandbySwitchoverUri(version string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getAddAclRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACL + bce.URI_PREFIX + CONSTANT_RULE
}
func getAddElasticNetworkCardAuxiliaryIpUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP
}
func getAddIpAddressGroupToIpAddressFamilyUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId + bce.URI_PREFIX + CONSTANT_BIND_IP_SET
}
func getAddIpAddressesToTheIpAddressGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId + bce.URI_PREFIX + CONSTANT_IP_ADDRESS
}
func getAddIpv6OnlyOutboundAndNoInboundPolicyUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_EGRESS_ONLY_RULE
}
func getAuthorizeRegularSecurityGroupRulesV2Uri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getAuthorizedEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + EnterpriseSecurityGroupId
}
func getBatchAddDnatRulesUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE + bce.URI_PREFIX + CONSTANT_BATCH_CREATE
}
func getBatchAddSnatRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + CONSTANT_SNAT_RULE + bce.URI_PREFIX + CONSTANT_BATCH_CREATE
}
func getBatchCreateSslVpnUsersUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
}
func getBatchDeleteElasticNetworkCardIntranetIpUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP + bce.URI_PREFIX + CONSTANT_BATCH_DEL
}
func getBatchIncreaseElasticNetworkCardIntranetIpUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP + bce.URI_PREFIX + CONSTANT_BATCH_ADD
}
func getBindEipUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getBindPhysicalDedicatedLineUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getClosePeerToPeerConnectionToSynchronizeDnsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getCloseVpcRelayUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_SHUTDOWN_RELAY + bce.URI_PREFIX + VpcId
}
func getCreateAHighlyAvailableVirtualIpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP
}
func getCreateAPeerToPeerConnectionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
}
func getCreateARegularSecurityGroupV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
}
func getCreateAnIpv6GatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY
}
func getCreateDedicatedGatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY
}
func getCreateDedicatedGatewayHealthCheckUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId + bce.URI_PREFIX + CONSTANT_HEALTH_CHECK
}
func getCreateDnatRuleUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE
}
func getCreateElasticNetworkCardUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI
}
func getCreateEnterpriseSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY
}
func getCreateGatewayLimitRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE
}
func getCreateIpAddressFamilyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP
}
func getCreateIpAddressGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET
}
func getCreateIpReservedUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getCreateIpv6GatewaySpeedLimitPolicyUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE
}
func getCreateNatUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT
}
func getCreateNetworkDetectionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE
}
func getCreateRoutingRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
}
func getCreateSnatRuleUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_SNAT_RULE
}
func getCreateSslVpnServerUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER
}
func getCreateSubnetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getCreateUserGatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW
}
func getCreateVpcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC
}
func getCreateVpnUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN
}
func getCreateVpnTunnelUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_VPNCONN
}
func getDeleteAclRuleUri(version string, AclRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACL + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + AclRuleId
}
func getDeleteDnatRuleUri(version string, NatId string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE + bce.URI_PREFIX + RuleId
}
func getDeleteElasticNetworkCardAuxiliaryIpUri(version string, EniId string, PrivateIpAddress string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP + bce.URI_PREFIX + PrivateIpAddress
}
func getDeleteEnterpriseSecurityGroupUri(version string, EnterpriseSecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + EnterpriseSecurityGroupId
}
func getDeleteEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + EnterpriseSecurityGroupRuleId
}
func getDeleteGatewayLimitRuleUri(version string, GlrId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE + bce.URI_PREFIX + GlrId
}
func getDeleteHighlyAvailableVirtualIpUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getDeleteIpAddressFamilyUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId
}
func getDeleteIpAddressFromIpAddressGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId + bce.URI_PREFIX + CONSTANT_DELETE_IP_ADDRESS
}
func getDeleteIpAddressGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId
}
func getDeleteIpReserveUri(version string, IpReserveId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE + bce.URI_PREFIX + IpReserveId
}
func getDeleteIpv6GatewayUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId
}
func getDeleteIpv6GatewaySpeedLimitPolicyUri(version string, GatewayId string, RateLimitRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE + bce.URI_PREFIX + RateLimitRuleId
}
func getDeleteIpv6OnlyAccessPolicyUri(version string, GatewayId string, EgressOnlyRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_EGRESS_ONLY_RULE + bce.URI_PREFIX + EgressOnlyRuleId
}
func getDeleteNetworkDetectionUri(version string, ProbeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE + bce.URI_PREFIX + ProbeId
}
func getDeleteRegularSecurityGroupRulesV2Uri(version string, SecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + SecurityGroupRuleId
}
func getDeleteRegularSecurityGroupV2Uri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getDeleteRoutingRulesUri(version string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getDeleteSnatRuleUri(version string, NatId string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_SNAT_RULE + bce.URI_PREFIX + RuleId
}
func getDeleteSslVpnServerUri(version string, VpnId string, SslVpnServerId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER + bce.URI_PREFIX + SslVpnServerId
}
func getDeleteSslVpnUserUri(version string, VpnId string, UserId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER + bce.URI_PREFIX + UserId
}
func getDeleteSubnetUri(version string, SubnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + SubnetId
}
func getDeleteUserGatewayUri(version string, CgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW + bce.URI_PREFIX + CgwId
}
func getDeleteVpcUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + VpcId
}
func getDeleteVpnTunnelUri(version string, VpnConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_VPNCONN + bce.URI_PREFIX + VpnConnId
}
func getElasticNetworkCardBindingEipUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getElasticNetworkCardMountedCloudProductInstanceUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getElasticNetworkCardUnbindingEipUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getElasticNetworkCardUninstallationCloudProductInstanceUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getElasticNetworkCardUpdateEnterpriseSecurityGroupUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getElasticNetworkCardUpdatesRegularSecurityGroupUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getEnablePeerToPeerConnectionToSynchronizeDnsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getGetNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getGetVpcResourceIpInfoUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_RESOURCE_IP
}
func getHighAvailabilityVirtualIpUnbindingEipUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getHighAvailabilityVirtualIpUnbindingInstanceUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getHighlyAvailableVirtualIpBindingEipUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getHighlyAvailableVirtualIpBindingInstanceUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getIpv6GatewayBandwidthUpgradeAndDowngradeUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId
}
func getListDnatRuleUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE
}
func getListIpReserveUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getListNatUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT
}
func getListSnatRuleUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_SNAT_RULE
}
func getModifyGatewayLimitRulesUri(version string, GlrId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE + bce.URI_PREFIX + GlrId
}
func getModifyNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getNatBindEipUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getNatUnBindEipUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getOpenVpcRelayUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_OPEN_RELAY + bce.URI_PREFIX + VpcId
}
func getPeerToPeerConnectionBandwidthUpgradeAndDowngradeUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getPeerToPeerConnectionRenewalUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getPrepaidPeerToPeerConnectionUnsubscribeUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getPurchaseReservedNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getQueryAclUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACL
}
func getQueryAclRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACL + bce.URI_PREFIX + CONSTANT_RULE
}
func getQueryIpAddressFamilyListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP
}
func getQueryIpv6GatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY
}
func getQueryNetworkDetectionDetailsUri(version string, ProbeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE + bce.URI_PREFIX + ProbeId
}
func getQueryNetworkDetectionListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE
}
func getQueryRoutingRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
}
func getQueryRoutingTableUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE
}
func getQuerySpecifiedSubnetUri(version string, SubnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + SubnetId
}
func getQuerySpecifiedVpcUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + VpcId
}
func getQuerySslVpnServerUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER
}
func getQuerySslVpnUsersUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
}
func getQuerySubnetListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getQueryTheDetailsOfTheDedicatedGatewayUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getQueryTheListOfDedicatedLineGatewaysUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY
}
func getQueryTheListOfElasticNetworkCardsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI
}
func getQueryTheListOfEnterpriseSecurityGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY
}
func getQueryTheListOfHighlyAvailableVirtualIpsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP
}
func getQueryTheListOfIpAddressGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET
}
func getQueryTheListOfPeerConnectionsUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
}
func getQueryTheListOfRegularSecurityGroupsV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
}
func getQueryTheListOfSpeedLimitPoliciesForIpv6GatewayUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE
}
func getQueryTheSpecifiedElasticNetworkCardUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getQueryTheSpecifiedHighlyAvailableVirtualIpUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getQueryTheSpecifiedIpAddressFamilyUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId
}
func getQueryTheSpecifiedIpAddressGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId
}
func getQueryTheStatusOfTheElasticNetworkCardUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_STATUS
}
func getQueryVpcIntranetIpUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + VpcId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP_ADDRESS_INFO
}
func getQueryVpcListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC
}
func getQueryVpnListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN
}
func getQueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_EGRESS_ONLY_RULE
}
func getRejectPeerToPeerConnectionRequestUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseDedicatedGatewayUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getReleaseNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getReleasePeerToPeerConnectionUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getRemoveElasticNetworkCardUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getRemoveIpAddressGroupFromIpAddressFamilyUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId + bce.URI_PREFIX + CONSTANT_UNBIND_IP_SET
}
func getRenewVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getResizeNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getRevokeRegularSecurityGroupRulesV2Uri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getSearchForVpnDetailsUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getSearchVpnTunnelUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_VPNCONN + bce.URI_PREFIX + VpnId
}
func getUnbindEipUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getUnbindPhysicalDedicatedLineUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getUpdateAclRulesUri(version string, AclRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACL + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + AclRuleId
}
func getUpdateDedicatedGatewayUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getUpdateDnatRuleUri(version string, NatId string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE + bce.URI_PREFIX + RuleId
}
func getUpdateElasticNetworkCardUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getUpdateEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + EnterpriseSecurityGroupRuleId
}
func getUpdateHighlyAvailableVirtualIpUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getUpdateIpAddressFamilyUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId
}
func getUpdateIpAddressGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId
}
func getUpdateIpv6GatewayReleaseProtectionSwitchUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateIpv6GatewaySpeedLimitPolicyUri(version string, GatewayId string, RateLimitRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE + bce.URI_PREFIX + RateLimitRuleId
}
func getUpdateNatReleaseProtectionSwitchUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateNetworkDetectionUri(version string, ProbeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE + bce.URI_PREFIX + ProbeId
}
func getUpdatePeerToPeerConnectionReleaseProtectionSwitchUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateRegularSecurityGroupRulesV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateRoutingRulesUri(version string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getUpdateSnatRuleUri(version string, NatId string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_SNAT_RULE + bce.URI_PREFIX + RuleId
}
func getUpdateSslVpnServerUri(version string, VpnId string, SslVpnServerId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER + bce.URI_PREFIX + SslVpnServerId
}
func getUpdateSslVpnUsersUri(version string, VpnId string, UserId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER + bce.URI_PREFIX + UserId
}
func getUpdateSubnetUri(version string, SubnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + SubnetId
}
func getUpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getUpdateUserGatewayUri(version string, CgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW + bce.URI_PREFIX + CgwId
}
func getUpdateVpcUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + VpcId
}
func getUpdateVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getUpdateVpnReleaseProtectionUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateVpnTunnelUri(version string, VpnConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_VPNCONN + bce.URI_PREFIX + VpnConnId
}
func getUserGatewayDetailsUri(version string, CgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW + bce.URI_PREFIX + CgwId
}
func getUserGatewayListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW
}
func getViewGatewayLimitRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE
}
func getViewPeerToPeerConnectionDetailsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getViewSecurityGroupDetailsV2Uri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
