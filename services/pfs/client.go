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

func getCancelL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getCreateFilesetUri() string {
	return bce.URI_PREFIX
}
func getCreateL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getCreateL2PolicyUri() string {
	return bce.URI_PREFIX
}
func getCreatePfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDeleteFilesetUri() string {
	return bce.URI_PREFIX
}
func getDeleteL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getDeleteL2PolicyUri() string {
	return bce.URI_PREFIX
}
func getDeletePfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDescFilesetUri() string {
	return bce.URI_PREFIX
}
func getDescL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getDescL2PolicyUri() string {
	return bce.URI_PREFIX
}
func getDescPfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getInstanceListClientsUri() string {
	return bce.URI_PREFIX
}
func getListFilesetUri() string {
	return bce.URI_PREFIX
}
func getListL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getListL2PolicyUri() string {
	return bce.URI_PREFIX
}
func getListPfsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getLstPerL2BktLnkExecLogUri() string {
	return bce.URI_PREFIX
}
func getMountTargetListClientsUri() string {
	return bce.URI_PREFIX
}
func getPauseL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getQryL2PolExecDetailUri() string {
	return bce.URI_PREFIX
}
func getQryL2PolExecLogUri() string {
	return bce.URI_PREFIX
}
func getResumeL2BucketLinkUri() string {
	return bce.URI_PREFIX
}
func getUpdPerL2BktLnkInfoUri() string {
	return bce.URI_PREFIX
}
func getUpdateFilesetUri() string {
	return bce.URI_PREFIX
}
func getUpdateL2PolicyUri() string {
	return bce.URI_PREFIX
}
func getUpdatePFSTagUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PFS + bce.URI_PREFIX + CONSTANT_TAG
}
