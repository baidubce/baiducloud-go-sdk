package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vpc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VPC = "vpc"

	CONSTANT_VPN = "vpn"

	CONSTANT_VPNCONN = "vpnconn"

	CONSTANT_SSL_VPN_SERVER = "sslVpnServer"

	CONSTANT_SHUTDOWN_RELAY = "shutdownRelay"

	CONSTANT_RESOURCE_IP = "resourceIp"

	CONSTANT_CGW = "cgw"

	CONSTANT_SSL_VPN_USER = "sslVpnUser"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_PRIVATE_IP_ADDRESS_INFO = "privateIpAddressInfo"

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

func getBatchCreateSslVpnUsersUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
}
func getBindEipUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId
}
func getCloseVpcRelayUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_SHUTDOWN_RELAY + bce.URI_PREFIX + vpcId
}
func getCreateIpReservedUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getCreateSslVpnServerUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER
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
func getCreateVpnTunnelUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_VPNCONN
}
func getDeleteIpReserveUri(version string, ipReserveId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE + bce.URI_PREFIX + ipReserveId
}
func getDeleteSslVpnServerUri(version string, vpnId string, sslVpnServerId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER + bce.URI_PREFIX + sslVpnServerId
}
func getDeleteSslVpnUserUri(version string, vpnId string, userId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER + bce.URI_PREFIX + userId
}
func getDeleteSubnetUri(version string, subnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + subnetId
}
func getDeleteUserGatewayUri(version string, cgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW + bce.URI_PREFIX + cgwId
}
func getDeleteVpcUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId
}
func getDeleteVpnTunnelUri(version string, vpnConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_VPNCONN + bce.URI_PREFIX + vpnConnId
}
func getGetVpcResourceIpInfoUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_RESOURCE_IP
}
func getListIpReserveUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getOpenVpcRelayUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_OPEN_RELAY + bce.URI_PREFIX + vpcId
}
func getQuerySpecifiedSubnetUri(version string, subnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + subnetId
}
func getQuerySpecifiedVpcUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId
}
func getQuerySslVpnServerUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER
}
func getQuerySslVpnUsersUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER
}
func getQuerySubnetListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getQueryVpcIntranetIpUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP_ADDRESS_INFO
}
func getQueryVpcListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC
}
func getQueryVpnListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN
}
func getReleaseVpnUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId
}
func getRenewVpnUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId
}
func getSearchForVpnDetailsUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId
}
func getSearchVpnTunnelUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_VPNCONN + bce.URI_PREFIX + vpnId
}
func getUnbindEipUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId
}
func getUpdateSslVpnServerUri(version string, vpnId string, sslVpnServerId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_SERVER + bce.URI_PREFIX + sslVpnServerId
}
func getUpdateSslVpnUsersUri(version string, vpnId string, userId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_SSL_VPN_USER + bce.URI_PREFIX + userId
}
func getUpdateSubnetUri(version string, subnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + subnetId
}
func getUpdateUserGatewayUri(version string, cgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW + bce.URI_PREFIX + cgwId
}
func getUpdateVpcUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId
}
func getUpdateVpnUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId
}
func getUpdateVpnReleaseProtectionUri(version string, vpnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + vpnId + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateVpnTunnelUri(version string, vpnConnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_VPNCONN + bce.URI_PREFIX + vpnConnId
}
func getUserGatewayDetailsUri(version string, cgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW + bce.URI_PREFIX + cgwId
}
func getUserGatewayListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPN + bce.URI_PREFIX + CONSTANT_CGW
}
