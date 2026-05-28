package dns

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "dns." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_DNS = "dns"

	CONSTANT_ZONE = "zone"

	CONSTANT_ORDER = "order"

	CONSTANT_CUSTOMLINE = "customline"

	CONSTANT_RECORD = "record"
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

func getAddLineGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE
}
func getCreatePaidZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + CONSTANT_ORDER
}
func getCreateRecordUri(version string, ZoneName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD
}
func getCreateZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE
}
func getDeleteLineGroupUri(version string, LineId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE + bce.URI_PREFIX + LineId
}
func getDeleteRecordUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getDeleteZoneUri(version string, ZoneName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName
}
func getListLineGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE
}
func getListRecordUri(version string, ZoneName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD
}
func getListZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE
}
func getRenewZoneUri(version string, Name string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + CONSTANT_ORDER + bce.URI_PREFIX + Name
}
func getUpdateLineGroupUri(version string, LineId int32) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_CUSTOMLINE + bce.URI_PREFIX + LineId
}
func getUpdateRecordUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getUpdateRecordDisableUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getUpdateRecordEnableUri(version string, ZoneName string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + ZoneName + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getUpgradeZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DNS + bce.URI_PREFIX + CONSTANT_ZONE + bce.URI_PREFIX + CONSTANT_ORDER
}
