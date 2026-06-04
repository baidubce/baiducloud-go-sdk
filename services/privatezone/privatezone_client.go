package privatezone

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddRecord
//
// PARAMS:
//   - request: the arguments to AddRecord
//
// RETURNS:
//   - AddRecordResponse: The return type of the AddRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddRecord(request *AddRecordRequest) (*AddRecordResponse, error) {
	result := &AddRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddRecordUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindVpc
//
// PARAMS:
//   - request: the arguments to BindVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindVpc(request *BindVpcRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindVpcUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindVpcToRule
//
// PARAMS:
//   - request: the arguments to BindVpcToRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindVpcToRule(request *BindVpcToRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindVpcToRuleUri(VERSION_V1, util.StringValue(request.RuleId))).
		WithQueryParamFilter("clienToken", util.StringValue(request.ClienToken)).
		WithBody(request).
		Do()
}

// CreatePrivateZone
//
// PARAMS:
//   - request: the arguments to CreatePrivateZone
//
// RETURNS:
//   - CreatePrivateZoneResponse: The return type of the CreatePrivateZone interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePrivateZone(request *CreatePrivateZoneRequest) (*CreatePrivateZoneResponse, error) {
	result := &CreatePrivateZoneResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePrivateZoneUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateResolver
//
// PARAMS:
//   - request: the arguments to CreateResolver
//
// RETURNS:
//   - CreateResolverResponse: The return type of the CreateResolver interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateResolver(request *CreateResolverRequest) (*CreateResolverResponse, error) {
	result := &CreateResolverResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateResolverUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateResolverRule
//
// PARAMS:
//   - request: the arguments to CreateResolverRule
//
// RETURNS:
//   - CreateResolverRuleResponse: The return type of the CreateResolverRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateResolverRule(request *CreateResolverRuleRequest) (*CreateResolverRuleResponse, error) {
	result := &CreateResolverRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateResolverRuleUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeletePrivateZone
//
// PARAMS:
//   - request: the arguments to DeletePrivateZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePrivateZone(request *DeletePrivateZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePrivateZoneUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteRecord
//
// PARAMS:
//   - request: the arguments to DeleteRecord
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRecord(request *DeleteRecordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRecordUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteResolver
//
// PARAMS:
//   - request: the arguments to DeleteResolver
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteResolver(request *DeleteResolverRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteResolverUri(VERSION_V1, util.StringValue(request.ResolverId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteResolverRule
//
// PARAMS:
//   - request: the arguments to DeleteResolverRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteResolverRule(request *DeleteResolverRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteResolverRuleUri(VERSION_V1, util.StringValue(request.RuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DisableRecord
//
// PARAMS:
//   - request: the arguments to DisableRecord
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableRecord(request *DisableRecordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisableRecordUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParam("disable", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// EnableRecord
//
// PARAMS:
//   - request: the arguments to EnableRecord
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableRecord(request *EnableRecordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableRecordUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParam("enable", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// GetDnsResolverDetail
//
// PARAMS:
//   - request: the arguments to GetDnsResolverDetail
//
// RETURNS:
//   - GetDnsResolverDetailResponse: The return type of the GetDnsResolverDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDnsResolverDetail(request *GetDnsResolverDetailRequest) (*GetDnsResolverDetailResponse, error) {
	result := &GetDnsResolverDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDnsResolverDetailUri(VERSION_V1, util.StringValue(request.ResolverId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDnsResolverList
//
// PARAMS:
//   - request: the arguments to GetDnsResolverList
//
// RETURNS:
//   - GetDnsResolverListResponse: The return type of the GetDnsResolverList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDnsResolverList(request *GetDnsResolverListRequest) (*GetDnsResolverListResponse, error) {
	result := &GetDnsResolverListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDnsResolverListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.StringValue(request.MaxKeys)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDnsResolverRuleDetail
//
// PARAMS:
//   - request: the arguments to GetDnsResolverRuleDetail
//
// RETURNS:
//   - GetDnsResolverRuleDetailResponse: The return type of the GetDnsResolverRuleDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDnsResolverRuleDetail(request *GetDnsResolverRuleDetailRequest) (*GetDnsResolverRuleDetailResponse, error) {
	result := &GetDnsResolverRuleDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDnsResolverRuleDetailUri(VERSION_V1, util.StringValue(request.RuleId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDnsResolverRuleList
//
// PARAMS:
//   - request: the arguments to GetDnsResolverRuleList
//
// RETURNS:
//   - GetDnsResolverRuleListResponse: The return type of the GetDnsResolverRuleList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDnsResolverRuleList(request *GetDnsResolverRuleListRequest) (*GetDnsResolverRuleListResponse, error) {
	result := &GetDnsResolverRuleListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDnsResolverRuleListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.StringValue(request.MaxKeys)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPrivateZone
//
// PARAMS:
//   - request: the arguments to GetPrivateZone
//
// RETURNS:
//   - GetPrivateZoneResponse: The return type of the GetPrivateZone interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPrivateZone(request *GetPrivateZoneRequest) (*GetPrivateZoneResponse, error) {
	result := &GetPrivateZoneResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPrivateZoneUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPrivateZone
//
// PARAMS:
//   - request: the arguments to ListPrivateZone
//
// RETURNS:
//   - ListPrivateZoneResponse: The return type of the ListPrivateZone interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPrivateZone(request *ListPrivateZoneRequest) (*ListPrivateZoneResponse, error) {
	result := &ListPrivateZoneResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPrivateZoneUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRecord
//
// PARAMS:
//   - request: the arguments to ListRecord
//
// RETURNS:
//   - ListRecordResponse: The return type of the ListRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRecord(request *ListRecordRequest) (*ListRecordResponse, error) {
	result := &ListRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRecordUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("rr", util.StringValue(request.Rr)).
		WithQueryParamFilter("searchMode", util.StringValue(request.SearchMode)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithQueryParamFilter("value", util.StringValue(request.Value)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindVpc
//
// PARAMS:
//   - request: the arguments to UnbindVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindVpc(request *UnbindVpcRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindVpcUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UnbindVpcToRule
//
// PARAMS:
//   - request: the arguments to UnbindVpcToRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindVpcToRule(request *UnbindVpcToRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindVpcToRuleUri(VERSION_V1, util.StringValue(request.RuleId))).
		WithQueryParamFilter("clienToken", util.StringValue(request.ClienToken)).
		WithBody(request).
		Do()
}

// UpdateDnsParser
//
// PARAMS:
//   - request: the arguments to UpdateDnsParser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDnsParser(request *UpdateDnsParserRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDnsParserUri(VERSION_V1, util.StringValue(request.ResolverId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateRecord
//
// PARAMS:
//   - request: the arguments to UpdateRecord
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRecord(request *UpdateRecordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRecordUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateResolverRule
//
// PARAMS:
//   - request: the arguments to UpdateResolverRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateResolverRule(request *UpdateResolverRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateResolverRuleUri(VERSION_V1, util.StringValue(request.ReluId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
