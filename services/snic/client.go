package snic

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "snic." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_ENDPOINT = "endpoint"

	CONSTANT_PUBLIC_SERVICE = "publicService"
)

// Client of snic service is a kind of BceClient, so derived from BceClient
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

func getCreateSnicUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT
}
func getDeleteSnicUri(version string, EndpointId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT + bce.URI_PREFIX + EndpointId
}
func getDescribeSnicUri(version string, EndpointId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT + bce.URI_PREFIX + EndpointId
}
func getListSnicUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT
}
func getQueryAvailablePublicServicesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT + bce.URI_PREFIX + CONSTANT_PUBLIC_SERVICE
}
func getUpdateSnicUri(version string, EndpointId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT + bce.URI_PREFIX + EndpointId
}
func getUpdateSnicEsgUri(version string, EndpointId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT + bce.URI_PREFIX + EndpointId
}
func getUpdateSnicSgUri(version string, EndpointId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENDPOINT + bce.URI_PREFIX + EndpointId
}
