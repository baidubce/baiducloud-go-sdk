package cfw

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// BindCfw
//
// PARAMS:
//   - request: the arguments to BindCfw
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindCfw(request *BindCfwRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindCfwUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithQueryParam("bind", "").
		WithBody(request).
		Do()
}

// CreateCfw
//
// PARAMS:
//   - request: the arguments to CreateCfw
//
// RETURNS:
//   - CreateCfwResponse: The return type of the CreateCfw interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCfw(request *CreateCfwRequest) (*CreateCfwResponse, error) {
	result := &CreateCfwResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCfwUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCfwRule
//
// PARAMS:
//   - request: the arguments to CreateCfwRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateCfwRule(request *CreateCfwRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCfwRuleUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithBody(request).
		Do()
}

// DeleteCfw
//
// PARAMS:
//   - request: the arguments to DeleteCfw
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCfw(request *DeleteCfwRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteCfwUri(VERSION_V1, util.StringValue(request.CfwId))).
		Do()
}

// DeleteCfwRule
//
// PARAMS:
//   - request: the arguments to DeleteCfwRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCfwRule(request *DeleteCfwRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDeleteCfwRuleUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithBody(request).
		Do()
}

// DisableCfwProtect
//
// PARAMS:
//   - request: the arguments to DisableCfwProtect
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableCfwProtect(request *DisableCfwProtectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisableCfwProtectUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithQueryParam("off", "").
		WithBody(request).
		Do()
}

// EnableCfwProtect
//
// PARAMS:
//   - request: the arguments to EnableCfwProtect
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableCfwProtect(request *EnableCfwProtectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableCfwProtectUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithQueryParam("on", "").
		WithBody(request).
		Do()
}

// GetCfw
//
// PARAMS:
//   - request: the arguments to GetCfw
//
// RETURNS:
//   - GetCfwResponse: The return type of the GetCfw interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetCfw(request *GetCfwRequest) (*GetCfwResponse, error) {
	result := &GetCfwResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetCfwUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCfw
//
// PARAMS:
//   - request: the arguments to ListCfw
//
// RETURNS:
//   - ListCfwResponse: The return type of the ListCfw interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListCfw(request *ListCfwRequest) (*ListCfwResponse, error) {
	result := &ListCfwResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListCfwUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListProtectInstances
//
// PARAMS:
//   - request: the arguments to ListProtectInstances
//
// RETURNS:
//   - ListProtectInstancesResponse: The return type of the ListProtectInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListProtectInstances(request *ListProtectInstancesRequest) (*ListProtectInstancesResponse, error) {
	result := &ListProtectInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListProtectInstancesUri(VERSION_V1)).
		WithQueryParamFilter("instanceType", util.StringValue(request.InstanceType)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("region", util.StringValue(request.Region)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindCfw
//
// PARAMS:
//   - request: the arguments to UnbindCfw
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindCfw(request *UnbindCfwRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindCfwUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithQueryParam("unbind", "").
		WithBody(request).
		Do()
}

// UpdateCfw
//
// PARAMS:
//   - request: the arguments to UpdateCfw
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateCfw(request *UpdateCfwRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateCfwUri(VERSION_V1, util.StringValue(request.CfwId))).
		WithBody(request).
		Do()
}

// UpdateCfwRule
//
// PARAMS:
//   - request: the arguments to UpdateCfwRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateCfwRule(request *UpdateCfwRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateCfwRuleUri(VERSION_V1, util.StringValue(request.CfwId), util.StringValue(request.CfwRuleId))).
		WithBody(request).
		Do()
}
