package aihc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "aihc." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of aihc service is a kind of BceClient, so derived from BceClient
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

func getCreateADatasetV2Uri() string {
	return bce.URI_PREFIX
}
func getCreateAModelV2Uri() string {
	return bce.URI_PREFIX
}
func getCreateDatasetVersionV2Uri() string {
	return bce.URI_PREFIX
}
func getDeleteDatasetV2Uri() string {
	return bce.URI_PREFIX
}
func getDeleteDatasetVersionV2Uri() string {
	return bce.URI_PREFIX
}
func getDeleteModelV2Uri() string {
	return bce.URI_PREFIX
}
func getDeleteModelVersionV2Uri() string {
	return bce.URI_PREFIX
}
func getGetAListOfModelVersionsV2Uri() string {
	return bce.URI_PREFIX
}
func getGetDatasetDetailsV2Uri() string {
	return bce.URI_PREFIX
}
func getGetDatasetVersionDetailsV2Uri() string {
	return bce.URI_PREFIX
}
func getGetModelDetailsV2Uri() string {
	return bce.URI_PREFIX
}
func getGetModelListV2Uri() string {
	return bce.URI_PREFIX
}
func getGetModelVersionDetailsV2Uri() string {
	return bce.URI_PREFIX
}
func getModifyDatasetV2Uri() string {
	return bce.URI_PREFIX
}
func getModifyTheModelV2Uri() string {
	return bce.URI_PREFIX
}
func getNewModelVersionV2Uri() string {
	return bce.URI_PREFIX
}
func getRetrieveTheDatasetListV2Uri() string {
	return bce.URI_PREFIX
}
func getRetrieveTheDatasetVersionListV2Uri() string {
	return bce.URI_PREFIX
}
