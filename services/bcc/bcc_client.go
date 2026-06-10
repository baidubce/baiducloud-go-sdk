package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V2 = "v2"
)

// AttachVolume
//
// PARAMS:
//   - request: the arguments to AttachVolume
//
// RETURNS:
//   - AttachVolumeResponse: The return type of the AttachVolume interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AttachVolume(request *AttachVolumeRequest) (*AttachVolumeResponse, error) {
	result := &AttachVolumeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAttachVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("attach", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindTagVolume
//
// PARAMS:
//   - request: the arguments to BindTagVolume
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTagVolume(request *BindTagVolumeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTagVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("bind", "").
		WithBody(request).
		Do()
}

// CreateVolume
//
// PARAMS:
//   - request: the arguments to CreateVolume
//
// RETURNS:
//   - CreateVolumeResponse: The return type of the CreateVolume interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVolume(request *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	result := &CreateVolumeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateVolumeUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DetachVolume
//
// PARAMS:
//   - request: the arguments to DetachVolume
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DetachVolume(request *DetachVolumeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDetachVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("detach", "").
		WithBody(request).
		Do()
}

// GetCdsPrice
//
// PARAMS:
//   - request: the arguments to GetCdsPrice
//
// RETURNS:
//   - GetCdsPriceResponse: The return type of the GetCdsPrice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetCdsPrice(request *GetCdsPriceRequest) (*GetCdsPriceResponse, error) {
	result := &GetCdsPriceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetCdsPriceUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDiskQuota
//
// PARAMS:
//   - request: the arguments to GetDiskQuota
//
// RETURNS:
//   - GetDiskQuotaResponse: The return type of the GetDiskQuota interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDiskQuota(request *GetDiskQuotaRequest) (*GetDiskQuotaResponse, error) {
	result := &GetDiskQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDiskQuotaUri(VERSION_V2)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetVolume
//
// PARAMS:
//   - request: the arguments to GetVolume
//
// RETURNS:
//   - GetVolumeResponse: The return type of the GetVolume interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetVolume(request *GetVolumeRequest) (*GetVolumeResponse, error) {
	result := &GetVolumeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetVolumeResizeProgress
//
// PARAMS:
//   - request: the arguments to GetVolumeResizeProgress
//
// RETURNS:
//   - GetVolumeResizeProgressResponse: The return type of the GetVolumeResizeProgress interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetVolumeResizeProgress(request *GetVolumeResizeProgressRequest) (*GetVolumeResizeProgressResponse, error) {
	result := &GetVolumeResizeProgressResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetVolumeResizeProgressUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListVolumes
//
// PARAMS:
//   - request: the arguments to ListVolumes
//
// RETURNS:
//   - ListVolumesResponse: The return type of the ListVolumes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListVolumes(request *ListVolumesRequest) (*ListVolumesResponse, error) {
	result := &ListVolumesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListVolumesUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("clusterId", util.StringValue(request.ClusterId)).
		WithQueryParamFilter("chargeFilter", util.StringValue(request.ChargeFilter)).
		WithQueryParamFilter("usageFilter", util.StringValue(request.UsageFilter)).
		WithQueryParamFilter("name", util.StringValue(request.Name)).
		WithQueryParamFilter("productCategory", util.StringValue(request.ProductCategory)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifyCdsAttribute
//
// PARAMS:
//   - request: the arguments to ModifyCdsAttribute
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyCdsAttribute(request *ModifyCdsAttributeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyCdsAttributeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("modify", "").
		WithBody(request).
		Do()
}

// ModifyVolumeChargeType
//
// PARAMS:
//   - request: the arguments to ModifyVolumeChargeType
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyVolumeChargeType(request *ModifyVolumeChargeTypeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyVolumeChargeTypeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("modifyChargeType", "").
		WithBody(request).
		Do()
}

// PurchaseReservedVolume
//
// PARAMS:
//   - request: the arguments to PurchaseReservedVolume
//
// RETURNS:
//   - PurchaseReservedVolumeResponse: The return type of the PurchaseReservedVolume interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedVolume(request *PurchaseReservedVolumeRequest) (*PurchaseReservedVolumeResponse, error) {
	result := &PurchaseReservedVolumeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPurchaseReservedVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("purchaseReserved", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseVolume
//
// PARAMS:
//   - request: the arguments to ReleaseVolume
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseVolume(request *ReleaseVolumeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getReleaseVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithBody(request).
		Do()
}

// RenameVolume
//
// PARAMS:
//   - request: the arguments to RenameVolume
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenameVolume(request *RenameVolumeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenameVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("rename", "").
		WithBody(request).
		Do()
}

// ResizeVolume
//
// PARAMS:
//   - request: the arguments to ResizeVolume
//
// RETURNS:
//   - ResizeVolumeResponse: The return type of the ResizeVolume interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeVolume(request *ResizeVolumeRequest) (*ResizeVolumeResponse, error) {
	result := &ResizeVolumeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RollbackVolume
//
// PARAMS:
//   - request: the arguments to RollbackVolume
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RollbackVolume(request *RollbackVolumeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRollbackVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("rollback", "").
		WithBody(request).
		Do()
}

// UnbindTagVolume
//
// PARAMS:
//   - request: the arguments to UnbindTagVolume
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTagVolume(request *UnbindTagVolumeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTagVolumeUri(VERSION_V2, util.StringValue(request.VolumeId))).
		WithQueryParam("unbind", "").
		WithBody(request).
		Do()
}
