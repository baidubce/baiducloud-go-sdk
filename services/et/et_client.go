package et

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// ApplyPhysicalDedicatedLine
//
// PARAMS:
//   - request: the arguments to ApplyPhysicalDedicatedLine
//
// RETURNS:
//   - ApplyPhysicalDedicatedLineResponse: The return type of the ApplyPhysicalDedicatedLine interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApplyPhysicalDedicatedLine(request *ApplyPhysicalDedicatedLineRequest) (*ApplyPhysicalDedicatedLineResponse, error) {
	result := &ApplyPhysicalDedicatedLineResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApplyPhysicalDedicatedLineUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AssociatedDedicatedChannel
//
// PARAMS:
//   - request: the arguments to AssociatedDedicatedChannel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AssociatedDedicatedChannel(request *AssociatedDedicatedChannelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAssociatedDedicatedChannelUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("associate", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateDedicatedChannel
//
// PARAMS:
//   - request: the arguments to CreateDedicatedChannel
//
// RETURNS:
//   - CreateDedicatedChannelResponse: The return type of the CreateDedicatedChannel interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedChannel(request *CreateDedicatedChannelRequest) (*CreateDedicatedChannelResponse, error) {
	result := &CreateDedicatedChannelResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDedicatedChannelUri(VERSION_V1, util.StringValue(request.EtId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDedicatedChannelBfd
//
// PARAMS:
//   - request: the arguments to CreateDedicatedChannelBfd
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedChannelBfd(request *CreateDedicatedChannelBfdRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDedicatedChannelBfdUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateDedicatedChannelRouteParameters
//
// PARAMS:
//   - request: the arguments to CreateDedicatedChannelRouteParameters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedChannelRouteParameters(request *CreateDedicatedChannelRouteParametersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCreateDedicatedChannelRouteParametersUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("addRoutes", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateDedicatedChannelRouteRules
//
// PARAMS:
//   - request: the arguments to CreateDedicatedChannelRouteRules
//
// RETURNS:
//   - CreateDedicatedChannelRouteRulesResponse: The return type of the CreateDedicatedChannelRouteRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedChannelRouteRules(request *CreateDedicatedChannelRouteRulesRequest) (*CreateDedicatedChannelRouteRulesResponse, error) {
	result := &CreateDedicatedChannelRouteRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDedicatedChannelRouteRulesUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDedicatedChannelUserObject
//
// PARAMS:
//   - request: the arguments to CreateDedicatedChannelUserObject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateDedicatedChannelUserObject(request *CreateDedicatedChannelUserObjectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCreateDedicatedChannelUserObjectUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("addUsers", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteDedicatedChannel
//
// PARAMS:
//   - request: the arguments to DeleteDedicatedChannel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDedicatedChannel(request *DeleteDedicatedChannelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDedicatedChannelUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteDedicatedChannelBfd
//
// PARAMS:
//   - request: the arguments to DeleteDedicatedChannelBfd
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDedicatedChannelBfd(request *DeleteDedicatedChannelBfdRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDedicatedChannelBfdUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteDedicatedChannelRouteRules
//
// PARAMS:
//   - request: the arguments to DeleteDedicatedChannelRouteRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDedicatedChannelRouteRules(request *DeleteDedicatedChannelRouteRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDedicatedChannelRouteRulesUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId), util.StringValue(request.RouteRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeletePhysicalDedicatedLine
//
// PARAMS:
//   - request: the arguments to DeletePhysicalDedicatedLine
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePhysicalDedicatedLine(request *DeletePhysicalDedicatedLineRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePhysicalDedicatedLineUri(VERSION_V1, util.StringValue(request.DcphyId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DisableDedicatedChannelIpv6
//
// PARAMS:
//   - request: the arguments to DisableDedicatedChannelIpv6
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableDedicatedChannelIpv6(request *DisableDedicatedChannelIpv6Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisableDedicatedChannelIpv6Uri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("disableIpv6", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// EnableDedicatedChannelIpv6
//
// PARAMS:
//   - request: the arguments to EnableDedicatedChannelIpv6
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableDedicatedChannelIpv6(request *EnableDedicatedChannelIpv6Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableDedicatedChannelIpv6Uri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("enableIpv6", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// QueryDedicatedChannel
//
// PARAMS:
//   - request: the arguments to QueryDedicatedChannel
//
// RETURNS:
//   - QueryDedicatedChannelResponse: The return type of the QueryDedicatedChannel interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryDedicatedChannel(request *QueryDedicatedChannelRequest) (*QueryDedicatedChannelResponse, error) {
	result := &QueryDedicatedChannelResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryDedicatedChannelUri(VERSION_V1, util.StringValue(request.EtId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("etChannelId", util.StringValue(request.EtChannelId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryDedicatedChannelRouteRules
//
// PARAMS:
//   - request: the arguments to QueryDedicatedChannelRouteRules
//
// RETURNS:
//   - QueryDedicatedChannelRouteRulesResponse: The return type of the QueryDedicatedChannelRouteRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryDedicatedChannelRouteRules(request *QueryDedicatedChannelRouteRulesRequest) (*QueryDedicatedChannelRouteRulesResponse, error) {
	result := &QueryDedicatedChannelRouteRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryDedicatedChannelRouteRulesUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("destAddress", util.StringValue(request.DestAddress)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryDedicatedLineDetail
//
// PARAMS:
//   - request: the arguments to QueryDedicatedLineDetail
//
// RETURNS:
//   - QueryDedicatedLineDetailResponse: The return type of the QueryDedicatedLineDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryDedicatedLineDetail(request *QueryDedicatedLineDetailRequest) (*QueryDedicatedLineDetailResponse, error) {
	result := &QueryDedicatedLineDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryDedicatedLineDetailUri(VERSION_V1, util.StringValue(request.DcphyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryDedicatedLines
//
// PARAMS:
//   - request: the arguments to QueryDedicatedLines
//
// RETURNS:
//   - QueryDedicatedLinesResponse: The return type of the QueryDedicatedLines interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryDedicatedLines(request *QueryDedicatedLinesRequest) (*QueryDedicatedLinesResponse, error) {
	result := &QueryDedicatedLinesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryDedicatedLinesUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveDedicatedChannelRouteParameters
//
// PARAMS:
//   - request: the arguments to RemoveDedicatedChannelRouteParameters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveDedicatedChannelRouteParameters(request *RemoveDedicatedChannelRouteParametersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRemoveDedicatedChannelRouteParametersUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("removeRoutes", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RemoveDedicatedChannelUserObject
//
// PARAMS:
//   - request: the arguments to RemoveDedicatedChannelUserObject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveDedicatedChannelUserObject(request *RemoveDedicatedChannelUserObjectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRemoveDedicatedChannelUserObjectUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("removeUsers", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ResubmitDedicatedChannel
//
// PARAMS:
//   - request: the arguments to ResubmitDedicatedChannel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResubmitDedicatedChannel(request *ResubmitDedicatedChannelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResubmitDedicatedChannelUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("reCreate", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UnrelatedDedicatedLineChannel
//
// PARAMS:
//   - request: the arguments to UnrelatedDedicatedLineChannel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnrelatedDedicatedLineChannel(request *UnrelatedDedicatedLineChannelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnrelatedDedicatedLineChannelUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("disassociate", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateDedicatedChannel
//
// PARAMS:
//   - request: the arguments to UpdateDedicatedChannel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDedicatedChannel(request *UpdateDedicatedChannelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDedicatedChannelUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateDedicatedChannelBfd
//
// PARAMS:
//   - request: the arguments to UpdateDedicatedChannelBfd
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDedicatedChannelBfd(request *UpdateDedicatedChannelBfdRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDedicatedChannelBfdUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateDedicatedChannelRouteRules
//
// PARAMS:
//   - request: the arguments to UpdateDedicatedChannelRouteRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDedicatedChannelRouteRules(request *UpdateDedicatedChannelRouteRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDedicatedChannelRouteRulesUri(VERSION_V1, util.StringValue(request.EtId), util.StringValue(request.EtChannelId), util.StringValue(request.RouteRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdatePhysicalDedicatedLine
//
// PARAMS:
//   - request: the arguments to UpdatePhysicalDedicatedLine
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePhysicalDedicatedLine(request *UpdatePhysicalDedicatedLineRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePhysicalDedicatedLineUri(VERSION_V1, util.StringValue(request.DcphyId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
