package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// CreateAPrimaryInstance
//
// PARAMS:
//   - request: the arguments to CreateAPrimaryInstance
//
// RETURNS:
//   - CreateAPrimaryInstanceResponse: The return type of the CreateAPrimaryInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAPrimaryInstance(request *CreateAPrimaryInstanceRequest) (*CreateAPrimaryInstanceResponse, error) {
	result := &CreateAPrimaryInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAPrimaryInstanceUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAProxyInstance
//
// PARAMS:
//   - request: the arguments to CreateAProxyInstance
//
// RETURNS:
//   - CreateAProxyInstanceResponse: The return type of the CreateAProxyInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAProxyInstance(request *CreateAProxyInstanceRequest) (*CreateAProxyInstanceResponse, error) {
	result := &CreateAProxyInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAProxyInstanceUri(VERSION_V1)).
		WithQueryParam("rdsproxy", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateReadOnlyInstance
//
// PARAMS:
//   - request: the arguments to CreateReadOnlyInstance
//
// RETURNS:
//   - CreateReadOnlyInstanceResponse: The return type of the CreateReadOnlyInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateReadOnlyInstance(request *CreateReadOnlyInstanceRequest) (*CreateReadOnlyInstanceResponse, error) {
	result := &CreateReadOnlyInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateReadOnlyInstanceUri(VERSION_V1)).
		WithQueryParam("readReplica", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTaskList
//
// PARAMS:
//   - request: the arguments to QueryTaskList
//
// RETURNS:
//   - QueryTaskListResponse: The return type of the QueryTaskList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTaskList(request *QueryTaskListRequest) (*QueryTaskListResponse, error) {
	result := &QueryTaskListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryTaskListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
