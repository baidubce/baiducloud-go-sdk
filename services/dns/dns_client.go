package dns

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddLineGroup
//
// PARAMS:
//   - request: the arguments to AddLineGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddLineGroup(request *AddLineGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddLineGroupUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreatePaidZone
//
// PARAMS:
//   - request: the arguments to CreatePaidZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreatePaidZone(request *CreatePaidZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePaidZoneUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateRecord
//
// PARAMS:
//   - request: the arguments to CreateRecord
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateRecord(request *CreateRecordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRecordUri(VERSION_V1, util.StringValue(request.ZoneName))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateZone
//
// PARAMS:
//   - request: the arguments to CreateZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateZone(request *CreateZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateZoneUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// DeleteLineGroup
//
// PARAMS:
//   - request: the arguments to DeleteLineGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteLineGroup(request *DeleteLineGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteLineGroupUri(VERSION_V1, util.StringValue(request.LineId))).
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
		WithURL(getDeleteRecordUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DeleteZone
//
// PARAMS:
//   - request: the arguments to DeleteZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteZone(request *DeleteZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteZoneUri(VERSION_V1, util.StringValue(request.ZoneName))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ListLineGroup
//
// PARAMS:
//   - request: the arguments to ListLineGroup
//
// RETURNS:
//   - ListLineGroupResponse: The return type of the ListLineGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListLineGroup(request *ListLineGroupRequest) (*ListLineGroupResponse, error) {
	result := &ListLineGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListLineGroupUri(VERSION_V1)).
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
		WithURL(getListRecordUri(VERSION_V1, util.StringValue(request.ZoneName))).
		WithQueryParamFilter("rr", util.StringValue(request.Rr)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListZone
//
// PARAMS:
//   - request: the arguments to ListZone
//
// RETURNS:
//   - ListZoneResponse: The return type of the ListZone interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListZone(request *ListZoneRequest) (*ListZoneResponse, error) {
	result := &ListZoneResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListZoneUri(VERSION_V1)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RenewZone
//
// PARAMS:
//   - request: the arguments to RenewZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenewZone(request *RenewZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenewZoneUri(VERSION_V1, util.StringValue(request.Name))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateLineGroup
//
// PARAMS:
//   - request: the arguments to UpdateLineGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateLineGroup(request *UpdateLineGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateLineGroupUri(VERSION_V1, util.Int32Value(request.LineId))).
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
		WithURL(getUpdateRecordUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateRecordDisable
//
// PARAMS:
//   - request: the arguments to UpdateRecordDisable
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRecordDisable(request *UpdateRecordDisableRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRecordDisableUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParam("disable", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UpdateRecordEnable
//
// PARAMS:
//   - request: the arguments to UpdateRecordEnable
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRecordEnable(request *UpdateRecordEnableRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRecordEnableUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParam("enable", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UpgradeZone
//
// PARAMS:
//   - request: the arguments to UpgradeZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpgradeZone(request *UpgradeZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpgradeZoneUri(VERSION_V1)).
		WithQueryParam("upgradeToDiscount", "").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
