package scs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "scs." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V2 = "v2"

	CONSTANT_INSTANCE = "instance"
)

// Client of scs service is a kind of BceClient, so derived from BceClient
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

func getInstanceListUri(Marker string, MaxKeys string, InstanceIds string, VnetIp string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE
}
