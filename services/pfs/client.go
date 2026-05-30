package pfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "pfs." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V1 = "v1"

	CONSTANT_PFS = "pfs"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_TAG = "tag"
)

// Client of pfs service is a kind of BceClient, so derived from BceClient
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

func getCreatePfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDeletePfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDescPfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListPfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getUpdatePFSTagUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_TAG
}
