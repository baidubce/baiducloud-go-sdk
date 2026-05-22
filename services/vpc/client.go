package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vpc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VPC = "vpc"

	CONSTANT_SECURITY_GROUP = "securityGroup"

	CONSTANT_VPN = "vpn"

	CONSTANT_SSL_VPN_SERVER = "sslVpnServer"

	CONSTANT_SHUTDOWN_RELAY = "shutdownRelay"

	CONSTANT_VPNCONN = "vpnconn"

	CONSTANT_RULE = "rule"

	CONSTANT_RESOURCE_IP = "resourceIp"

	CONSTANT_CGW = "cgw"

	CONSTANT_PEERCONN = "peerconn"

	CONSTANT_ROUTE = "route"

	CONSTANT_PRIVATE_IP_ADDRESS_INFO = "privateIpAddressInfo"

	CONSTANT_ET_GATEWAY = "etGateway"

	CONSTANT_HEALTH_CHECK = "healthCheck"

	CONSTANT_ENTERPRISE = "enterprise"

	CONSTANT_SECURITY = "security"

	CONSTANT_SSL_VPN_USER = "sslVpnUser"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_GATEWAY = "gateway"

	CONSTANT_LIMITRULE = "limitrule"

	CONSTANT_UPDATE = "update"

	CONSTANT_IPRESERVE = "ipreserve"

	CONSTANT_OPEN_RELAY = "openRelay"

	CONSTANT_DELETE_PROTECT = "deleteProtect"
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
func getAuthorizeRegularSecurityGroupRulesV2Uri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getAuthorizedEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + EnterpriseSecurityGroupId
}
func getBatchCreateSslVpnUsersUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
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
func getCreateAPeerToPeerConnectionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
}
func getCreateARegularSecurityGroupV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
}
func getCreateDedicatedGatewayUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY
}
func getCreateDedicatedGatewayHealthCheckUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId + bce.URI_PREFIX + CONSTANT_HEALTH_CHECK
}
func getCreateEnterpriseSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY
}
func getCreateGatewayLimitRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE
}
func getCreateIpReservedUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getCreateRoutingRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
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
func getDeleteEnterpriseSecurityGroupUri(version string, EnterpriseSecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + EnterpriseSecurityGroupId
}
func getDeleteEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + EnterpriseSecurityGroupRuleId
}
func getDeleteGatewayLimitRuleUri(version string, GlrId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE + bce.URI_PREFIX + GlrId
}
func getDeleteIpReserveUri(version string, IpReserveId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE + bce.URI_PREFIX + IpReserveId
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
func getEnablePeerToPeerConnectionToSynchronizeDnsUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getGetVpcResourceIpInfoUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_RESOURCE_IP
}
func getListIpReserveUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getModifyGatewayLimitRulesUri(version string, GlrId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE + bce.URI_PREFIX + GlrId
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
func getQueryTheListOfEnterpriseSecurityGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY
}
func getQueryTheListOfPeerConnectionsUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
}
func getQueryTheListOfRegularSecurityGroupsV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
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
func getRejectPeerToPeerConnectionRequestUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseDedicatedGatewayUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getReleasePeerToPeerConnectionUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getRenewVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
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
func getUpdateDedicatedGatewayUri(version string, EtGatewayId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET_GATEWAY + bce.URI_PREFIX + EtGatewayId
}
func getUpdateEnterpriseSecurityGroupRulesUri(version string, EnterpriseSecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTERPRISE + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + EnterpriseSecurityGroupRuleId
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
