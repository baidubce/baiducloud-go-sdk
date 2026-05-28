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
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		Do()
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
		WithQueryParamFilter("action", util.StringValue(request.Action)).
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
