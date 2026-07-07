package bci

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bci." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_BATCH_DEL = "batchDel"

	CONSTANT_IMAGE_CACHE = "imageCache"
)

// Client of bci service is a kind of BceClient, so derived from BceClient
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

func getBatchDeleteImageCachesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE_CACHE + bce.URI_PREFIX + CONSTANT_BATCH_DEL
}
func getBatchDeleteInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_DEL
}
func getCreateImageCacheUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE_CACHE
}
func getCreateInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDeleteInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getGetInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getListImageCachesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE_CACHE
}
func getListInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
