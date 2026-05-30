package pfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// CreatePfs
//
// PARAMS:
//   - request: the arguments to CreatePfs
//
// RETURNS:
//   - CreatePfsResponse: The return type of the CreatePfs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePfs(request *CreatePfsRequest) (*CreatePfsResponse, error) {
	result := &CreatePfsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePfsUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeletePfs
//
// PARAMS:
//   - request: the arguments to DeletePfs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePfs(request *DeletePfsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePfsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		Do()
}

// DescPfs
//
// PARAMS:
//   - request: the arguments to DescPfs
//
// RETURNS:
//   - DescPfsResponse: The return type of the DescPfs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescPfs(request *DescPfsRequest) (*DescPfsResponse, error) {
	result := &DescPfsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescPfsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPfs
//
// PARAMS:
//   - request: the arguments to ListPfs
//
// RETURNS:
//   - ListPfsResponse: The return type of the ListPfs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPfs(request *ListPfsRequest) (*ListPfsResponse, error) {
	result := &ListPfsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPfsUri()).
		WithQueryParamFilter("manner", "marker").
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("manner", util.StringValue(request.Manner)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("filterTag", util.StringValue(request.FilterTag)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdatePFSTag
//
// PARAMS:
//   - request: the arguments to UpdatePFSTag
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePFSTag(request *UpdatePFSTagRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdatePFSTagUri()).
		WithBody(request).
		Do()
}
