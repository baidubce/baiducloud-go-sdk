package dns

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "dns." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_DNS = "dns"

	CONSTANT_ZONE = "zone"

	CONSTANT_RECORD = "record"

	CONSTANT_CUSTOMLINE = "customline"

	CONSTANT_ORDER = "order"
)

// Client of dns service is a kind of BceClient, so derived from BceClient
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

func getAddDomainNameUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE
}
func getAddLineGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE
}
func getAddParsingRecordsUri(version string, ZoneName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD
}
func getDeleteLineGroupUri(version string, LineId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE + bce.URI_PREFIX + LineId
}
func getDeleteParsingRecordsUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getDomainNameRenewalUri(version string, Name string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + CONSTANT_ORDER + bce.URI_PREFIX + Name
}
func getModifyParsingRecordsUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getModifyTheParsingRecordStatusUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getPurchaseAPaidDomainNameUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + CONSTANT_ORDER
}
func getQueryAndParseRecordListUri(version string, ZoneName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD
}
func getQueryDomainNameListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE
}
func getQueryTheListOfLineGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE
}
func getRemoveDomainNameUri(version string, ZoneName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName
}
func getUpdateLineGroupUri(version string, LineId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE + bce.URI_PREFIX + LineId
}
func getUpgradeTheFreeDomainNameToTheUniversalVersionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + CONSTANT_ORDER
}
