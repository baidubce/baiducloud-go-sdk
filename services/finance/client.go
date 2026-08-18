package finance

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "finance." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_RENEW = "renew"

	CONSTANT_RESOURCE = "resource"

	CONSTANT_LIST = "list"

	CONSTANT_RULE = "rule"

	CONSTANT_CREATE = "create"
)

// Client of finance service is a kind of BceClient, so derived from BceClient
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

func getCreateRenewResourceRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RENEW + bce.URI_PREFIX + CONSTANT_RESOURCE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + CONSTANT_CREATE
}
func getGetRenewResourceListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RENEW + bce.URI_PREFIX + CONSTANT_RESOURCE + bce.URI_PREFIX + CONSTANT_LIST
}
