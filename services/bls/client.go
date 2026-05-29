package bls

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bls." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_PROJECT = "project"

	CONSTANT_LOGSTORE = "logstore"

	CONSTANT_LOGRECORD = "logrecord"

	CONSTANT_LIST = "list"

	CONSTANT_LOGHISTOGRAM = "loghistogram"
)

// Client of bls service is a kind of BceClient, so derived from BceClient
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

func getCreateProjectUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT
}
func getDeleteProjectUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT + bce.URI_PREFIX + Uuid
}
func getDescribeProjectUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT + bce.URI_PREFIX + Uuid
}
func getListProjectUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT + bce.URI_PREFIX + CONSTANT_LIST
}
func getPullLogRecordUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGRECORD
}
func getPushLogRecordUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGRECORD
}
func getQueryLogHistogramUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGHISTOGRAM
}
func getQueryLogRecordUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGRECORD
}
func getUpdateProjectUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT
}
