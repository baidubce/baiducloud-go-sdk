package csn

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddRouteRule
//
// PARAMS:
//   - request: the arguments to AddRouteRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddRouteRule(request *AddRouteRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddRouteRuleUri(VERSION_V1, util.StringValue(request.CsnRtId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AttachCsnInstance
//
// PARAMS:
//   - request: the arguments to AttachCsnInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AttachCsnInstance(request *AttachCsnInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAttachCsnInstanceUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParam("attach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindCsnBp
//
// PARAMS:
//   - request: the arguments to BindCsnBp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindCsnBp(request *BindCsnBpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindCsnBpUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateAssociationRelation
//
// PARAMS:
//   - request: the arguments to CreateAssociationRelation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAssociationRelation(request *CreateAssociationRelationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAssociationRelationUri(VERSION_V1, util.StringValue(request.CsnRtId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateCsn
//
// PARAMS:
//   - request: the arguments to CreateCsn
//
// RETURNS:
//   - CreateCsnResponse: The return type of the CreateCsn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCsn(request *CreateCsnRequest) (*CreateCsnResponse, error) {
	result := &CreateCsnResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCsnUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCsnBp
//
// PARAMS:
//   - request: the arguments to CreateCsnBp
//
// RETURNS:
//   - CreateCsnBpResponse: The return type of the CreateCsnBp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCsnBp(request *CreateCsnBpRequest) (*CreateCsnBpResponse, error) {
	result := &CreateCsnBpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCsnBpUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRegionBandwidth
//
// PARAMS:
//   - request: the arguments to CreateRegionBandwidth
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateRegionBandwidth(request *CreateRegionBandwidthRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRegionBandwidthUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateStudyRelation
//
// PARAMS:
//   - request: the arguments to CreateStudyRelation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateStudyRelation(request *CreateStudyRelationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateStudyRelationUri(VERSION_V1, util.StringValue(request.CsnRtId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteAssociationRelation
//
// PARAMS:
//   - request: the arguments to DeleteAssociationRelation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAssociationRelation(request *DeleteAssociationRelationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAssociationRelationUri(VERSION_V1, util.StringValue(request.CsnRtId), util.StringValue(request.AttachId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteCsn
//
// PARAMS:
//   - request: the arguments to DeleteCsn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCsn(request *DeleteCsnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteCsnUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteCsnBp
//
// PARAMS:
//   - request: the arguments to DeleteCsnBp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCsnBp(request *DeleteCsnBpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteCsnBpUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteRegionBandwidth
//
// PARAMS:
//   - request: the arguments to DeleteRegionBandwidth
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRegionBandwidth(request *DeleteRegionBandwidthRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteRegionBandwidthUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteRouteRule
//
// PARAMS:
//   - request: the arguments to DeleteRouteRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRouteRule(request *DeleteRouteRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRouteRuleUri(VERSION_V1, util.StringValue(request.CsnRtId), util.StringValue(request.CsnRtRuleId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteStudyRelation
//
// PARAMS:
//   - request: the arguments to DeleteStudyRelation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteStudyRelation(request *DeleteStudyRelationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteStudyRelationUri(VERSION_V1, util.StringValue(request.CsnRtId), util.StringValue(request.AttachId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DetachCsnInstance
//
// PARAMS:
//   - request: the arguments to DetachCsnInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DetachCsnInstance(request *DetachCsnInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDetachCsnInstanceUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParam("detach", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// QueryAssociationRelation
//
// PARAMS:
//   - request: the arguments to QueryAssociationRelation
//
// RETURNS:
//   - QueryAssociationRelationResponse: The return type of the QueryAssociationRelation interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAssociationRelation(request *QueryAssociationRelationRequest) (*QueryAssociationRelationResponse, error) {
	result := &QueryAssociationRelationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAssociationRelationUri(VERSION_V1, util.StringValue(request.CsnRtId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryCsnBpDetail
//
// PARAMS:
//   - request: the arguments to QueryCsnBpDetail
//
// RETURNS:
//   - QueryCsnBpDetailResponse: The return type of the QueryCsnBpDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCsnBpDetail(request *QueryCsnBpDetailRequest) (*QueryCsnBpDetailResponse, error) {
	result := &QueryCsnBpDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryCsnBpDetailUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryCsnBpList
//
// PARAMS:
//   - request: the arguments to QueryCsnBpList
//
// RETURNS:
//   - QueryCsnBpListResponse: The return type of the QueryCsnBpList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCsnBpList(request *QueryCsnBpListRequest) (*QueryCsnBpListResponse, error) {
	result := &QueryCsnBpListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryCsnBpListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryCsnBpPrice
//
// PARAMS:
//   - request: the arguments to QueryCsnBpPrice
//
// RETURNS:
//   - QueryCsnBpPriceResponse: The return type of the QueryCsnBpPrice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCsnBpPrice(request *QueryCsnBpPriceRequest) (*QueryCsnBpPriceResponse, error) {
	result := &QueryCsnBpPriceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryCsnBpPriceUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryCsnDetail
//
// PARAMS:
//   - request: the arguments to QueryCsnDetail
//
// RETURNS:
//   - QueryCsnDetailResponse: The return type of the QueryCsnDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCsnDetail(request *QueryCsnDetailRequest) (*QueryCsnDetailResponse, error) {
	result := &QueryCsnDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryCsnDetailUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryCsnInstance
//
// PARAMS:
//   - request: the arguments to QueryCsnInstance
//
// RETURNS:
//   - QueryCsnInstanceResponse: The return type of the QueryCsnInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCsnInstance(request *QueryCsnInstanceRequest) (*QueryCsnInstanceResponse, error) {
	result := &QueryCsnInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryCsnInstanceUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryCsnList
//
// PARAMS:
//   - request: the arguments to QueryCsnList
//
// RETURNS:
//   - QueryCsnListResponse: The return type of the QueryCsnList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCsnList(request *QueryCsnListRequest) (*QueryCsnListResponse, error) {
	result := &QueryCsnListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryCsnListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRegionBandwidth
//
// PARAMS:
//   - request: the arguments to QueryRegionBandwidth
//
// RETURNS:
//   - QueryRegionBandwidthResponse: The return type of the QueryRegionBandwidth interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRegionBandwidth(request *QueryRegionBandwidthRequest) (*QueryRegionBandwidthResponse, error) {
	result := &QueryRegionBandwidthResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRegionBandwidthUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRegionBandwidthByCsn
//
// PARAMS:
//   - request: the arguments to QueryRegionBandwidthByCsn
//
// RETURNS:
//   - QueryRegionBandwidthByCsnResponse: The return type of the QueryRegionBandwidthByCsn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRegionBandwidthByCsn(request *QueryRegionBandwidthByCsnRequest) (*QueryRegionBandwidthByCsnResponse, error) {
	result := &QueryRegionBandwidthByCsnResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRegionBandwidthByCsnUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRouteRule
//
// PARAMS:
//   - request: the arguments to QueryRouteRule
//
// RETURNS:
//   - QueryRouteRuleResponse: The return type of the QueryRouteRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRouteRule(request *QueryRouteRuleRequest) (*QueryRouteRuleResponse, error) {
	result := &QueryRouteRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRouteRuleUri(VERSION_V1, util.StringValue(request.CsnRtId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryRouteTableList
//
// PARAMS:
//   - request: the arguments to QueryRouteTableList
//
// RETURNS:
//   - QueryRouteTableListResponse: The return type of the QueryRouteTableList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryRouteTableList(request *QueryRouteTableListRequest) (*QueryRouteTableListResponse, error) {
	result := &QueryRouteTableListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRouteTableListUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryStudyRelation
//
// PARAMS:
//   - request: the arguments to QueryStudyRelation
//
// RETURNS:
//   - QueryStudyRelationResponse: The return type of the QueryStudyRelation interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryStudyRelation(request *QueryStudyRelationRequest) (*QueryStudyRelationResponse, error) {
	result := &QueryStudyRelationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryStudyRelationUri(VERSION_V1, util.StringValue(request.CsnRtId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTgwList
//
// PARAMS:
//   - request: the arguments to QueryTgwList
//
// RETURNS:
//   - QueryTgwListResponse: The return type of the QueryTgwList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTgwList(request *QueryTgwListRequest) (*QueryTgwListResponse, error) {
	result := &QueryTgwListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTgwListUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTgwRoute
//
// PARAMS:
//   - request: the arguments to QueryTgwRoute
//
// RETURNS:
//   - QueryTgwRouteResponse: The return type of the QueryTgwRoute interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTgwRoute(request *QueryTgwRouteRequest) (*QueryTgwRouteResponse, error) {
	result := &QueryTgwRouteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTgwRouteUri(VERSION_V1, util.StringValue(request.CsnId), util.StringValue(request.TgwId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResizeCsnBp
//
// PARAMS:
//   - request: the arguments to ResizeCsnBp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeCsnBp(request *ResizeCsnBpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeCsnBpUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UnbindCsnBp
//
// PARAMS:
//   - request: the arguments to UnbindCsnBp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindCsnBp(request *UnbindCsnBpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindCsnBpUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateCsn
//
// PARAMS:
//   - request: the arguments to UpdateCsn
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateCsn(request *UpdateCsnRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateCsnUri(VERSION_V1, util.StringValue(request.CsnId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateCsnBp
//
// PARAMS:
//   - request: the arguments to UpdateCsnBp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateCsnBp(request *UpdateCsnBpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateCsnBpUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateRegionBandwidth
//
// PARAMS:
//   - request: the arguments to UpdateRegionBandwidth
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRegionBandwidth(request *UpdateRegionBandwidthRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRegionBandwidthUri(VERSION_V1, util.StringValue(request.CsnBpId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateTgw
//
// PARAMS:
//   - request: the arguments to UpdateTgw
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTgw(request *UpdateTgwRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateTgwUri(VERSION_V1, util.StringValue(request.CsnId), util.StringValue(request.TgwId))).
		WithBody(request).
		Do()
}
