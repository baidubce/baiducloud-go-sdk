package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vpc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VPC = "vpc"

	CONSTANT_VPN = "vpn"

	CONSTANT_SSL_VPN_SERVER = "sslVpnServer"

	CONSTANT_SHUTDOWN_RELAY = "shutdownRelay"

	CONSTANT_VPNCONN = "vpnconn"

	CONSTANT_RESOURCE_IP = "resourceIp"

	CONSTANT_CGW = "cgw"

	CONSTANT_PEERCONN = "peerconn"

	CONSTANT_PRIVATE_IP_ADDRESS_INFO = "privateIpAddressInfo"

	CONSTANT_SSL_VPN_USER = "sslVpnUser"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_GATEWAY = "gateway"

	CONSTANT_LIMITRULE = "limitrule"

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
func getBatchCreateSslVpnUsersUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
}
func getBindEipUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
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
func getCreateGatewayLimitRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE
}
func getCreateIpReservedUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
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
func getDeleteGatewayLimitRuleUri(version string, GlrId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GATEWAY + bce.URI_PREFIX + CONSTANT_LIMITRULE + bce.URI_PREFIX + GlrId
}
func getDeleteIpReserveUri(version string, IpReserveId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE + bce.URI_PREFIX + IpReserveId
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
func getQueryTheListOfPeerConnectionsUri(version string, VpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN
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
func getReleasePeerToPeerConnectionUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId
}
func getReleaseVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
}
func getRenewVpnUri(version string, VpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + VpnId
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
func getUpdatePeerToPeerConnectionReleaseProtectionSwitchUri(version string, PeerConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PEERCONN + bce.URI_PREFIX + PeerConnId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
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
