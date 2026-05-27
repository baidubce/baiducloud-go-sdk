package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vpc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VPC = "vpc"

	CONSTANT_NAT = "nat"

	CONSTANT_DNAT_RULE = "dnatRule"

	CONSTANT_IP_GROUP = "ipGroup"

	CONSTANT_UNBIND_IP_SET = "unbindIpSet"

	CONSTANT_SECURITY_GROUP = "securityGroup"

	CONSTANT_RULE = "rule"

	CONSTANT_RESOURCE_IP = "resourceIp"

	CONSTANT_UPDATE = "update"

	CONSTANT_IP_SET = "ipSet"

	CONSTANT_PRIVATE_IP_ADDRESS_INFO = "privateIpAddressInfo"

	CONSTANT_ENI = "eni"

	CONSTANT_SNAT_RULE = "snatRule"

	CONSTANT_ET_GATEWAY = "etGateway"

	CONSTANT_VPN = "vpn"

	CONSTANT_CGW = "cgw"

	CONSTANT_HAVIP = "havip"

	CONSTANT_PEERCONN = "peerconn"

	CONSTANT_BATCH_CREATE = "batchCreate"

	CONSTANT_VPNCONN = "vpnconn"

	CONSTANT_ENTERPRISE = "enterprise"

	CONSTANT_SECURITY = "security"

	CONSTANT_ACL = "acl"

	CONSTANT_GATEWAY = "gateway"

	CONSTANT_LIMITRULE = "limitrule"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_I_PV6_GATEWAY = "IPv6Gateway"

	CONSTANT_EGRESS_ONLY_RULE = "egressOnlyRule"

	CONSTANT_IPRESERVE = "ipreserve"

	CONSTANT_SSL_VPN_SERVER = "sslVpnServer"

	CONSTANT_DELETE_PROTECT = "deleteProtect"

	CONSTANT_PRIVATE_IP = "privateIp"

	CONSTANT_BATCH_ADD = "batchAdd"

	CONSTANT_ROUTE = "route"

	CONSTANT_RATE_LIMIT_RULE = "rateLimitRule"

	CONSTANT_PROBE = "probe"

	CONSTANT_SHUTDOWN_RELAY = "shutdownRelay"

	CONSTANT_HEALTH_CHECK = "healthCheck"

	CONSTANT_SSL_VPN_USER = "sslVpnUser"

	CONSTANT_DELETE_IP_ADDRESS = "deleteIpAddress"

	CONSTANT_BIND_IP_SET = "bindIpSet"

	CONSTANT_BATCH_DEL = "batchDel"

	CONSTANT_OPEN_RELAY = "openRelay"

	CONSTANT_STATUS = "status"

	CONSTANT_IP_ADDRESS = "ipAddress"
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

func getAcceptPeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getActiveStandbySwitchoverUri(version string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getAddAclRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACL + bce.URI_PREFIX + CONSTANT_RULE
}
func getAddEniIpUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP
}
func getAddIpAddressToIpGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId + bce.URI_PREFIX + CONSTANT_IP_ADDRESS
}
func getAddIpGroupToIpSetUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId + bce.URI_PREFIX + CONSTANT_BIND_IP_SET
}
func getAttachEniInstanceUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getAuthorizeEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + EnterpriseSecurityGroupId
}
func getAuthorizeSecurityGroupRulesUri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getBatchAddDnatRulesUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE + bce.URI_PREFIX + CONSTANT_BATCH_CREATE
}
func getBatchAddEniIpUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP + bce.URI_PREFIX + CONSTANT_BATCH_ADD
}
func getBatchAddSnatRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + CONSTANT_SNAT_RULE + bce.URI_PREFIX + CONSTANT_BATCH_CREATE
}
func getBatchCreateSslVpnUsersUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
}
func getBatchDeleteEniIpUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP + bce.URI_PREFIX + CONSTANT_BATCH_DEL
}
func getBindEipUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getBindEniEipUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getBindHaVipEipUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getBindHaVipInstanceUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getBindPhysicalDedicatedLineUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getClosePeerConnSyncDnsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getCloseVpcRelayUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_SHUTDOWN_RELAY + bce.URI_PREFIX + VpcId
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
func getCreateEgressOnlyRuleUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_EGRESS_ONLY_RULE
}
func getCreateEniUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI
}
func getCreateEnterpriseSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY
}
func getCreateGatewayLimitRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE
}
func getCreateHaVipUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP
}
func getCreateIpGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET
}
func getCreateIpReservedUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getCreateIpSetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP
}
func getCreateIpv6GatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY
}
func getCreateNatUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT
}
func getCreatePeerConnUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
}
func getCreateProbeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE
}
func getCreateRateLimitRuleUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE
}
func getCreateRoutingRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
}
func getCreateSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
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
func getDeleteEniIpUri(version string, EniId string, PrivateIpAddress string) string {
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
func getDeleteHaVipUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getDeleteIpGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId
}
func getDeleteIpReserveUri(version string, IpReserveId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE + bce.URI_PREFIX + IpReserveId
}
func getDeleteIpSetUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId
}
func getDeleteIpv6GatewayUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId
}
func getDeleteIpv6GatewayEgressOnlyRuleUri(version string, GatewayId string, EgressOnlyRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_EGRESS_ONLY_RULE + bce.URI_PREFIX + EgressOnlyRuleId
}
func getDeleteIpv6GatewayRateLimitRuleUri(version string, GatewayId string, RateLimitRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE + bce.URI_PREFIX + RateLimitRuleId
}
func getDeleteProbeUri(version string, ProbeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE + bce.URI_PREFIX + ProbeId
}
func getDeleteRoutingRulesUri(version string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getDeleteSecurityGroupUri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getDeleteSecurityGroupRulesUri(version string, SecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + SecurityGroupRuleId
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
func getDetachEniInstanceUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getGetEniDetailUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getGetEniStatusUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId + bce.URI_PREFIX + CONSTANT_STATUS
}
func getGetHaVipDetailUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getGetNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getGetPeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getGetProbeDetailUri(version string, ProbeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE + bce.URI_PREFIX + ProbeId
}
func getGetSecurityGroupDetailsUri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getGetVpcResourceIpInfoUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_RESOURCE_IP
}
func getListDnatRuleUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE
}
func getListEgressOnlyRuleUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_EGRESS_ONLY_RULE
}
func getListEniUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI
}
func getListHaVipUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP
}
func getListIpReserveUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getListNatUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT
}
func getListPeerConnUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
}
func getListProbesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE
}
func getListRateLimitRuleUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE
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
func getOpenPeerConnSyncDnsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getOpenVpcRelayUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_OPEN_RELAY + bce.URI_PREFIX + VpcId
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
func getQueryEnterpriseSecurityGroupListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY
}
func getQueryIpGroupDetailUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId
}
func getQueryIpGroupListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET
}
func getQueryIpSetDetailUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId
}
func getQueryIpSetListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP
}
func getQueryIpv6GatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY
}
func getQueryRoutingRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
}
func getQueryRoutingTableUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE
}
func getQuerySecurityGroupsListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
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
func getQueryVpcIntranetIpUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + VpcId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP_ADDRESS_INFO
}
func getQueryVpcListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC
}
func getQueryVpnListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN
}
func getRefundPeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getRejectPeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseDedicatedGatewayUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getReleaseNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getReleasePeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getRemoveEniUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getRemoveIpAddressFromIpGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId + bce.URI_PREFIX + CONSTANT_DELETE_IP_ADDRESS
}
func getRemoveIpGroupFromIpSetUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId + bce.URI_PREFIX + CONSTANT_UNBIND_IP_SET
}
func getRenewPeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getRenewVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getResizeIpv6GatewayUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId
}
func getResizeNatUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId
}
func getRevokeSecurityGroupRulesUri(version string, SecurityGroupId string) string {
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
func getUnbindEniEipUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getUnbindHaVipEipUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getUnbindHaVipInstanceUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
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
func getUpdateDeleteProtectUri(version string, GatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateDnatRuleUri(version string, NatId string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DNAT_RULE + bce.URI_PREFIX + RuleId
}
func getUpdateEniUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getUpdateEniEnterpriseSecurityGroupUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getUpdateEniSecurityGroupUri(version string, EniId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + EniId
}
func getUpdateEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + EnterpriseSecurityGroupRuleId
}
func getUpdateHaVipUri(version string, HaVipId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_HAVIP + bce.URI_PREFIX + HaVipId
}
func getUpdateIpGroupUri(version string, IpSetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_SET + bce.URI_PREFIX + IpSetId
}
func getUpdateIpSetUri(version string, IpGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IP_GROUP + bce.URI_PREFIX + IpGroupId
}
func getUpdateNatReleaseProtectionSwitchUri(version string, NatId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_NAT + bce.URI_PREFIX + NatId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdatePeerConnUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getUpdatePeerConnBandwidthUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getUpdatePeerConnDeleteProtectUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateProbeUri(version string, ProbeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROBE + bce.URI_PREFIX + ProbeId
}
func getUpdateRateLimitRuleUri(version string, GatewayId string, RateLimitRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_I_PV6_GATEWAY + bce.URI_PREFIX + GatewayId + bce.URI_PREFIX + CONSTANT_RATE_LIMIT_RULE + bce.URI_PREFIX + RateLimitRuleId
}
func getUpdateRoutingRulesUri(version string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getUpdateSecurityGroupRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + CONSTANT_UPDATE
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
