package bcm

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bcm." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V3 = "v3"

	CONSTANT_BCM = "bcm"
)

// Client of bcm service is a kind of BceClient, so derived from BceClient
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

func getCreateAlarmTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getCreateInstanceGroupUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getCreateNotifyTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDeleteAlarmTemplatesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDeleteInstanceGroupUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDeleteInstanceGroupInstancesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDeleteNotifyTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeAlarmTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeAlarmTemplatesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeInstanceGroupUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeInstanceGroupsUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeNotifyTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeNotifyTemplatesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeReceiversUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getDescribeSystemTemplateRulesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getExportAlarmTemplatesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getImportAlarmTemplatesUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getUpdateAlarmTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getUpdateInstanceGroupUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
func getUpdateNotifyTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_BCM
}
