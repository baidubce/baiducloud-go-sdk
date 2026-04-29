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

func getAddParsingRecordsUri(version string, zoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + zoneId + bce.URI_PREFIX + CONSTANT_RECORD
}
func getAssociateVpcUri(version string, zoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + zoneId
}
func getCreateAPrivateZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE
}
func getDeleteParsingRecordsUri(version string, recordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + recordId
}
func getDeletePrivateZoneUri(version string, zoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + zoneId
}
func getDisassociateVpcUri(version string, zoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + zoneId
}
func getModifyParsingRecordsUri(version string, recordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + recordId
}
func getQueryAndParseRecordListUri(version string, zoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + zoneId + bce.URI_PREFIX + CONSTANT_RECORD
}
func getQueryTheListOfPrivateZonesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE
}
func getSearchForDetailsOfPrivatzoneUri(version string, zoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + zoneId
}
func getSetParsingRecordStatusUri(version string, recordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + recordId
}
