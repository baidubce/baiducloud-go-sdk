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

// BindTagImage
//
// PARAMS:
//   - request: the arguments to BindTagImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTagImage(request *BindTagImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTagImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		WithQueryParam("bind", "").
		WithBody(request).
		Do()
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

// CancelRemoteCopyImage
//
// PARAMS:
//   - request: the arguments to CancelRemoteCopyImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CancelRemoteCopyImage(request *CancelRemoteCopyImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelRemoteCopyImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		WithQueryParam("cancelRemoteCopy", "").
		Do()
}

// CreateImage
//
// PARAMS:
//   - request: the arguments to CreateImage
//
// RETURNS:
//   - CreateImageResponse: The return type of the CreateImage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateImage(request *CreateImageRequest) (*CreateImageResponse, error) {
	result := &CreateImageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateImageUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// DeleteImage
//
// PARAMS:
//   - request: the arguments to DeleteImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteImage(request *DeleteImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		Do()
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

// GetAvailableImagesBySpec
//
// PARAMS:
//   - request: the arguments to GetAvailableImagesBySpec
//
// RETURNS:
//   - GetAvailableImagesBySpecResponse: The return type of the GetAvailableImagesBySpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAvailableImagesBySpec(request *GetAvailableImagesBySpecRequest) (*GetAvailableImagesBySpecResponse, error) {
	result := &GetAvailableImagesBySpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAvailableImagesBySpecUri(VERSION_V2)).
		WithQueryParamFilter("spec", util.StringValue(request.Spec)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("osName", util.StringValue(request.OsName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// GetImage
//
// PARAMS:
//   - request: the arguments to GetImage
//
// RETURNS:
//   - GetImageResponse: The return type of the GetImage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetImage(request *GetImageRequest) (*GetImageResponse, error) {
	result := &GetImageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetImageUri(VERSION_V2, util.StringValue(request.ImageId))).
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

// ImportImage
//
// PARAMS:
//   - request: the arguments to ImportImage
//
// RETURNS:
//   - ImportImageResponse: The return type of the ImportImage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImportImage(request *ImportImageRequest) (*ImportImageResponse, error) {
	result := &ImportImageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImportImageUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListImages
//
// PARAMS:
//   - request: the arguments to ListImages
//
// RETURNS:
//   - ListImagesResponse: The return type of the ListImages interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListImages(request *ListImagesRequest) (*ListImagesResponse, error) {
	result := &ListImagesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListImagesUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("imageType", util.StringValue(request.ImageType)).
		WithQueryParamFilter("imageName", util.StringValue(request.ImageName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListOs
//
// PARAMS:
//   - request: the arguments to ListOs
//
// RETURNS:
//   - ListOsResponse: The return type of the ListOs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListOs(request *ListOsRequest) (*ListOsResponse, error) {
	result := &ListOsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListOsUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSharedUser
//
// PARAMS:
//   - request: the arguments to ListSharedUser
//
// RETURNS:
//   - ListSharedUserResponse: The return type of the ListSharedUser interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSharedUser(request *ListSharedUserRequest) (*ListSharedUserResponse, error) {
	result := &ListSharedUserResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSharedUserUri(VERSION_V2, util.StringValue(request.ImageId))).
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

// RemoteCopyImage
//
// PARAMS:
//   - request: the arguments to RemoteCopyImage
//
// RETURNS:
//   - RemoteCopyImageResponse: The return type of the RemoteCopyImage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RemoteCopyImage(request *RemoteCopyImageRequest) (*RemoteCopyImageResponse, error) {
	result := &RemoteCopyImageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoteCopyImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		WithQueryParam("remoteCopy", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RenameImage
//
// PARAMS:
//   - request: the arguments to RenameImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenameImage(request *RenameImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenameImageUri(VERSION_V2)).
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

// ShareImage
//
// PARAMS:
//   - request: the arguments to ShareImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ShareImage(request *ShareImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getShareImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		WithQueryParam("share", "").
		WithBody(request).
		Do()
}

// UnShareImage
//
// PARAMS:
//   - request: the arguments to UnShareImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnShareImage(request *UnShareImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUnShareImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		WithQueryParam("unshare", "").
		WithBody(request).
		Do()
}

// UnbindTagImage
//
// PARAMS:
//   - request: the arguments to UnbindTagImage
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTagImage(request *UnbindTagImageRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTagImageUri(VERSION_V2, util.StringValue(request.ImageId))).
		WithQueryParam("unbind", "").
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
