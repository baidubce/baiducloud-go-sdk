package eip

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"strings"
)

const (
	VERSION_V1 = "v1"
)

// AddEipGroupCount
//
// PARAMS:
//   - request: the arguments to AddEipGroupCount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddEipGroupCount(request *AddEipGroupCountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAddEipGroupCountUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddTbspAreaBlocking
//
// PARAMS:
//   - request: the arguments to AddTbspAreaBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddTbspAreaBlocking(request *AddTbspAreaBlockingRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddTbspAreaBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddTbspIpWhitelist
//
// PARAMS:
//   - request: the arguments to AddTbspIpWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddTbspIpWhitelist(request *AddTbspIpWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddTbspIpWhitelistUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// AddTbspProtocolBlocking
//
// PARAMS:
//   - request: the arguments to AddTbspProtocolBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddTbspProtocolBlocking(request *AddTbspProtocolBlockingRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddTbspProtocolBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ApplyForEip
//
// PARAMS:
//   - request: the arguments to ApplyForEip
//
// RETURNS:
//   - ApplyForEipResponse: The return type of the ApplyForEip interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApplyForEip(request *ApplyForEipRequest) (*ApplyForEipResponse, error) {
	result := &ApplyForEipResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApplyForEipUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BandwidthPackageInquiry
//
// PARAMS:
//   - request: the arguments to BandwidthPackageInquiry
//
// RETURNS:
//   - BandwidthPackageInquiryResponse: The return type of the BandwidthPackageInquiry interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BandwidthPackageInquiry(request *BandwidthPackageInquiryRequest) (*BandwidthPackageInquiryResponse, error) {
	result := &BandwidthPackageInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBandwidthPackageInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindEip
//
// PARAMS:
//   - request: the arguments to BindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindEip(request *BindEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// BindTbspProtectionObject
//
// PARAMS:
//   - request: the arguments to BindTbspProtectionObject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTbspProtectionObject(request *BindTbspProtectionObjectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTbspProtectionObjectUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CancelEipTransfer
//
// PARAMS:
//   - request: the arguments to CancelEipTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CancelEipTransfer(request *CancelEipTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCancelEipTransferUri(VERSION_V1)).
		WithQueryParam("cancel", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// CreateASharedTrafficPackage
//
// PARAMS:
//   - request: the arguments to CreateASharedTrafficPackage
//
// RETURNS:
//   - CreateASharedTrafficPackageResponse: The return type of the CreateASharedTrafficPackage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateASharedTrafficPackage(request *CreateASharedTrafficPackageRequest) (*CreateASharedTrafficPackageResponse, error) {
	result := &CreateASharedTrafficPackageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateASharedTrafficPackageUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEipBp
//
// PARAMS:
//   - request: the arguments to CreateEipBp
//
// RETURNS:
//   - CreateEipBpResponse: The return type of the CreateEipBp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEipBp(request *CreateEipBpRequest) (*CreateEipBpResponse, error) {
	result := &CreateEipBpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEipBpUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEipGroup
//
// PARAMS:
//   - request: the arguments to CreateEipGroup
//
// RETURNS:
//   - CreateEipGroupResponse: The return type of the CreateEipGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEipGroup(request *CreateEipGroupRequest) (*CreateEipGroupResponse, error) {
	result := &CreateEipGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEipGroupUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEipTransfer
//
// PARAMS:
//   - request: the arguments to CreateEipTransfer
//
// RETURNS:
//   - CreateEipTransferResponse: The return type of the CreateEipTransfer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEipTransfer(request *CreateEipTransferRequest) (*CreateEipTransferResponse, error) {
	result := &CreateEipTransferResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEipTransferUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTbsp
//
// PARAMS:
//   - request: the arguments to CreateTbsp
//
// RETURNS:
//   - CreateTbspResponse: The return type of the CreateTbsp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateTbsp(request *CreateTbspRequest) (*CreateTbspResponse, error) {
	result := &CreateTbspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTbspUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DetailTbsp
//
// PARAMS:
//   - request: the arguments to DetailTbsp
//
// RETURNS:
//   - DetailTbspResponse: The return type of the DetailTbsp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DetailTbsp(request *DetailTbspRequest) (*DetailTbspResponse, error) {
	result := &DetailTbspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDetailTbspUri(VERSION_V1, util.StringValue(request.Id))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DirectEip
//
// PARAMS:
//   - request: the arguments to DirectEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DirectEip(request *DirectEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDirectEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("direct", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// DisableTbspIpClean
//
// PARAMS:
//   - request: the arguments to DisableTbspIpClean
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableTbspIpClean(request *DisableTbspIpCleanRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisableTbspIpCleanUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("turnOffClean", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// EipBandwidthScalingCapacity
//
// PARAMS:
//   - request: the arguments to EipBandwidthScalingCapacity
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EipBandwidthScalingCapacity(request *EipBandwidthScalingCapacityRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipBandwidthScalingCapacityUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// EipInquiry
//
// PARAMS:
//   - request: the arguments to EipInquiry
//
// RETURNS:
//   - EipInquiryResponse: The return type of the EipInquiry interface.
//   - error: nil if success otherwise the specific error
func (c *Client) EipInquiry(request *EipInquiryRequest) (*EipInquiryResponse, error) {
	result := &EipInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EipPostpaidToPrepaid
//
// PARAMS:
//   - request: the arguments to EipPostpaidToPrepaid
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EipPostpaidToPrepaid(request *EipPostpaidToPrepaidRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipPostpaidToPrepaidUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParamFilter("action", "TO_PREPAY").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// EipRenewal
//
// PARAMS:
//   - request: the arguments to EipRenewal
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EipRenewal(request *EipRenewalRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipRenewalUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// EnableTbspIpClean
//
// PARAMS:
//   - request: the arguments to EnableTbspIpClean
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableTbspIpClean(request *EnableTbspIpCleanRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableTbspIpCleanUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("turnOnClean", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		Do()
}

// GetEipBp
//
// PARAMS:
//   - request: the arguments to GetEipBp
//
// RETURNS:
//   - GetEipBpResponse: The return type of the GetEipBp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetEipBp(request *GetEipBpRequest) (*GetEipBpResponse, error) {
	result := &GetEipBpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetEipBpUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEipGroup
//
// PARAMS:
//   - request: the arguments to GetEipGroup
//
// RETURNS:
//   - GetEipGroupResponse: The return type of the GetEipGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetEipGroup(request *GetEipGroupRequest) (*GetEipGroupResponse, error) {
	result := &GetEipGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetEipGroupUri(VERSION_V1, util.StringValue(request.Id))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListBaseDdos
//
// PARAMS:
//   - request: the arguments to ListBaseDdos
//
// RETURNS:
//   - ListBaseDdosResponse: The return type of the ListBaseDdos interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListBaseDdos(request *ListBaseDdosRequest) (*ListBaseDdosResponse, error) {
	result := &ListBaseDdosResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListBaseDdosUri(VERSION_V1)).
		WithQueryParamFilter("ips", util.StringValue(request.Ips)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListBaseDdosAttackRecord
//
// PARAMS:
//   - request: the arguments to ListBaseDdosAttackRecord
//
// RETURNS:
//   - ListBaseDdosAttackRecordResponse: The return type of the ListBaseDdosAttackRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListBaseDdosAttackRecord(request *ListBaseDdosAttackRecordRequest) (*ListBaseDdosAttackRecordResponse, error) {
	result := &ListBaseDdosAttackRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListBaseDdosAttackRecordUri(VERSION_V1, util.StringValue(request.Ip))).
		WithQueryParamFilter("startTime", util.StringValue(request.StartTime)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListEipBp
//
// PARAMS:
//   - request: the arguments to ListEipBp
//
// RETURNS:
//   - ListEipBpResponse: The return type of the ListEipBp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListEipBp(request *ListEipBpRequest) (*ListEipBpResponse, error) {
	result := &ListEipBpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListEipBpUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("bindType", util.StringValue(request.BindType)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListEipGroup
//
// PARAMS:
//   - request: the arguments to ListEipGroup
//
// RETURNS:
//   - ListEipGroupResponse: The return type of the ListEipGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListEipGroup(request *ListEipGroupRequest) (*ListEipGroupResponse, error) {
	result := &ListEipGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListEipGroupUri(VERSION_V1)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListEipTransfer
//
// PARAMS:
//   - request: the arguments to ListEipTransfer
//
// RETURNS:
//   - ListEipTransferResponse: The return type of the ListEipTransfer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListEipTransfer(request *ListEipTransferRequest) (*ListEipTransferResponse, error) {
	result := &ListEipTransferResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListEipTransferUri(VERSION_V1)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("direction", util.StringValue(request.Direction)).
		WithQueryParamFilter("transferId", util.StringValue(request.TransferId)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("fuzzyTransferId", util.StringValue(request.FuzzyTransferId)).
		WithQueryParamFilter("fuzzyInstanceId", util.StringValue(request.FuzzyInstanceId)).
		WithQueryParamFilter("fuzzyInstanceName", util.StringValue(request.FuzzyInstanceName)).
		WithQueryParamFilter("fuzzyInstanceIp", util.StringValue(request.FuzzyInstanceIp)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRecycleEips
//
// PARAMS:
//   - request: the arguments to ListRecycleEips
//
// RETURNS:
//   - ListRecycleEipsResponse: The return type of the ListRecycleEips interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRecycleEips(request *ListRecycleEipsRequest) (*ListRecycleEipsResponse, error) {
	result := &ListRecycleEipsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRecycleEipsUri(VERSION_V1)).
		WithQueryParamFilter("eip", util.StringValue(request.Eip)).
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

// ListTbsp
//
// PARAMS:
//   - request: the arguments to ListTbsp
//
// RETURNS:
//   - ListTbspResponse: The return type of the ListTbsp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTbsp(request *ListTbspRequest) (*ListTbspResponse, error) {
	result := &ListTbspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspUri(VERSION_V1)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTbspAreaBlocking
//
// PARAMS:
//   - request: the arguments to ListTbspAreaBlocking
//
// RETURNS:
//   - ListTbspAreaBlockingResponse: The return type of the ListTbspAreaBlocking interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTbspAreaBlocking(request *ListTbspAreaBlockingRequest) (*ListTbspAreaBlockingResponse, error) {
	result := &ListTbspAreaBlockingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspAreaBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTbspIpClean
//
// PARAMS:
//   - request: the arguments to ListTbspIpClean
//
// RETURNS:
//   - ListTbspIpCleanResponse: The return type of the ListTbspIpClean interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTbspIpClean(request *ListTbspIpCleanRequest) (*ListTbspIpCleanResponse, error) {
	result := &ListTbspIpCleanResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspIpCleanUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTbspIpWhitelist
//
// PARAMS:
//   - request: the arguments to ListTbspIpWhitelist
//
// RETURNS:
//   - ListTbspIpWhitelistResponse: The return type of the ListTbspIpWhitelist interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTbspIpWhitelist(request *ListTbspIpWhitelistRequest) (*ListTbspIpWhitelistResponse, error) {
	result := &ListTbspIpWhitelistResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspIpWhitelistUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("ipCidr", util.StringValue(request.IpCidr)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTbspProtocolBlocking
//
// PARAMS:
//   - request: the arguments to ListTbspProtocolBlocking
//
// RETURNS:
//   - ListTbspProtocolBlockingResponse: The return type of the ListTbspProtocolBlocking interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTbspProtocolBlocking(request *ListTbspProtocolBlockingRequest) (*ListTbspProtocolBlockingResponse, error) {
	result := &ListTbspProtocolBlockingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspProtocolBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListUnban
//
// PARAMS:
//   - request: the arguments to ListUnban
//
// RETURNS:
//   - ListUnbanResponse: The return type of the ListUnban interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListUnban(request *ListUnbanRequest) (*ListUnbanResponse, error) {
	result := &ListUnbanResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListUnbanUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyTbspIpCleanThreshold
//
// PARAMS:
//   - request: the arguments to ModifyTbspIpCleanThreshold
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTbspIpCleanThreshold(request *ModifyTbspIpCleanThresholdRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTbspIpCleanThresholdUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("modifyThreshold", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ModifyTbspIpProtectLevel
//
// PARAMS:
//   - request: the arguments to ModifyTbspIpProtectLevel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTbspIpProtectLevel(request *ModifyTbspIpProtectLevelRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTbspIpProtectLevelUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// MoveInEips
//
// PARAMS:
//   - request: the arguments to MoveInEips
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) MoveInEips(request *MoveInEipsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMoveInEipsUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("move_in", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// MoveOutEips
//
// PARAMS:
//   - request: the arguments to MoveOutEips
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) MoveOutEips(request *MoveOutEipsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMoveOutEipsUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("move_out", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// OptionalReleaseEip
//
// PARAMS:
//   - request: the arguments to OptionalReleaseEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) OptionalReleaseEip(request *OptionalReleaseEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getOptionalReleaseEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParamFilter("releaseToRecycle", util.BoolValue(request.ReleaseToRecycle)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// PurchaseReservedEipGroup
//
// PARAMS:
//   - request: the arguments to PurchaseReservedEipGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedEipGroup(request *PurchaseReservedEipGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPurchaseReservedEipGroupUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// QueryEipList
//
// PARAMS:
//   - request: the arguments to QueryEipList
//
// RETURNS:
//   - QueryEipListResponse: The return type of the QueryEipList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryEipList(request *QueryEipListRequest) (*QueryEipListResponse, error) {
	result := &QueryEipListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryEipListUri(VERSION_V1)).
		WithQueryParamFilter("ipVersion", util.StringValue(request.IpVersion)).
		WithQueryParamFilter("eip", util.StringValue(request.Eip)).
		WithQueryParamFilter("instanceType", util.StringValue(request.InstanceType)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("eipIds", strings.Join(util.PtrSliceToStringSlice(request.EipIds), ",")).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheDetailsOfSharedTrafficPackages
//
// PARAMS:
//   - request: the arguments to QueryTheDetailsOfSharedTrafficPackages
//
// RETURNS:
//   - QueryTheDetailsOfSharedTrafficPackagesResponse: The return type of the QueryTheDetailsOfSharedTrafficPackages interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheDetailsOfSharedTrafficPackages(request *QueryTheDetailsOfSharedTrafficPackagesRequest) (*QueryTheDetailsOfSharedTrafficPackagesResponse, error) {
	result := &QueryTheDetailsOfSharedTrafficPackagesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheDetailsOfSharedTrafficPackagesUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfSharedTrafficPackages
//
// PARAMS:
//   - request: the arguments to QueryTheListOfSharedTrafficPackages
//
// RETURNS:
//   - QueryTheListOfSharedTrafficPackagesResponse: The return type of the QueryTheListOfSharedTrafficPackages interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfSharedTrafficPackages(request *QueryTheListOfSharedTrafficPackagesRequest) (*QueryTheListOfSharedTrafficPackagesResponse, error) {
	result := &QueryTheListOfSharedTrafficPackagesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfSharedTrafficPackagesUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("id", util.StringValue(request.Id)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("deductPolicy", util.StringValue(request.DeductPolicy)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReceiveEipTransfer
//
// PARAMS:
//   - request: the arguments to ReceiveEipTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReceiveEipTransfer(request *ReceiveEipTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getReceiveEipTransferUri(VERSION_V1)).
		WithQueryParam("accept", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RefundEip
//
// PARAMS:
//   - request: the arguments to RefundEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RefundEip(request *RefundEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRefundEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RefundEipGroup
//
// PARAMS:
//   - request: the arguments to RefundEipGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RefundEipGroup(request *RefundEipGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRefundEipGroupUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RejectEipTransfer
//
// PARAMS:
//   - request: the arguments to RejectEipTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RejectEipTransfer(request *RejectEipTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRejectEipTransferUri(VERSION_V1)).
		WithQueryParam("reject", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ReleaseEip
//
// PARAMS:
//   - request: the arguments to ReleaseEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseEip(request *ReleaseEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParamFilter("releaseToRecycle", util.BoolValue(request.ReleaseToRecycle)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleaseEipBp
//
// PARAMS:
//   - request: the arguments to ReleaseEipBp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseEipBp(request *ReleaseEipBpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseEipBpUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleaseEipFromRecycle
//
// PARAMS:
//   - request: the arguments to ReleaseEipFromRecycle
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseEipFromRecycle(request *ReleaseEipFromRecycleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseEipFromRecycleUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// ReleaseEipGroup
//
// PARAMS:
//   - request: the arguments to ReleaseEipGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseEipGroup(request *ReleaseEipGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseEipGroupUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RemoveTbspAreaBlocking
//
// PARAMS:
//   - request: the arguments to RemoveTbspAreaBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveTbspAreaBlocking(request *RemoveTbspAreaBlockingRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveTbspAreaBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("blockType", util.StringValue(request.BlockType)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RemoveTbspIpWhitelist
//
// PARAMS:
//   - request: the arguments to RemoveTbspIpWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveTbspIpWhitelist(request *RemoveTbspIpWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveTbspIpWhitelistUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("whitelistId", util.StringValue(request.WhitelistId)).
		Do()
}

// RemoveTbspProtocolBlocking
//
// PARAMS:
//   - request: the arguments to RemoveTbspProtocolBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveTbspProtocolBlocking(request *RemoveTbspProtocolBlockingRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveTbspProtocolBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// RenewTbsp
//
// PARAMS:
//   - request: the arguments to RenewTbsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenewTbsp(request *RenewTbspRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenewTbspUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ResizeEipBpBandwidth
//
// PARAMS:
//   - request: the arguments to ResizeEipBpBandwidth
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeEipBpBandwidth(request *ResizeEipBpBandwidthRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeEipBpBandwidthUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ResizeEipGroupBandwidth
//
// PARAMS:
//   - request: the arguments to ResizeEipGroupBandwidth
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeEipGroupBandwidth(request *ResizeEipGroupBandwidthRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeEipGroupBandwidthUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// ResizeTbsp
//
// PARAMS:
//   - request: the arguments to ResizeTbsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeTbsp(request *ResizeTbspRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeTbspUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// RestoreEipFromRecycle
//
// PARAMS:
//   - request: the arguments to RestoreEipFromRecycle
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestoreEipFromRecycle(request *RestoreEipFromRecycleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRestoreEipFromRecycleUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("restore", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// SharedBandwidthInquiry
//
// PARAMS:
//   - request: the arguments to SharedBandwidthInquiry
//
// RETURNS:
//   - SharedBandwidthInquiryResponse: The return type of the SharedBandwidthInquiry interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SharedBandwidthInquiry(request *SharedBandwidthInquiryRequest) (*SharedBandwidthInquiryResponse, error) {
	result := &SharedBandwidthInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSharedBandwidthInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SharedDataPackageInquiry
//
// PARAMS:
//   - request: the arguments to SharedDataPackageInquiry
//
// RETURNS:
//   - SharedDataPackageInquiryResponse: The return type of the SharedDataPackageInquiry interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SharedDataPackageInquiry(request *SharedDataPackageInquiryRequest) (*SharedDataPackageInquiryResponse, error) {
	result := &SharedDataPackageInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSharedDataPackageInquiryUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StartEipAutoRenew
//
// PARAMS:
//   - request: the arguments to StartEipAutoRenew
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StartEipAutoRenew(request *StartEipAutoRenewRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStartEipAutoRenewUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("startAutoRenew", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// StopEipAutoRenew
//
// PARAMS:
//   - request: the arguments to StopEipAutoRenew
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StopEipAutoRenew(request *StopEipAutoRenewRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStopEipAutoRenewUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("stopAutoRenew", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UnDirectEip
//
// PARAMS:
//   - request: the arguments to UnDirectEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnDirectEip(request *UnDirectEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnDirectEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("unDirect", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UnbindEip
//
// PARAMS:
//   - request: the arguments to UnbindEip
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindEip(request *UnbindEipRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindEipUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
}

// UnbindTbspProtectionObject
//
// PARAMS:
//   - request: the arguments to UnbindTbspProtectionObject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTbspProtectionObject(request *UnbindTbspProtectionObjectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTbspProtectionObjectUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateBaseDdosThreshold
//
// PARAMS:
//   - request: the arguments to UpdateBaseDdosThreshold
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateBaseDdosThreshold(request *UpdateBaseDdosThresholdRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateBaseDdosThresholdUri(VERSION_V1, util.StringValue(request.Ip))).
		WithQueryParam("modifyThreshold", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEipBpAutoReleaseTime
//
// PARAMS:
//   - request: the arguments to UpdateEipBpAutoReleaseTime
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEipBpAutoReleaseTime(request *UpdateEipBpAutoReleaseTimeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEipBpAutoReleaseTimeUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("retime", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEipBpName
//
// PARAMS:
//   - request: the arguments to UpdateEipBpName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEipBpName(request *UpdateEipBpNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEipBpNameUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("rename", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEipDeleteProtect
//
// PARAMS:
//   - request: the arguments to UpdateEipDeleteProtect
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEipDeleteProtect(request *UpdateEipDeleteProtectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEipDeleteProtectUri(VERSION_V1, util.StringValue(request.Eip))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}

// UpdateEipGroup
//
// PARAMS:
//   - request: the arguments to UpdateEipGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateEipGroup(request *UpdateEipGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateEipGroupUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("update", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
}
