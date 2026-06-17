package sts

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AssumeRole
//
// PARAMS:
//   - request: the arguments to AssumeRole
//
// RETURNS:
//   - AssumeRoleResponse: The return type of the AssumeRole interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AssumeRole(request *AssumeRoleRequest) (*AssumeRoleResponse, error) {
	result := &AssumeRoleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAssumeRoleUri(VERSION_V1)).
		WithQueryParam("assumeRole", "").
		WithQueryParamFilter("durationSeconds", util.StringValue(request.DurationSeconds)).
		WithQueryParamFilter("accountId", util.StringValue(request.AccountId)).
		WithQueryParamFilter("userId", util.StringValue(request.UserId)).
		WithQueryParamFilter("roleName", util.StringValue(request.RoleName)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSessionToken
//
// PARAMS:
//   - request: the arguments to GetSessionToken
//
// RETURNS:
//   - GetSessionTokenResponse: The return type of the GetSessionToken interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSessionToken(request *GetSessionTokenRequest) (*GetSessionTokenResponse, error) {
	result := &GetSessionTokenResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetSessionTokenUri(VERSION_V1)).
		WithQueryParamFilter("durationSeconds", util.StringValue(request.DurationSeconds)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
