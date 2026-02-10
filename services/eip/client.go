package eip

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "eip." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_TBSP = "tbsp"

	CONSTANT_AREA_BLOCKING = "areaBlocking"

	CONSTANT_EIP = "eip"

	CONSTANT_PRICE = "price"

	CONSTANT_PROTOCOL_BLOCKING = "protocolBlocking"

	CONSTANT_EIPBP = "eipbp"

	CONSTANT_IP_PROTECT_LEVEL = "ipProtectLevel"

	CONSTANT_IP_CLEAN = "ipClean"

	CONSTANT_IP_WHITELIST = "ipWhitelist"

	CONSTANT_EIPTP = "eiptp"

	CONSTANT_EIPGROUP = "eipgroup"
)

// Client of eip service is a kind of BceClient, so derived from BceClient
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

func getAddTbspAreaBlockingUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_AREA_BLOCKING
}
func getAddTbspIpWhitelistUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_WHITELIST
}
func getAddTbspProtocolBlockingUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_PROTOCOL_BLOCKING
}
func getBandwidthPackageInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getBindTbspProtectionObjectUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id
}
func getCreateTbspUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP
}
func getDetailTbspUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id
}
func getDisableTbspIpCleanUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getEipInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getEnableTbspIpCleanUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getListTbspUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP
}
func getListTbspAreaBlockingUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_AREA_BLOCKING
}
func getListTbspIpCleanUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getListTbspIpWhitelistUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_WHITELIST
}
func getListTbspProtocolBlockingUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_PROTOCOL_BLOCKING
}
func getModifyTbspIpCleanThresholdUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getModifyTbspIpProtectLevelUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_PROTECT_LEVEL
}
func getRemoveTbspAreaBlockingUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_AREA_BLOCKING
}
func getRemoveTbspIpWhitelistUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_IP_WHITELIST
}
func getRemoveTbspProtocolBlockingUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id + bce.URI_PREFIX + CONSTANT_PROTOCOL_BLOCKING
}
func getRenewTbspUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id
}
func getResizeTbspUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id
}
func getSharedBandwidthInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getSharedDataPackageInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPTP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getUnbindTbspProtectionObjectUri(version string, id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + id
}
