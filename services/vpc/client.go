package vpc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vpc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VPC = "vpc"

	CONSTANT_SHUTDOWN_RELAY = "shutdownRelay"

	CONSTANT_RESOURCE_IP = "resourceIp"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_IPRESERVE = "ipreserve"

	CONSTANT_OPEN_RELAY = "openRelay"

	CONSTANT_PRIVATE_IP_ADDRESS_INFO = "privateIpAddressInfo"
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

func getCloseVpcRelayUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_SHUTDOWN_RELAY + bce.URI_PREFIX + vpcId
}
func getCreateAReservedNetworkSegmentUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getCreateSubnetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getCreateVpcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC
}
func getDeleteReservedNetworkSegmentUri(version string, ipReserveId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE + bce.URI_PREFIX + ipReserveId
}
func getDeleteSubnetUri(version string, subnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + subnetId
}
func getDeleteVpcUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId
}
func getEnableVpcRelayUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_OPEN_RELAY + bce.URI_PREFIX + vpcId
}
func getQuerySpecifiedSubnetUri(version string, subnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + subnetId
}
func getQuerySpecifiedVpcUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId
}
func getQuerySubnetListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getQueryTheIpAddressesOccupiedByProductsWithinVpcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_RESOURCE_IP
}
func getQueryTheReservedNetworkSegmentListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_IPRESERVE
}
func getQueryVpcIntranetIpUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId + bce.URI_PREFIX + CONSTANT_PRIVATE_IP_ADDRESS_INFO
}
func getQueryVpcListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC
}
func getUpdateSubnetUri(version string, subnetId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + subnetId
}
func getUpdateVpcUri(version string, vpcId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + vpcId
}
