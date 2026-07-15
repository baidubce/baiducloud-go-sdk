package cloudassistant

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "cloudassistant." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_CA = "ca"

	CONSTANT_ACTION = "action"

	CONSTANT_ACTION_RUN = "actionRun"

	CONSTANT_AGENT = "agent"

	CONSTANT_BATCH = "batch"

	CONSTANT_LIST = "list"

	CONSTANT_LOG = "log"
)

// Client of cloudassistant service is a kind of BceClient, so derived from BceClient
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

func getActionListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION + bce.URI_PREFIX + CONSTANT_LIST
}
func getActionLogUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_LOG
}
func getActionRunUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION_RUN
}
func getActionRunListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION_RUN + bce.URI_PREFIX + CONSTANT_LIST
}
func getBatchGetAgentUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_AGENT + bce.URI_PREFIX + CONSTANT_BATCH
}
func getCreateActionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION
}
func getDeleteActionUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION + bce.URI_PREFIX + Id
}
func getGetActionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION
}
func getGetActionRunUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION_RUN
}
func getUpdateActionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CA + bce.URI_PREFIX + CONSTANT_ACTION
}
