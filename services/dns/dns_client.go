package dns

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddDomainName
//
// PARAMS:
//   - request: the arguments to AddDomainName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddDomainName(request *AddDomainNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddDomainNameUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

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

// AddParsingRecords
//
// PARAMS:
//   - request: the arguments to AddParsingRecords
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddParsingRecords(request *AddParsingRecordsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddParsingRecordsUri(VERSION_V1, util.StringValue(request.ZoneName))).
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

// DeleteParsingRecords
//
// PARAMS:
//   - request: the arguments to DeleteParsingRecords
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteParsingRecords(request *DeleteParsingRecordsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteParsingRecordsUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DomainNameRenewal
//
// PARAMS:
//   - request: the arguments to DomainNameRenewal
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DomainNameRenewal(request *DomainNameRenewalRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDomainNameRenewalUri(VERSION_V1, util.StringValue(request.Name))).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ModifyParsingRecords
//
// PARAMS:
//   - request: the arguments to ModifyParsingRecords
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyParsingRecords(request *ModifyParsingRecordsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyParsingRecordsUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ModifyTheParsingRecordStatus
//
// PARAMS:
//   - request: the arguments to ModifyTheParsingRecordStatus
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTheParsingRecordStatus(request *ModifyTheParsingRecordStatusRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTheParsingRecordStatusUri(VERSION_V1, util.StringValue(request.ZoneName), util.StringValue(request.RecordId))).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// PurchaseAPaidDomainName
//
// PARAMS:
//   - request: the arguments to PurchaseAPaidDomainName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PurchaseAPaidDomainName(request *PurchaseAPaidDomainNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPurchaseAPaidDomainNameUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
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
		WithURL(getQueryAndParseRecordListUri(VERSION_V1, util.StringValue(request.ZoneName))).
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

// QueryDomainNameList
//
// PARAMS:
//   - request: the arguments to QueryDomainNameList
//
// RETURNS:
//   - QueryDomainNameListResponse: The return type of the QueryDomainNameList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryDomainNameList(request *QueryDomainNameListRequest) (*QueryDomainNameListResponse, error) {
	result := &QueryDomainNameListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryDomainNameListUri(VERSION_V1)).
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

// QueryTheListOfLineGroups
//
// PARAMS:
//   - request: the arguments to QueryTheListOfLineGroups
//
// RETURNS:
//   - QueryTheListOfLineGroupsResponse: The return type of the QueryTheListOfLineGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfLineGroups(request *QueryTheListOfLineGroupsRequest) (*QueryTheListOfLineGroupsResponse, error) {
	result := &QueryTheListOfLineGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfLineGroupsUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveDomainName
//
// PARAMS:
//   - request: the arguments to RemoveDomainName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveDomainName(request *RemoveDomainNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveDomainNameUri(VERSION_V1, util.StringValue(request.ZoneName))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
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

// UpgradeTheFreeDomainNameToTheUniversalVersion
//
// PARAMS:
//   - request: the arguments to UpgradeTheFreeDomainNameToTheUniversalVersion
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpgradeTheFreeDomainNameToTheUniversalVersion(request *UpgradeTheFreeDomainNameToTheUniversalVersionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpgradeTheFreeDomainNameToTheUniversalVersionUri(VERSION_V1)).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
