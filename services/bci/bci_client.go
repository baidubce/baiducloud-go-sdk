package bci

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// BatchDeleteImageCaches
//
// PARAMS:
//   - request: the arguments to BatchDeleteImageCaches
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteImageCaches(request *BatchDeleteImageCachesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchDeleteImageCachesUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// BatchDeleteInstances
//
// PARAMS:
//   - request: the arguments to BatchDeleteInstances
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteInstances(request *BatchDeleteInstancesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchDeleteInstancesUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreateImageCache
//
// PARAMS:
//   - request: the arguments to CreateImageCache
//
// RETURNS:
//   - CreateImageCacheResponse: The return type of the CreateImageCache interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateImageCache(request *CreateImageCacheRequest) (*CreateImageCacheResponse, error) {
	result := &CreateImageCacheResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateImageCacheUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateInstance
//
// PARAMS:
//   - request: the arguments to CreateInstance
//
// RETURNS:
//   - CreateInstanceResponse: The return type of the CreateInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(request *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	result := &CreateInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteInstance
//
// PARAMS:
//   - request: the arguments to DeleteInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteInstanceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("relatedReleaseFlag", util.BoolValue(request.RelatedReleaseFlag)).
		Do()
}

// GetInstance
//
// PARAMS:
//   - request: the arguments to GetInstance
//
// RETURNS:
//   - GetInstanceResponse: The return type of the GetInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstance(request *GetInstanceRequest) (*GetInstanceResponse, error) {
	result := &GetInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListImageCaches
//
// PARAMS:
//   - request: the arguments to ListImageCaches
//
// RETURNS:
//   - ListImageCachesResponse: The return type of the ListImageCaches interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListImageCaches(request *ListImageCachesRequest) (*ListImageCachesResponse, error) {
	result := &ListImageCachesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListImageCachesUri(VERSION_V1)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstances
//
// PARAMS:
//   - request: the arguments to ListInstances
//
// RETURNS:
//   - ListInstancesResponse: The return type of the ListInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(request *ListInstancesRequest) (*ListInstancesResponse, error) {
	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstancesUri(VERSION_V1)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
