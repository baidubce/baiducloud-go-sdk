package cfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// BatchCreationOfPermissionGroupRules
//
// PARAMS:
//   - request: the arguments to BatchCreationOfPermissionGroupRules
//
// RETURNS:
//   - BatchCreationOfPermissionGroupRulesResponse: The return type of the BatchCreationOfPermissionGroupRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchCreationOfPermissionGroupRules(request *BatchCreationOfPermissionGroupRulesRequest) (*BatchCreationOfPermissionGroupRulesResponse, error) {
	result := &BatchCreationOfPermissionGroupRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchCreationOfPermissionGroupRulesUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFileSystem
//
// PARAMS:
//   - request: the arguments to CreateFileSystem
//
// RETURNS:
//   - CreateFileSystemResponse: The return type of the CreateFileSystem interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) (*CreateFileSystemResponse, error) {
	result := &CreateFileSystemResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateFileSystemUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateMountingTarget
//
// PARAMS:
//   - request: the arguments to CreateMountingTarget
//
// RETURNS:
//   - CreateMountingTargetResponse: The return type of the CreateMountingTarget interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateMountingTarget(request *CreateMountingTargetRequest) (*CreateMountingTargetResponse, error) {
	result := &CreateMountingTargetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateMountingTargetUri(VERSION_V1, util.StringValue(request.FsId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePermissionGroup
//
// PARAMS:
//   - request: the arguments to CreatePermissionGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreatePermissionGroup(request *CreatePermissionGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePermissionGroupUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreatePermissionGroupRules
//
// PARAMS:
//   - request: the arguments to CreatePermissionGroupRules
//
// RETURNS:
//   - CreatePermissionGroupRulesResponse: The return type of the CreatePermissionGroupRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePermissionGroupRules(request *CreatePermissionGroupRulesRequest) (*CreatePermissionGroupRulesResponse, error) {
	result := &CreatePermissionGroupRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePermissionGroupRulesUri(VERSION_V1, util.StringValue(request.AgName))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeletePermissionGroup
//
// PARAMS:
//   - request: the arguments to DeletePermissionGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePermissionGroup(request *DeletePermissionGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePermissionGroupUri(VERSION_V1, util.StringValue(request.AgName))).
		Do()
}

// DeletePermissionGroupRule
//
// PARAMS:
//   - request: the arguments to DeletePermissionGroupRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePermissionGroupRule(request *DeletePermissionGroupRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePermissionGroupRuleUri(VERSION_V1, util.StringValue(request.AgName), util.StringValue(request.ArId))).
		Do()
}

// DropFileSystem
//
// PARAMS:
//   - request: the arguments to DropFileSystem
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DropFileSystem(request *DropFileSystemRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDropFileSystemUri(VERSION_V1, util.StringValue(request.FsId))).
		Do()
}

// DropMountTarget
//
// PARAMS:
//   - request: the arguments to DropMountTarget
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DropMountTarget(request *DropMountTargetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDropMountTargetUri(VERSION_V1, util.StringValue(request.FsId), util.StringValue(request.MountId))).
		Do()
}

// ModifyTheMountTargetPermissionGroup
//
// PARAMS:
//   - request: the arguments to ModifyTheMountTargetPermissionGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTheMountTargetPermissionGroup(request *ModifyTheMountTargetPermissionGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTheMountTargetPermissionGroupUri(VERSION_V1, util.StringValue(request.FsId), util.StringValue(request.MountID))).
		WithBody(request).
		Do()
}

// QueryFileSystem
//
// PARAMS:
//   - request: the arguments to QueryFileSystem
//
// RETURNS:
//   - QueryFileSystemResponse: The return type of the QueryFileSystem interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryFileSystem(request *QueryFileSystemRequest) (*QueryFileSystemResponse, error) {
	result := &QueryFileSystemResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryFileSystemUri(VERSION_V1)).
		WithQueryParamFilter("userId", util.StringValue(request.UserId)).
		WithQueryParamFilter("fsId", util.StringValue(request.FsId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("filterTag", util.StringValue(request.FilterTag)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryMountedClient
//
// PARAMS:
//   - request: the arguments to QueryMountedClient
//
// RETURNS:
//   - QueryMountedClientResponse: The return type of the QueryMountedClient interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryMountedClient(request *QueryMountedClientRequest) (*QueryMountedClientResponse, error) {
	result := &QueryMountedClientResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryMountedClientUri(VERSION_V1)).
		WithQueryParamFilter("fsId", util.StringValue(request.FsId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryMountingTarget
//
// PARAMS:
//   - request: the arguments to QueryMountingTarget
//
// RETURNS:
//   - QueryMountingTargetResponse: The return type of the QueryMountingTarget interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryMountingTarget(request *QueryMountingTargetRequest) (*QueryMountingTargetResponse, error) {
	result := &QueryMountingTargetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryMountingTargetUri(VERSION_V1, util.StringValue(request.FId))).
		WithQueryParam("", "").
		WithQueryParamFilter("mountId", util.StringValue(request.MountId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryPermissionGroup
//
// PARAMS:
//   - request: the arguments to QueryPermissionGroup
//
// RETURNS:
//   - QueryPermissionGroupResponse: The return type of the QueryPermissionGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryPermissionGroup(request *QueryPermissionGroupRequest) (*QueryPermissionGroupResponse, error) {
	result := &QueryPermissionGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryPermissionGroupUri(VERSION_V1)).
		WithQueryParamFilter("agName", util.StringValue(request.AgName)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryPermissionGroupRules
//
// PARAMS:
//   - request: the arguments to QueryPermissionGroupRules
//
// RETURNS:
//   - QueryPermissionGroupRulesResponse: The return type of the QueryPermissionGroupRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryPermissionGroupRules(request *QueryPermissionGroupRulesRequest) (*QueryPermissionGroupRulesResponse, error) {
	result := &QueryPermissionGroupRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryPermissionGroupRulesUri(VERSION_V1, util.StringValue(request.AgName))).
		WithQueryParamFilter("arId", util.StringValue(request.ArId)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFileSystem
//
// PARAMS:
//   - request: the arguments to UpdateFileSystem
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateFileSystem(request *UpdateFileSystemRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateFileSystemUri(VERSION_V1, util.StringValue(request.FsId))).
		WithBody(request).
		Do()
}

// UpdateFileSystemLabels
//
// PARAMS:
//   - request: the arguments to UpdateFileSystemLabels
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateFileSystemLabels(request *UpdateFileSystemLabelsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateFileSystemLabelsUri(VERSION_V1)).
		WithQueryParam("tag", "").
		WithQueryParamFilter("tag", util.StringValue(request.Tag)).
		WithBody(request).
		Do()
}

// UpdatePermissionGroup
//
// PARAMS:
//   - request: the arguments to UpdatePermissionGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePermissionGroup(request *UpdatePermissionGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePermissionGroupUri(VERSION_V1, util.StringValue(request.AgName))).
		WithBody(request).
		Do()
}

// UpdatePermissionGroupRules
//
// PARAMS:
//   - request: the arguments to UpdatePermissionGroupRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePermissionGroupRules(request *UpdatePermissionGroupRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePermissionGroupRulesUri(VERSION_V1, util.StringValue(request.AgName), util.StringValue(request.ArId))).
		WithBody(request).
		Do()
}
