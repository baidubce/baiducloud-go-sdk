package pfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// CreateFileset
//
// PARAMS:
//   - request: the arguments to CreateFileset
//
// RETURNS:
//   - CreateFilesetResponse: The return type of the CreateFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFileset(request *CreateFilesetRequest) (*CreateFilesetResponse, error) {
	result := &CreateFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateFilesetUri()).
		WithQueryParamFilter("action", "CreateFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

// DeleteFileset
//
// PARAMS:
//   - request: the arguments to DeleteFileset
//
// RETURNS:
//   - DeleteFilesetResponse: The return type of the DeleteFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteFileset(request *DeleteFilesetRequest) (*DeleteFilesetResponse, error) {
	result := &DeleteFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteFilesetUri()).
		WithQueryParamFilter("action", "DeleteFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
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

// DescFileset
//
// PARAMS:
//   - request: the arguments to DescFileset
//
// RETURNS:
//   - DescFilesetResponse: The return type of the DescFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescFileset(request *DescFilesetRequest) (*DescFilesetResponse, error) {
	result := &DescFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescFilesetUri()).
		WithQueryParamFilter("action", "DescribeFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// InstanceListClients
//
// PARAMS:
//   - request: the arguments to InstanceListClients
//
// RETURNS:
//   - InstanceListClientsResponse: The return type of the InstanceListClients interface.
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceListClients(request *InstanceListClientsRequest) (*InstanceListClientsResponse, error) {
	result := &InstanceListClientsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceListClientsUri()).
		WithQueryParamFilter("action", "InstanceListClients").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListFileset
//
// PARAMS:
//   - request: the arguments to ListFileset
//
// RETURNS:
//   - ListFilesetResponse: The return type of the ListFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListFileset(request *ListFilesetRequest) (*ListFilesetResponse, error) {
	result := &ListFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListFilesetUri()).
		WithQueryParamFilter("action", "ListFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
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

// MountTargetListClients
//
// PARAMS:
//   - request: the arguments to MountTargetListClients
//
// RETURNS:
//   - MountTargetListClientsResponse: The return type of the MountTargetListClients interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MountTargetListClients(request *MountTargetListClientsRequest) (*MountTargetListClientsResponse, error) {
	result := &MountTargetListClientsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMountTargetListClientsUri()).
		WithQueryParamFilter("action", "MountTargetListClients").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFileset
//
// PARAMS:
//   - request: the arguments to UpdateFileset
//
// RETURNS:
//   - UpdateFilesetResponse: The return type of the UpdateFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateFileset(request *UpdateFilesetRequest) (*UpdateFilesetResponse, error) {
	result := &UpdateFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateFilesetUri()).
		WithQueryParamFilter("action", "UpdateFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
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
