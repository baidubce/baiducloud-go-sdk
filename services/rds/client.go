package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "rds." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_V1 = "v1"

	CONSTANT_TASK = "task"
)

// Client of rds service is a kind of BceClient, so derived from BceClient
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

func getCreateAPrimaryInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateAProxyInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateReadOnlyInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getQueryTaskListUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TASK
}
