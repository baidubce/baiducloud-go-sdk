package eip

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddTbspAreaBlocking
//
// PARAMS:
//   - request: the arguments to AddTbspAreaBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddTbspAreaBlocking(request *AddTbspAreaBlockingRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddTbspAreaBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// AddTbspIpWhitelist
//
// PARAMS:
//   - request: the arguments to AddTbspIpWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddTbspIpWhitelist(request *AddTbspIpWhitelistRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddTbspIpWhitelistUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// AddTbspProtocolBlocking
//
// PARAMS:
//   - request: the arguments to AddTbspProtocolBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddTbspProtocolBlocking(request *AddTbspProtocolBlockingRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddTbspProtocolBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
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
	if request.BandwidthInMbps == nil {
		return nil, fmt.Errorf("BandwidthInMbps is required and must be specified")
	}
	result := &BandwidthPackageInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBandwidthPackageInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
}

// BindTbspProtectionObject
//
// PARAMS:
//   - request: the arguments to BindTbspProtectionObject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTbspProtectionObject(request *BindTbspProtectionObjectRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	if len(request.IpList) == 0 {
		return fmt.Errorf("IpList is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTbspProtectionObjectUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("bind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
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
	if request.Name == nil {
		return nil, fmt.Errorf("Name is required and must be specified")
	}
	if request.LineType == nil {
		return nil, fmt.Errorf("LineType is required and must be specified")
	}
	if request.IpCapacity == nil {
		return nil, fmt.Errorf("IpCapacity is required and must be specified")
	}
	if request.ReservationLength == nil {
		return nil, fmt.Errorf("ReservationLength is required and must be specified")
	}
	if request.ReservationTimeUnit == nil {
		return nil, fmt.Errorf("ReservationTimeUnit is required and must be specified")
	}
	result := &CreateTbspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTbspUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
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
	if request.Id == nil {
		return nil, fmt.Errorf("id is required and must be specified")
	}
	result := &DetailTbspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDetailTbspUri(VERSION_V1, util.StringValue(request.Id))).
		WithResult(result).
		Do()
	return result, err
}

// DisableTbspIpClean
//
// PARAMS:
//   - request: the arguments to DisableTbspIpClean
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableTbspIpClean(request *DisableTbspIpCleanRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisableTbspIpCleanUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("turnOffClean", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
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
	if request.BandwidthInMbps == nil {
		return nil, fmt.Errorf("BandwidthInMbps is required and must be specified")
	}
	if request.Billing == nil {
		return nil, fmt.Errorf("Billing is required and must be specified")
	}
	result := &EipInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
}

// EnableTbspIpClean
//
// PARAMS:
//   - request: the arguments to EnableTbspIpClean
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableTbspIpClean(request *EnableTbspIpCleanRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableTbspIpCleanUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("turnOnClean", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		Do()
	return err
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
	return result, err
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
	if request.Id == nil {
		return nil, fmt.Errorf("id is required and must be specified")
	}
	result := &ListTbspAreaBlockingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspAreaBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithResult(result).
		Do()
	return result, err
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
	if request.Id == nil {
		return nil, fmt.Errorf("id is required and must be specified")
	}
	result := &ListTbspIpCleanResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspIpCleanUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	return result, err
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
	if request.Id == nil {
		return nil, fmt.Errorf("id is required and must be specified")
	}
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
	return result, err
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
	if request.Id == nil {
		return nil, fmt.Errorf("id is required and must be specified")
	}
	result := &ListTbspProtocolBlockingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTbspProtocolBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithResult(result).
		Do()
	return result, err
}

// ModifyTbspIpCleanThreshold
//
// PARAMS:
//   - request: the arguments to ModifyTbspIpCleanThreshold
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTbspIpCleanThreshold(request *ModifyTbspIpCleanThresholdRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTbspIpCleanThresholdUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("modifyThreshold", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// ModifyTbspIpProtectLevel
//
// PARAMS:
//   - request: the arguments to ModifyTbspIpProtectLevel
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTbspIpProtectLevel(request *ModifyTbspIpProtectLevelRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTbspIpProtectLevelUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// RemoveTbspAreaBlocking
//
// PARAMS:
//   - request: the arguments to RemoveTbspAreaBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveTbspAreaBlocking(request *RemoveTbspAreaBlockingRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveTbspAreaBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("blockType", util.StringValue(request.BlockType)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// RemoveTbspIpWhitelist
//
// PARAMS:
//   - request: the arguments to RemoveTbspIpWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveTbspIpWhitelist(request *RemoveTbspIpWhitelistRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveTbspIpWhitelistUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("whitelistId", util.StringValue(request.WhitelistId)).
		Do()
	return err
}

// RemoveTbspProtocolBlocking
//
// PARAMS:
//   - request: the arguments to RemoveTbspProtocolBlocking
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveTbspProtocolBlocking(request *RemoveTbspProtocolBlockingRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveTbspProtocolBlockingUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParamFilter("ip", util.StringValue(request.Ip)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		Do()
	return err
}

// RenewTbsp
//
// PARAMS:
//   - request: the arguments to RenewTbsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenewTbsp(request *RenewTbspRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	if request.RenewTime == nil {
		return fmt.Errorf("RenewTime is required and must be specified")
	}
	if request.RenewTimeUnit == nil {
		return fmt.Errorf("RenewTimeUnit is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenewTbspUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}

// ResizeTbsp
//
// PARAMS:
//   - request: the arguments to ResizeTbsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeTbsp(request *ResizeTbspRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	if request.IpCapacity == nil {
		return fmt.Errorf("IpCapacity is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeTbspUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
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
	if request.BandwidthInMbps == nil {
		return nil, fmt.Errorf("BandwidthInMbps is required and must be specified")
	}
	if request.IpNum == nil {
		return nil, fmt.Errorf("IpNum is required and must be specified")
	}
	if request.Billing == nil {
		return nil, fmt.Errorf("Billing is required and must be specified")
	}
	result := &SharedBandwidthInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSharedBandwidthInquiryUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
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
	if request.ReservationLength == nil {
		return nil, fmt.Errorf("ReservationLength is required and must be specified")
	}
	if request.Capacity == nil {
		return nil, fmt.Errorf("Capacity is required and must be specified")
	}
	result := &SharedDataPackageInquiryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSharedDataPackageInquiryUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
}

// UnbindTbspProtectionObject
//
// PARAMS:
//   - request: the arguments to UnbindTbspProtectionObject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTbspProtectionObject(request *UnbindTbspProtectionObjectRequest) error {
	if request.Id == nil {
		return fmt.Errorf("id is required and must be specified")
	}
	if len(request.IpList) == 0 {
		return fmt.Errorf("IpList is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTbspProtectionObjectUri(VERSION_V1, util.StringValue(request.Id))).
		WithQueryParam("unbind", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		Do()
	return err
}
