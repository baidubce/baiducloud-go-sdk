package eip

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "eip." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_EIP = "eip"

	CONSTANT_PRICE = "price"

	CONSTANT_TBSP = "tbsp"

	CONSTANT_TRANSFER = "transfer"

	CONSTANT_EIPGROUP = "eipgroup"

	CONSTANT_EIPBP = "eipbp"

	CONSTANT_PROTOCOL_BLOCKING = "protocolBlocking"

	CONSTANT_AREA_BLOCKING = "areaBlocking"

	CONSTANT_DDOS = "ddos"

	CONSTANT_RECYCLE = "recycle"

	CONSTANT_UNBAN = "unban"

	CONSTANT_RECORD = "record"

	CONSTANT_IP_CLEAN = "ipClean"

	CONSTANT_EIPTP = "eiptp"

	CONSTANT_IP_WHITELIST = "ipWhitelist"

	CONSTANT_REFUND = "refund"

	CONSTANT_IP_PROTECT_LEVEL = "ipProtectLevel"

	CONSTANT_DELETE_PROTECT = "deleteProtect"
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

func getAddEipGroupCountUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getAddTbspAreaBlockingUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_AREA_BLOCKING
}
func getAddTbspIpWhitelistUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_WHITELIST
}
func getAddTbspProtocolBlockingUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_PROTOCOL_BLOCKING
}
func getApplyForEipUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP
}
func getBandwidthPackageInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getBindEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getBindTbspProtectionObjectUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id
}
func getCancelEipTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TRANSFER
}
func getCreateASharedTrafficPackageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPTP
}
func getCreateEipBpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP
}
func getCreateEipGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP
}
func getCreateEipTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TRANSFER
}
func getCreateTbspUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP
}
func getDetailTbspUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id
}
func getDirectEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getDisableTbspIpCleanUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getEipBandwidthScalingCapacityUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getEipInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getEipPostpaidToPrepaidUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getEipRenewalUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getEnableTbspIpCleanUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getGetEipBpUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + Id
}
func getGetEipGroupUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getListBaseDdosUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DDOS
}
func getListBaseDdosAttackRecordUri(version string, Ip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DDOS + bce.URI_PREFIX + Ip + bce.URI_PREFIX + CONSTANT_RECORD
}
func getListEipBpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP
}
func getListEipGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP
}
func getListEipTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TRANSFER
}
func getListRecycleEipsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + CONSTANT_RECYCLE
}
func getListTbspUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP
}
func getListTbspAreaBlockingUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_AREA_BLOCKING
}
func getListTbspIpCleanUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getListTbspIpWhitelistUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_WHITELIST
}
func getListTbspProtocolBlockingUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_PROTOCOL_BLOCKING
}
func getListUnbanUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_UNBAN + bce.URI_PREFIX + CONSTANT_RECORD
}
func getModifyTbspIpCleanThresholdUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_CLEAN
}
func getModifyTbspIpProtectLevelUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_PROTECT_LEVEL
}
func getMoveInEipsUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getMoveOutEipsUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getOptionalReleaseEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getPurchaseReservedEipGroupUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getQueryEipListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP
}
func getQueryTheDetailsOfSharedTrafficPackagesUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPTP + bce.URI_PREFIX + Id
}
func getQueryTheListOfSharedTrafficPackagesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPTP
}
func getReceiveEipTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TRANSFER
}
func getRefundEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + CONSTANT_REFUND + bce.URI_PREFIX + Eip
}
func getRefundEipGroupUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + CONSTANT_REFUND + bce.URI_PREFIX + Id
}
func getRejectEipTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TRANSFER
}
func getReleaseEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getReleaseEipBpUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + Id
}
func getReleaseEipFromRecycleUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + CONSTANT_RECYCLE + bce.URI_PREFIX + Eip
}
func getReleaseEipGroupUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getRemoveTbspAreaBlockingUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_AREA_BLOCKING
}
func getRemoveTbspIpWhitelistUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_IP_WHITELIST
}
func getRemoveTbspProtocolBlockingUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id + bce.URI_PREFIX + CONSTANT_PROTOCOL_BLOCKING
}
func getRenewTbspUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id
}
func getResizeEipBpBandwidthUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + Id
}
func getResizeEipGroupBandwidthUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
func getResizeTbspUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id
}
func getRestoreEipFromRecycleUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + CONSTANT_RECYCLE + bce.URI_PREFIX + Eip
}
func getSharedBandwidthInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getSharedDataPackageInquiryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPTP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getStartEipAutoRenewUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getStopEipAutoRenewUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getUnDirectEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getUnbindEipUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip
}
func getUnbindTbspProtectionObjectUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TBSP + bce.URI_PREFIX + Id
}
func getUpdateBaseDdosThresholdUri(version string, Ip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DDOS + bce.URI_PREFIX + Ip
}
func getUpdateEipBpAutoReleaseTimeUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + Id
}
func getUpdateEipBpNameUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPBP + bce.URI_PREFIX + Id
}
func getUpdateEipDeleteProtectUri(version string, Eip string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIP + bce.URI_PREFIX + Eip + bce.URI_PREFIX + CONSTANT_DELETE_PROTECT
}
func getUpdateEipGroupUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EIPGROUP + bce.URI_PREFIX + Id
}
