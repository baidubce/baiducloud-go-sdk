package sts

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "sts." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_SESSION_TOKEN = "sessionToken"

	CONSTANT_CREDENTIAL = "credential"
)

// Client of sts service is a kind of BceClient, so derived from BceClient
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

func getAssumeRoleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CREDENTIAL
}
func getGetSessionTokenUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SESSION_TOKEN
}
