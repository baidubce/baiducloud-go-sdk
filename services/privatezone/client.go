package privatezone

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "privatezone." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_PRIVATEZONE = "privatezone"

	CONSTANT_RECORD = "record"
)

// Client of privatezone service is a kind of BceClient, so derived from BceClient
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

func getAddRecordUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId + bce.URI_PREFIX + CONSTANT_RECORD
}
func getBindVpcUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getCreatePrivateZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE
}
func getDeletePrivateZoneUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getDeleteRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getDisableRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getEnableRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getGetPrivateZoneUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getListPrivateZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE
}
func getListRecordUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId + bce.URI_PREFIX + CONSTANT_RECORD
}
func getUnbindVpcUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getUpdateRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
