package finance

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const (
	VERSION_V1 = "v1"
)

// CreateRenewResourceRule
//
// PARAMS:
//   - request: the arguments to CreateRenewResourceRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateRenewResourceRule(request *CreateRenewResourceRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRenewResourceRuleUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// GetRenewResourceList
//
// PARAMS:
//   - request: the arguments to GetRenewResourceList
//
// RETURNS:
//   - GetRenewResourceListResponse: The return type of the GetRenewResourceList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRenewResourceList(request *GetRenewResourceListRequest) (*GetRenewResourceListResponse, error) {
	result := &GetRenewResourceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetRenewResourceListUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
