package privatezone

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddParsingRecords
//
// PARAMS:
//   - request: the arguments to AddParsingRecords
//
// RETURNS:
//   - AddParsingRecordsResponse: The return type of the AddParsingRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddParsingRecords(request *AddParsingRecordsRequest) (*AddParsingRecordsResponse, error) {
	result := &AddParsingRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddParsingRecordsUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
}

// AssociateVpc
//
// PARAMS:
//   - request: the arguments to AssociateVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AssociateVpc(request *AssociateVpcRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAssociateVpcUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParam(util.StringValue(request.Action), "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// CreateAPrivateZone
//
// PARAMS:
//   - request: the arguments to CreateAPrivateZone
//
// RETURNS:
//   - CreateAPrivateZoneResponse: The return type of the CreateAPrivateZone interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAPrivateZone(request *CreateAPrivateZoneRequest) (*CreateAPrivateZoneResponse, error) {
	result := &CreateAPrivateZoneResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAPrivateZoneUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
}

// DeleteParsingRecords
//
// PARAMS:
//   - request: the arguments to DeleteParsingRecords
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteParsingRecords(request *DeleteParsingRecordsRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteParsingRecordsUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DeletePrivateZone
//
// PARAMS:
//   - request: the arguments to DeletePrivateZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePrivateZone(request *DeletePrivateZoneRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePrivateZoneUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// DisassociateVpc
//
// PARAMS:
//   - request: the arguments to DisassociateVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisassociateVpc(request *DisassociateVpcRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisassociateVpcUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParam(util.StringValue(request.Action), "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// ModifyParsingRecords
//
// PARAMS:
//   - request: the arguments to ModifyParsingRecords
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyParsingRecords(request *ModifyParsingRecordsRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyParsingRecordsUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// QueryAndParseRecordList
//
// PARAMS:
//   - request: the arguments to QueryAndParseRecordList
//
// RETURNS:
//   - QueryAndParseRecordListResponse: The return type of the QueryAndParseRecordList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAndParseRecordList(request *QueryAndParseRecordListRequest) (*QueryAndParseRecordListResponse, error) {
	result := &QueryAndParseRecordListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAndParseRecordListUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("rr", util.StringValue(request.Rr)).
		WithQueryParamFilter("searchMode", util.StringValue(request.SearchMode)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithQueryParamFilter("value", util.StringValue(request.Value)).
		WithResult(result).
		Do()
	return result, err
}

// QueryTheListOfPrivateZones
//
// PARAMS:
//   - request: the arguments to QueryTheListOfPrivateZones
//
// RETURNS:
//   - QueryTheListOfPrivateZonesResponse: The return type of the QueryTheListOfPrivateZones interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfPrivateZones(request *QueryTheListOfPrivateZonesRequest) (*QueryTheListOfPrivateZonesResponse, error) {
	result := &QueryTheListOfPrivateZonesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfPrivateZonesUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	return result, err
}

// SearchForDetailsOfPrivatzone
//
// PARAMS:
//   - request: the arguments to SearchForDetailsOfPrivatzone
//
// RETURNS:
//   - SearchForDetailsOfPrivatzoneResponse: The return type of the SearchForDetailsOfPrivatzone interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SearchForDetailsOfPrivatzone(request *SearchForDetailsOfPrivatzoneRequest) (*SearchForDetailsOfPrivatzoneResponse, error) {
	result := &SearchForDetailsOfPrivatzoneResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSearchForDetailsOfPrivatzoneUri(VERSION_V1, util.StringValue(request.ZoneId))).
		WithResult(result).
		Do()
	return result, err
}

// SetParsingRecordStatus
//
// PARAMS:
//   - request: the arguments to SetParsingRecordStatus
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetParsingRecordStatus(request *SetParsingRecordStatusRequest) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSetParsingRecordStatusUri(VERSION_V1, util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParam(util.StringValue(request.Action), "").
		Do()
	return err
}
