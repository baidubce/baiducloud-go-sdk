package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V2 = "v2"
)

// AddIpv6
//
// PARAMS:
//   - request: the arguments to AddIpv6
//
// RETURNS:
//   - AddIpv6Response: The return type of the AddIpv6 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AddIpv6(request *AddIpv6Request) (*AddIpv6Response, error) {
	result := &AddIpv6Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddIpv6Uri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

// AutoReleaseInstance
//
// PARAMS:
//   - request: the arguments to AutoReleaseInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AutoReleaseInstance(request *AutoReleaseInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAutoReleaseInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("autorelease", "").
		WithBody(request).
		Do()
}

// BatchAddIp
//
// PARAMS:
//   - request: the arguments to BatchAddIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchAddIp(request *BatchAddIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBatchAddIpUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// BatchChangeToPostpaid
//
// PARAMS:
//   - request: the arguments to BatchChangeToPostpaid
//
// RETURNS:
//   - BatchChangeToPostpaidResponse: The return type of the BatchChangeToPostpaid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchChangeToPostpaid(request *BatchChangeToPostpaidRequest) (*BatchChangeToPostpaidResponse, error) {
	result := &BatchChangeToPostpaidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchChangeToPostpaidUri(VERSION_V2)).
		WithQueryParam("toPostpay", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchChangeToPrepaid
//
// PARAMS:
//   - request: the arguments to BatchChangeToPrepaid
//
// RETURNS:
//   - BatchChangeToPrepaidResponse: The return type of the BatchChangeToPrepaid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchChangeToPrepaid(request *BatchChangeToPrepaidRequest) (*BatchChangeToPrepaidResponse, error) {
	result := &BatchChangeToPrepaidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchChangeToPrepaidUri(VERSION_V2)).
		WithQueryParam("toPrepay", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchDeleteIp
//
// PARAMS:
//   - request: the arguments to BatchDeleteIp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteIp(request *BatchDeleteIpRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBatchDeleteIpUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// BatchRefundResource
//
// PARAMS:
//   - request: the arguments to BatchRefundResource
//
// RETURNS:
//   - BatchRefundResourceResponse: The return type of the BatchRefundResource interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchRefundResource(request *BatchRefundResourceRequest) (*BatchRefundResourceResponse, error) {
	result := &BatchRefundResourceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchRefundResourceUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchStartInstance
//
// PARAMS:
//   - request: the arguments to BatchStartInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchStartInstance(request *BatchStartInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBatchStartInstanceUri(VERSION_V2)).
		WithQueryParam("start", "").
		WithBody(request).
		Do()
}

// BatchStopInstance
//
// PARAMS:
//   - request: the arguments to BatchStopInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchStopInstance(request *BatchStopInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBatchStopInstanceUri(VERSION_V2)).
		WithQueryParam("stop", "").
		WithBody(request).
		Do()
}

// BindInstanceToSecurityGroup
//
// PARAMS:
//   - request: the arguments to BindInstanceToSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindInstanceToSecurityGroup(request *BindInstanceToSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindInstanceToSecurityGroupUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("bind", "").
		WithBody(request).
		Do()
}

// BindInstanceToTags
//
// PARAMS:
//   - request: the arguments to BindInstanceToTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindInstanceToTags(request *BindInstanceToTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindInstanceToTagsUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("bind", "").
		WithBody(request).
		Do()
}

// BindRole
//
// PARAMS:
//   - request: the arguments to BindRole
//
// RETURNS:
//   - BindRoleResponse: The return type of the BindRole interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BindRole(request *BindRoleRequest) (*BindRoleResponse, error) {
	result := &BindRoleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBindRoleUri(VERSION_V2)).
		WithQueryParam("bind", "").
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

// ChangeToPrepaid
//
// PARAMS:
//   - request: the arguments to ChangeToPrepaid
//
// RETURNS:
//   - ChangeToPrepaidResponse: The return type of the ChangeToPrepaid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ChangeToPrepaid(request *ChangeToPrepaidRequest) (*ChangeToPrepaidResponse, error) {
	result := &ChangeToPrepaidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getChangeToPrepaidUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("toPrepay", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChangeVpc
//
// PARAMS:
//   - request: the arguments to ChangeVpc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ChangeVpc(request *ChangeVpcRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getChangeVpcUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// CreateAutoRenewRule
//
// PARAMS:
//   - request: the arguments to CreateAutoRenewRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAutoRenewRule(request *CreateAutoRenewRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAutoRenewRuleUri(VERSION_V2)).
		WithBody(request).
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

// CreateInstanceBySpec
//
// PARAMS:
//   - request: the arguments to CreateInstanceBySpec
//
// RETURNS:
//   - CreateInstanceBySpecResponse: The return type of the CreateInstanceBySpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceBySpec(request *CreateInstanceBySpecRequest) (*CreateInstanceBySpecResponse, error) {
	result := &CreateInstanceBySpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceBySpecUri(VERSION_V2)).
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

// DelIpv6
//
// PARAMS:
//   - request: the arguments to DelIpv6
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DelIpv6(request *DelIpv6Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDelIpv6Uri(VERSION_V2)).
		WithBody(request).
		Do()
}

// DeleteAutoRenewRule
//
// PARAMS:
//   - request: the arguments to DeleteAutoRenewRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAutoRenewRule(request *DeleteAutoRenewRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAutoRenewRuleUri(VERSION_V2)).
		WithBody(request).
		Do()
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

// DeleteInstanceDeploySet
//
// PARAMS:
//   - request: the arguments to DeleteInstanceDeploySet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceDeploySet(request *DeleteInstanceDeploySetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteInstanceDeploySetUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// DeletePrepayInstance
//
// PARAMS:
//   - request: the arguments to DeletePrepayInstance
//
// RETURNS:
//   - DeletePrepayInstanceResponse: The return type of the DeletePrepayInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeletePrepayInstance(request *DeletePrepayInstanceRequest) (*DeletePrepayInstanceResponse, error) {
	result := &DeletePrepayInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeletePrepayInstanceUri(VERSION_V2)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteRecycledInstance
//
// PARAMS:
//   - request: the arguments to DeleteRecycledInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRecycledInstance(request *DeleteRecycledInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRecycledInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
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

// EnterRescueMode
//
// PARAMS:
//   - request: the arguments to EnterRescueMode
//
// RETURNS:
//   - EnterRescueModeResponse: The return type of the EnterRescueMode interface.
//   - error: nil if success otherwise the specific error
func (c *Client) EnterRescueMode(request *EnterRescueModeRequest) (*EnterRescueModeResponse, error) {
	result := &EnterRescueModeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnterRescueModeUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExitRescueMode
//
// PARAMS:
//   - request: the arguments to ExitRescueMode
//
// RETURNS:
//   - ExitRescueModeResponse: The return type of the ExitRescueMode interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ExitRescueMode(request *ExitRescueModeRequest) (*ExitRescueModeResponse, error) {
	result := &ExitRescueModeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getExitRescueModeUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// GetInstance
//
// PARAMS:
//   - request: the arguments to GetInstance
//
// RETURNS:
//   - GetInstanceResponse: The return type of the GetInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstance(request *GetInstanceRequest) (*GetInstanceResponse, error) {
	result := &GetInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceNoChargeList
//
// PARAMS:
//   - request: the arguments to GetInstanceNoChargeList
//
// RETURNS:
//   - GetInstanceNoChargeListResponse: The return type of the GetInstanceNoChargeList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceNoChargeList(request *GetInstanceNoChargeListRequest) (*GetInstanceNoChargeListResponse, error) {
	result := &GetInstanceNoChargeListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceNoChargeListUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("internalIp", util.StringValue(request.InternalIp)).
		WithQueryParamFilter("keypairId", util.StringValue(request.KeypairId)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("instanceIds", util.StringValue(request.InstanceIds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceUserDataInfo
//
// PARAMS:
//   - request: the arguments to GetInstanceUserDataInfo
//
// RETURNS:
//   - GetInstanceUserDataInfoResponse: The return type of the GetInstanceUserDataInfo interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceUserDataInfo(request *GetInstanceUserDataInfoRequest) (*GetInstanceUserDataInfoResponse, error) {
	result := &GetInstanceUserDataInfoResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetInstanceUserDataInfoUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceVnc
//
// PARAMS:
//   - request: the arguments to GetInstanceVnc
//
// RETURNS:
//   - GetInstanceVncResponse: The return type of the GetInstanceVnc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceVnc(request *GetInstanceVncRequest) (*GetInstanceVncResponse, error) {
	result := &GetInstanceVncResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceVncUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRoleList
//
// PARAMS:
//   - request: the arguments to GetRoleList
//
// RETURNS:
//   - GetRoleListResponse: The return type of the GetRoleList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRoleList() (*GetRoleListResponse, error) {
	result := &GetRoleListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRoleListUri(VERSION_V2)).
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

// InstanceBatchResizeBySpec
//
// PARAMS:
//   - request: the arguments to InstanceBatchResizeBySpec
//
// RETURNS:
//   - InstanceBatchResizeBySpecResponse: The return type of the InstanceBatchResizeBySpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceBatchResizeBySpec(request *InstanceBatchResizeBySpecRequest) (*InstanceBatchResizeBySpecResponse, error) {
	result := &InstanceBatchResizeBySpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceBatchResizeBySpecUri(VERSION_V2)).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InstanceDeletionProtection
//
// PARAMS:
//   - request: the arguments to InstanceDeletionProtection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) InstanceDeletionProtection(request *InstanceDeletionProtectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceDeletionProtectionUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// InstanceRecovery
//
// PARAMS:
//   - request: the arguments to InstanceRecovery
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) InstanceRecovery(request *InstanceRecoveryRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceRecoveryUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// ListAvailableResizeSpec
//
// PARAMS:
//   - request: the arguments to ListAvailableResizeSpec
//
// RETURNS:
//   - ListAvailableResizeSpecResponse: The return type of the ListAvailableResizeSpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAvailableResizeSpec(request *ListAvailableResizeSpecRequest) (*ListAvailableResizeSpecResponse, error) {
	result := &ListAvailableResizeSpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListAvailableResizeSpecUri(VERSION_V2)).
		WithQueryParam("resizeList", "").
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

// ListInstanceByIds
//
// PARAMS:
//   - request: the arguments to ListInstanceByIds
//
// RETURNS:
//   - ListInstanceByIdsResponse: The return type of the ListInstanceByIds interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstanceByIds(request *ListInstanceByIdsRequest) (*ListInstanceByIdsResponse, error) {
	result := &ListInstanceByIdsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListInstanceByIdsUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstanceEnis
//
// PARAMS:
//   - request: the arguments to ListInstanceEnis
//
// RETURNS:
//   - ListInstanceEnisResponse: The return type of the ListInstanceEnis interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstanceEnis(request *ListInstanceEnisRequest) (*ListInstanceEnisResponse, error) {
	result := &ListInstanceEnisResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstanceEnisUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstances
//
// PARAMS:
//   - request: the arguments to ListInstances
//
// RETURNS:
//   - ListInstancesResponse: The return type of the ListInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(request *ListInstancesRequest) (*ListInstancesResponse, error) {
	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstancesUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("internalIp", util.StringValue(request.InternalIp)).
		WithQueryParamFilter("dedicatedHostId", util.StringValue(request.DedicatedHostId)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("showRdmaTopo", util.StringValue(request.ShowRdmaTopo)).
		WithQueryParamFilter("instanceIds", util.StringValue(request.InstanceIds)).
		WithQueryParamFilter("instanceNames", util.StringValue(request.InstanceNames)).
		WithQueryParamFilter("fuzzyInstanceName", util.StringValue(request.FuzzyInstanceName)).
		WithQueryParamFilter("volumeIds", util.StringValue(request.VolumeIds)).
		WithQueryParamFilter("deploySetIds", util.StringValue(request.DeploySetIds)).
		WithQueryParamFilter("securityGroupIds", util.StringValue(request.SecurityGroupIds)).
		WithQueryParamFilter("paymentTiming", util.StringValue(request.PaymentTiming)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("tags", util.StringValue(request.Tags)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("privateIps", util.StringValue(request.PrivateIps)).
		WithQueryParamFilter("ehcClusterId", util.StringValue(request.EhcClusterId)).
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

// ListRecycleInstance
//
// PARAMS:
//   - request: the arguments to ListRecycleInstance
//
// RETURNS:
//   - ListRecycleInstanceResponse: The return type of the ListRecycleInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRecycleInstance(request *ListRecycleInstanceRequest) (*ListRecycleInstanceResponse, error) {
	result := &ListRecycleInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListRecycleInstanceUri(VERSION_V2)).
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

// ModifyInstanceAttributes
//
// PARAMS:
//   - request: the arguments to ModifyInstanceAttributes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceAttributes(request *ModifyInstanceAttributesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyInstanceAttributesUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyAttribute", "").
		WithBody(request).
		Do()
}

// ModifyInstanceDesc
//
// PARAMS:
//   - request: the arguments to ModifyInstanceDesc
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceDesc(request *ModifyInstanceDescRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyInstanceDescUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyDesc", "").
		WithBody(request).
		Do()
}

// ModifyInstanceHostname
//
// PARAMS:
//   - request: the arguments to ModifyInstanceHostname
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceHostname(request *ModifyInstanceHostnameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyInstanceHostnameUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("changeHostname", "").
		WithBody(request).
		Do()
}

// ModifyInstancePassword
//
// PARAMS:
//   - request: the arguments to ModifyInstancePassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstancePassword(request *ModifyInstancePasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyInstancePasswordUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("changePass", "").
		WithBody(request).
		Do()
}

// ModifyRelatedDeletePolicy
//
// PARAMS:
//   - request: the arguments to ModifyRelatedDeletePolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyRelatedDeletePolicy(request *ModifyRelatedDeletePolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyRelatedDeletePolicyUri(VERSION_V2, util.StringValue(request.InstanceId))).
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

// PurchaseReservedInstance
//
// PARAMS:
//   - request: the arguments to PurchaseReservedInstance
//
// RETURNS:
//   - PurchaseReservedInstanceResponse: The return type of the PurchaseReservedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedInstance(request *PurchaseReservedInstanceRequest) (*PurchaseReservedInstanceResponse, error) {
	result := &PurchaseReservedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPurchaseReservedInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("relatedRenewFlag", util.StringValue(request.RelatedRenewFlag)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// RebootInstance
//
// PARAMS:
//   - request: the arguments to RebootInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(request *RebootInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRebootInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("reboot", "").
		WithBody(request).
		Do()
}

// RebuildBatchInstance
//
// PARAMS:
//   - request: the arguments to RebuildBatchInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RebuildBatchInstance(request *RebuildBatchInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRebuildBatchInstanceUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// RebuildInstance
//
// PARAMS:
//   - request: the arguments to RebuildInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RebuildInstance(request *RebuildInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRebuildInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("rebuild", "").
		WithBody(request).
		Do()
}

// ReleaseInstanceByPost
//
// PARAMS:
//   - request: the arguments to ReleaseInstanceByPost
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseInstanceByPost(request *ReleaseInstanceByPostRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getReleaseInstanceByPostUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ReleaseMultipleInstanceByPost
//
// PARAMS:
//   - request: the arguments to ReleaseMultipleInstanceByPost
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseMultipleInstanceByPost(request *ReleaseMultipleInstanceByPostRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getReleaseMultipleInstanceByPostUri(VERSION_V2)).
		WithBody(request).
		Do()
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

// ResizeInstanceBySpec
//
// PARAMS:
//   - request: the arguments to ResizeInstanceBySpec
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResizeInstanceBySpec(request *ResizeInstanceBySpecRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeInstanceBySpecUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("resize", "").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
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

// StartInstance
//
// PARAMS:
//   - request: the arguments to StartInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StartInstance(request *StartInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStartInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("start", "").
		Do()
}

// StopInstance
//
// PARAMS:
//   - request: the arguments to StopInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StopInstance(request *StopInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStopInstanceUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("stop", "").
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

// UnbindInstanceFromSecurityGroup
//
// PARAMS:
//   - request: the arguments to UnbindInstanceFromSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindInstanceFromSecurityGroup(request *UnbindInstanceFromSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindInstanceFromSecurityGroupUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("unbind", "").
		WithBody(request).
		Do()
}

// UnbindInstanceFromTags
//
// PARAMS:
//   - request: the arguments to UnbindInstanceFromTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindInstanceFromTags(request *UnbindInstanceFromTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindInstanceFromTagsUri(VERSION_V2, util.StringValue(request.InstanceId))).
		WithQueryParam("unbind", "").
		WithBody(request).
		Do()
}

// UnbindRole
//
// PARAMS:
//   - request: the arguments to UnbindRole
//
// RETURNS:
//   - UnbindRoleResponse: The return type of the UnbindRole interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UnbindRole(request *UnbindRoleRequest) (*UnbindRoleResponse, error) {
	result := &UnbindRoleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUnbindRoleUri(VERSION_V2)).
		WithQueryParam("unbind", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// UpdateInstanceSubnet
//
// PARAMS:
//   - request: the arguments to UpdateInstanceSubnet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceSubnet(request *UpdateInstanceSubnetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceSubnetUri(VERSION_V2)).
		WithBody(request).
		Do()
}
