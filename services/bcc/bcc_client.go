package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V2 = "v2"

	VERSION_V1 = "v1"
)

// AcceptReservedInstanceTransfer
//
// PARAMS:
//   - request: the arguments to AcceptReservedInstanceTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AcceptReservedInstanceTransfer(request *AcceptReservedInstanceTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAcceptReservedInstanceTransferUri(VERSION_V2)).
		WithBody(request).
		Do()
}

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

// AttachAsp
//
// PARAMS:
//   - request: the arguments to AttachAsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AttachAsp(request *AttachAspRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAttachAspUri(VERSION_V2, util.StringValue(request.AspId))).
		WithQueryParam("attach", "").
		WithBody(request).
		Do()
}

// AttachKeypair
//
// PARAMS:
//   - request: the arguments to AttachKeypair
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AttachKeypair(request *AttachKeypairRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAttachKeypairUri(VERSION_V2, util.StringValue(request.KeypairId))).
		WithQueryParam("attach", "").
		WithBody(request).
		Do()
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

// AuthorizeSecurityGroupRule
//
// PARAMS:
//   - request: the arguments to AuthorizeSecurityGroupRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuthorizeSecurityGroupRule(request *AuthorizeSecurityGroupRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAuthorizeSecurityGroupRuleUri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		WithQueryParam("authorizeRule", "").
		WithQueryParamFilter("sgVersion", util.Int32Value(request.SgVersion)).
		WithBody(request).
		Do()
}

// AuthorizeServerEvent
//
// PARAMS:
//   - request: the arguments to AuthorizeServerEvent
//
// RETURNS:
//   - AuthorizeServerEventResponse: The return type of the AuthorizeServerEvent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AuthorizeServerEvent(request *AuthorizeServerEventRequest) (*AuthorizeServerEventResponse, error) {
	result := &AuthorizeServerEventResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAuthorizeServerEventUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
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

// AutoRenewReservedInstance
//
// PARAMS:
//   - request: the arguments to AutoRenewReservedInstance
//
// RETURNS:
//   - AutoRenewReservedInstanceResponse: The return type of the AutoRenewReservedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AutoRenewReservedInstance(request *AutoRenewReservedInstanceRequest) (*AutoRenewReservedInstanceResponse, error) {
	result := &AutoRenewReservedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAutoRenewReservedInstanceUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AutoRenewVolumeCluster
//
// PARAMS:
//   - request: the arguments to AutoRenewVolumeCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AutoRenewVolumeCluster(request *AutoRenewVolumeClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAutoRenewVolumeClusterUri(VERSION_V2)).
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

// BindInstanceSecurityGroup
//
// PARAMS:
//   - request: the arguments to BindInstanceSecurityGroup
//
// RETURNS:
//   - BindInstanceSecurityGroupResponse: The return type of the BindInstanceSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BindInstanceSecurityGroup(request *BindInstanceSecurityGroupRequest) (*BindInstanceSecurityGroupResponse, error) {
	result := &BindInstanceSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindInstanceSecurityGroupUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// BindReservedInstanceToTags
//
// PARAMS:
//   - request: the arguments to BindReservedInstanceToTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindReservedInstanceToTags(request *BindReservedInstanceToTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindReservedInstanceToTagsUri(VERSION_V2)).
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

// BindTagSnapchain
//
// PARAMS:
//   - request: the arguments to BindTagSnapchain
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTagSnapchain(request *BindTagSnapchainRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTagSnapchainUri(VERSION_V2, util.StringValue(request.ChainId))).
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

// BindTagVolumeCluster
//
// PARAMS:
//   - request: the arguments to BindTagVolumeCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTagVolumeCluster(request *BindTagVolumeClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTagVolumeClusterUri(VERSION_V2, util.StringValue(request.ClusterId))).
		WithQueryParam("bind", "").
		WithBody(request).
		Do()
}

// CancelAutoRenewReservedInstance
//
// PARAMS:
//   - request: the arguments to CancelAutoRenewReservedInstance
//
// RETURNS:
//   - CancelAutoRenewReservedInstanceResponse: The return type of the CancelAutoRenewReservedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CancelAutoRenewReservedInstance(request *CancelAutoRenewReservedInstanceRequest) (*CancelAutoRenewReservedInstanceResponse, error) {
	result := &CancelAutoRenewReservedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelAutoRenewReservedInstanceUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CancelAutoRenewVolumeCluster
//
// PARAMS:
//   - request: the arguments to CancelAutoRenewVolumeCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CancelAutoRenewVolumeCluster(request *CancelAutoRenewVolumeClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelAutoRenewVolumeClusterUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// CancelBidOrder
//
// PARAMS:
//   - request: the arguments to CancelBidOrder
//
// RETURNS:
//   - CancelBidOrderResponse: The return type of the CancelBidOrder interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CancelBidOrder(request *CancelBidOrderRequest) (*CancelBidOrderResponse, error) {
	result := &CancelBidOrderResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelBidOrderUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// CancelSnapshotShare
//
// PARAMS:
//   - request: the arguments to CancelSnapshotShare
//
// RETURNS:
//   - CancelSnapshotShareResponse: The return type of the CancelSnapshotShare interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CancelSnapshotShare(request *CancelSnapshotShareRequest) (*CancelSnapshotShareResponse, error) {
	result := &CancelSnapshotShareResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelSnapshotShareUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// CheckServerEvent
//
// PARAMS:
//   - request: the arguments to CheckServerEvent
//
// RETURNS:
//   - CheckServerEventResponse: The return type of the CheckServerEvent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CheckServerEvent(request *CheckServerEventRequest) (*CheckServerEventResponse, error) {
	result := &CheckServerEventResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCheckServerEventUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAsp
//
// PARAMS:
//   - request: the arguments to CreateAsp
//
// RETURNS:
//   - CreateAspResponse: The return type of the CreateAsp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAsp(request *CreateAspRequest) (*CreateAspResponse, error) {
	result := &CreateAspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAspUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAuthorizationRule
//
// PARAMS:
//   - request: the arguments to CreateAuthorizationRule
//
// RETURNS:
//   - CreateAuthorizationRuleResponse: The return type of the CreateAuthorizationRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAuthorizationRule(request *CreateAuthorizationRuleRequest) (*CreateAuthorizationRuleResponse, error) {
	result := &CreateAuthorizationRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAuthorizationRuleUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// CreateBidInstance
//
// PARAMS:
//   - request: the arguments to CreateBidInstance
//
// RETURNS:
//   - CreateBidInstanceResponse: The return type of the CreateBidInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateBidInstance(request *CreateBidInstanceRequest) (*CreateBidInstanceResponse, error) {
	result := &CreateBidInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateBidInstanceUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDeploySet
//
// PARAMS:
//   - request: the arguments to CreateDeploySet
//
// RETURNS:
//   - CreateDeploySetResponse: The return type of the CreateDeploySet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDeploySet(request *CreateDeploySetRequest) (*CreateDeploySetResponse, error) {
	result := &CreateDeploySetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDeploySetUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEhcCluster
//
// PARAMS:
//   - request: the arguments to CreateEhcCluster
//
// RETURNS:
//   - CreateEhcClusterResponse: The return type of the CreateEhcCluster interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEhcCluster(request *CreateEhcClusterRequest) (*CreateEhcClusterResponse, error) {
	result := &CreateEhcClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEhcClusterUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// CreateKeypair
//
// PARAMS:
//   - request: the arguments to CreateKeypair
//
// RETURNS:
//   - CreateKeypairResponse: The return type of the CreateKeypair interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateKeypair(request *CreateKeypairRequest) (*CreateKeypairResponse, error) {
	result := &CreateKeypairResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateKeypairUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateReservedInstanceTransfer
//
// PARAMS:
//   - request: the arguments to CreateReservedInstanceTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateReservedInstanceTransfer(request *CreateReservedInstanceTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateReservedInstanceTransferUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// CreateReservedInstances
//
// PARAMS:
//   - request: the arguments to CreateReservedInstances
//
// RETURNS:
//   - CreateReservedInstancesResponse: The return type of the CreateReservedInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateReservedInstances(request *CreateReservedInstancesRequest) (*CreateReservedInstancesResponse, error) {
	result := &CreateReservedInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateReservedInstancesUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSecurityGroup
//
// PARAMS:
//   - request: the arguments to CreateSecurityGroup
//
// RETURNS:
//   - CreateSecurityGroupResponse: The return type of the CreateSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	result := &CreateSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSecurityGroupUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSnapshot
//
// PARAMS:
//   - request: the arguments to CreateSnapshot
//
// RETURNS:
//   - CreateSnapshotResponse: The return type of the CreateSnapshot interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	result := &CreateSnapshotResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSnapshotUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSnapshotShare
//
// PARAMS:
//   - request: the arguments to CreateSnapshotShare
//
// RETURNS:
//   - CreateSnapshotShareResponse: The return type of the CreateSnapshotShare interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSnapshotShare(request *CreateSnapshotShareRequest) (*CreateSnapshotShareResponse, error) {
	result := &CreateSnapshotShareResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSnapshotShareUri(VERSION_V2)).
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

// CreateVolumeCluster
//
// PARAMS:
//   - request: the arguments to CreateVolumeCluster
//
// RETURNS:
//   - CreateVolumeClusterResponse: The return type of the CreateVolumeCluster interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVolumeCluster(request *CreateVolumeClusterRequest) (*CreateVolumeClusterResponse, error) {
	result := &CreateVolumeClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateVolumeClusterUri(VERSION_V2)).
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

// DeleteAsp
//
// PARAMS:
//   - request: the arguments to DeleteAsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAsp(request *DeleteAspRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAspUri(VERSION_V2, util.StringValue(request.AspId))).
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

// DeleteDeploySet
//
// PARAMS:
//   - request: the arguments to DeleteDeploySet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDeploySet(request *DeleteDeploySetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDeploySetUri(VERSION_V2, util.StringValue(request.DeployId))).
		Do()
}

// DeleteEhcCluster
//
// PARAMS:
//   - request: the arguments to DeleteEhcCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteEhcCluster(request *DeleteEhcClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteEhcClusterUri(VERSION_V2)).
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

// DeleteInstUserOpAuthorizeRule
//
// PARAMS:
//   - request: the arguments to DeleteInstUserOpAuthorizeRule
//
// RETURNS:
//   - DeleteInstUserOpAuthorizeRuleResponse: The return type of the DeleteInstUserOpAuthorizeRule interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstUserOpAuthorizeRule(request *DeleteInstUserOpAuthorizeRuleRequest) (*DeleteInstUserOpAuthorizeRuleResponse, error) {
	result := &DeleteInstUserOpAuthorizeRuleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteInstUserOpAuthorizeRuleUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// DeleteKeypair
//
// PARAMS:
//   - request: the arguments to DeleteKeypair
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteKeypair(request *DeleteKeypairRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteKeypairUri(VERSION_V2, util.StringValue(request.KeypairId))).
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

// DeleteSecurityGroup
//
// PARAMS:
//   - request: the arguments to DeleteSecurityGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSecurityGroupUri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		Do()
}

// DeleteSecurityGroupRule
//
// PARAMS:
//   - request: the arguments to DeleteSecurityGroupRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSecurityGroupRuleUri(VERSION_V2, util.StringValue(request.SecurityGroupRuleId))).
		WithQueryParamFilter("sgVersion", util.Int32Value(request.SgVersion)).
		Do()
}

// DeleteSnapshot
//
// PARAMS:
//   - request: the arguments to DeleteSnapshot
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSnapshotUri(VERSION_V2, util.StringValue(request.SnapshotId))).
		Do()
}

// DeletesInstanceDeploySet
//
// PARAMS:
//   - request: the arguments to DeletesInstanceDeploySet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletesInstanceDeploySet(request *DeletesInstanceDeploySetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeletesInstanceDeploySetUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// DescribeAuthorizeRules
//
// PARAMS:
//   - request: the arguments to DescribeAuthorizeRules
//
// RETURNS:
//   - DescribeAuthorizeRulesResponse: The return type of the DescribeAuthorizeRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeAuthorizeRules(request *DescribeAuthorizeRulesRequest) (*DescribeAuthorizeRulesResponse, error) {
	result := &DescribeAuthorizeRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeAuthorizeRulesUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribePlannedEventRecords
//
// PARAMS:
//   - request: the arguments to DescribePlannedEventRecords
//
// RETURNS:
//   - DescribePlannedEventRecordsResponse: The return type of the DescribePlannedEventRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribePlannedEventRecords(request *DescribePlannedEventRecordsRequest) (*DescribePlannedEventRecordsResponse, error) {
	result := &DescribePlannedEventRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribePlannedEventRecordsUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribePlannedEvents
//
// PARAMS:
//   - request: the arguments to DescribePlannedEvents
//
// RETURNS:
//   - DescribePlannedEventsResponse: The return type of the DescribePlannedEvents interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribePlannedEvents(request *DescribePlannedEventsRequest) (*DescribePlannedEventsResponse, error) {
	result := &DescribePlannedEventsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribePlannedEventsUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeRegions
//
// PARAMS:
//   - request: the arguments to DescribeRegions
//
// RETURNS:
//   - DescribeRegionsResponse: The return type of the DescribeRegions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (*DescribeRegionsResponse, error) {
	result := &DescribeRegionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeRegionsUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeUnplannedEventRecords
//
// PARAMS:
//   - request: the arguments to DescribeUnplannedEventRecords
//
// RETURNS:
//   - DescribeUnplannedEventRecordsResponse: The return type of the DescribeUnplannedEventRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeUnplannedEventRecords(request *DescribeUnplannedEventRecordsRequest) (*DescribeUnplannedEventRecordsResponse, error) {
	result := &DescribeUnplannedEventRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeUnplannedEventRecordsUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeUnplannedEvents
//
// PARAMS:
//   - request: the arguments to DescribeUnplannedEvents
//
// RETURNS:
//   - DescribeUnplannedEventsResponse: The return type of the DescribeUnplannedEvents interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeUnplannedEvents(request *DescribeUnplannedEventsRequest) (*DescribeUnplannedEventsResponse, error) {
	result := &DescribeUnplannedEventsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeUnplannedEventsUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DetachAsp
//
// PARAMS:
//   - request: the arguments to DetachAsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DetachAsp(request *DetachAspRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDetachAspUri(VERSION_V2, util.StringValue(request.AspId))).
		WithQueryParam("detach", "").
		WithBody(request).
		Do()
}

// DetachKeypair
//
// PARAMS:
//   - request: the arguments to DetachKeypair
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DetachKeypair(request *DetachKeypairRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDetachKeypairUri(VERSION_V2, util.StringValue(request.KeypairId))).
		WithQueryParam("detach", "").
		WithBody(request).
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

// EhcClusterList
//
// PARAMS:
//   - request: the arguments to EhcClusterList
//
// RETURNS:
//   - EhcClusterListResponse: The return type of the EhcClusterList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) EhcClusterList(request *EhcClusterListRequest) (*EhcClusterListResponse, error) {
	result := &EhcClusterListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEhcClusterListUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// GetAsp
//
// PARAMS:
//   - request: the arguments to GetAsp
//
// RETURNS:
//   - GetAspResponse: The return type of the GetAsp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAsp(request *GetAspRequest) (*GetAspResponse, error) {
	result := &GetAspResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAspUri(VERSION_V2, util.StringValue(request.AspId))).
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

// GetBidInstancePrice
//
// PARAMS:
//   - request: the arguments to GetBidInstancePrice
//
// RETURNS:
//   - GetBidInstancePriceResponse: The return type of the GetBidInstancePrice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetBidInstancePrice(request *GetBidInstancePriceRequest) (*GetBidInstancePriceResponse, error) {
	result := &GetBidInstancePriceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetBidInstancePriceUri(VERSION_V2)).
		WithBody(request).
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

// GetDeploySet
//
// PARAMS:
//   - request: the arguments to GetDeploySet
//
// RETURNS:
//   - GetDeploySetResponse: The return type of the GetDeploySet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDeploySet(request *GetDeploySetRequest) (*GetDeploySetResponse, error) {
	result := &GetDeploySetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDeploySetUri(VERSION_V2, util.StringValue(request.Id))).
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

// GetPriceBySpec
//
// PARAMS:
//   - request: the arguments to GetPriceBySpec
//
// RETURNS:
//   - GetPriceBySpecResponse: The return type of the GetPriceBySpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPriceBySpec(request *GetPriceBySpecRequest) (*GetPriceBySpecResponse, error) {
	result := &GetPriceBySpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetPriceBySpecUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetReservedInstance
//
// PARAMS:
//   - request: the arguments to GetReservedInstance
//
// RETURNS:
//   - GetReservedInstanceResponse: The return type of the GetReservedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetReservedInstance(request *GetReservedInstanceRequest) (*GetReservedInstanceResponse, error) {
	result := &GetReservedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetReservedInstanceUri(VERSION_V2)).
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

// GetReservedInstancePrice
//
// PARAMS:
//   - request: the arguments to GetReservedInstancePrice
//
// RETURNS:
//   - GetReservedInstancePriceResponse: The return type of the GetReservedInstancePrice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetReservedInstancePrice(request *GetReservedInstancePriceRequest) (*GetReservedInstancePriceResponse, error) {
	result := &GetReservedInstancePriceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetReservedInstancePriceUri(VERSION_V2)).
		WithBody(request).
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

// GetSnapshot
//
// PARAMS:
//   - request: the arguments to GetSnapshot
//
// RETURNS:
//   - GetSnapshotResponse: The return type of the GetSnapshot interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSnapshot(request *GetSnapshotRequest) (*GetSnapshotResponse, error) {
	result := &GetSnapshotResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSnapshotUri(VERSION_V2, util.StringValue(request.SnapshotId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTask
//
// PARAMS:
//   - request: the arguments to GetTask
//
// RETURNS:
//   - GetTaskResponse: The return type of the GetTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTask(request *GetTaskRequest) (*GetTaskResponse, error) {
	result := &GetTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetTaskUri(VERSION_V2)).
		WithBody(request).
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

// GetVolumeCluster
//
// PARAMS:
//   - request: the arguments to GetVolumeCluster
//
// RETURNS:
//   - GetVolumeClusterResponse: The return type of the GetVolumeCluster interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetVolumeCluster(request *GetVolumeClusterRequest) (*GetVolumeClusterResponse, error) {
	result := &GetVolumeClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetVolumeClusterUri(VERSION_V2, util.StringValue(request.ClusterId))).
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

// GetZoneBySpec
//
// PARAMS:
//   - request: the arguments to GetZoneBySpec
//
// RETURNS:
//   - GetZoneBySpecResponse: The return type of the GetZoneBySpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetZoneBySpec(request *GetZoneBySpecRequest) (*GetZoneBySpecResponse, error) {
	result := &GetZoneBySpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetZoneBySpecUri(VERSION_V1)).
		WithQueryParamFilter("instanceType", util.StringValue(request.InstanceType)).
		WithQueryParamFilter("productType", util.StringValue(request.ProductType)).
		WithQueryParamFilter("spec", util.StringValue(request.Spec)).
		WithQueryParamFilter("specId", util.StringValue(request.SpecId)).
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

// ImportKeypair
//
// PARAMS:
//   - request: the arguments to ImportKeypair
//
// RETURNS:
//   - ImportKeypairResponse: The return type of the ImportKeypair interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImportKeypair(request *ImportKeypairRequest) (*ImportKeypairResponse, error) {
	result := &ImportKeypairResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getImportKeypairUri(VERSION_V2)).
		WithQueryParam("import", "").
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

// KeypairDetail
//
// PARAMS:
//   - request: the arguments to KeypairDetail
//
// RETURNS:
//   - KeypairDetailResponse: The return type of the KeypairDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) KeypairDetail(request *KeypairDetailRequest) (*KeypairDetailResponse, error) {
	result := &KeypairDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getKeypairDetailUri(VERSION_V2, util.StringValue(request.KeypairId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAsps
//
// PARAMS:
//   - request: the arguments to ListAsps
//
// RETURNS:
//   - ListAspsResponse: The return type of the ListAsps interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAsps(request *ListAspsRequest) (*ListAspsResponse, error) {
	result := &ListAspsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAspsUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("aspName", util.StringValue(request.AspName)).
		WithQueryParamFilter("volumeName", util.StringValue(request.VolumeName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// ListBidFlavor
//
// PARAMS:
//   - request: the arguments to ListBidFlavor
//
// RETURNS:
//   - ListBidFlavorResponse: The return type of the ListBidFlavor interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListBidFlavor() (*ListBidFlavorResponse, error) {
	result := &ListBidFlavorResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListBidFlavorUri(VERSION_V2)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListDeploySet
//
// PARAMS:
//   - request: the arguments to ListDeploySet
//
// RETURNS:
//   - ListDeploySetResponse: The return type of the ListDeploySet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListDeploySet() (*ListDeploySetResponse, error) {
	result := &ListDeploySetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListDeploySetUri(VERSION_V2)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListFlavorSpec
//
// PARAMS:
//   - request: the arguments to ListFlavorSpec
//
// RETURNS:
//   - ListFlavorSpecResponse: The return type of the ListFlavorSpec interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListFlavorSpec(request *ListFlavorSpecRequest) (*ListFlavorSpecResponse, error) {
	result := &ListFlavorSpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListFlavorSpecUri(VERSION_V2)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("specs", util.StringValue(request.Specs)).
		WithQueryParamFilter("specIds", util.StringValue(request.SpecIds)).
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

// ListKeypair
//
// PARAMS:
//   - request: the arguments to ListKeypair
//
// RETURNS:
//   - ListKeypairResponse: The return type of the ListKeypair interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListKeypair(request *ListKeypairRequest) (*ListKeypairResponse, error) {
	result := &ListKeypairResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListKeypairUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
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

// ListReservedInstanceTransferIn
//
// PARAMS:
//   - request: the arguments to ListReservedInstanceTransferIn
//
// RETURNS:
//   - ListReservedInstanceTransferInResponse: The return type of the ListReservedInstanceTransferIn interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListReservedInstanceTransferIn(request *ListReservedInstanceTransferInRequest) (*ListReservedInstanceTransferInResponse, error) {
	result := &ListReservedInstanceTransferInResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListReservedInstanceTransferInUri(VERSION_V2)).
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

// ListReservedInstanceTransferOut
//
// PARAMS:
//   - request: the arguments to ListReservedInstanceTransferOut
//
// RETURNS:
//   - ListReservedInstanceTransferOutResponse: The return type of the ListReservedInstanceTransferOut interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListReservedInstanceTransferOut(request *ListReservedInstanceTransferOutRequest) (*ListReservedInstanceTransferOutResponse, error) {
	result := &ListReservedInstanceTransferOutResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListReservedInstanceTransferOutUri(VERSION_V2)).
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

// ListSecurityGroups
//
// PARAMS:
//   - request: the arguments to ListSecurityGroups
//
// RETURNS:
//   - ListSecurityGroupsResponse: The return type of the ListSecurityGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroups(request *ListSecurityGroupsRequest) (*ListSecurityGroupsResponse, error) {
	result := &ListSecurityGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSecurityGroupsUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("securityGroupId", util.StringValue(request.SecurityGroupId)).
		WithQueryParamFilter("securityGroupIds", util.StringValue(request.SecurityGroupIds)).
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

// ListSnapchain
//
// PARAMS:
//   - request: the arguments to ListSnapchain
//
// RETURNS:
//   - ListSnapchainResponse: The return type of the ListSnapchain interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSnapchain(request *ListSnapchainRequest) (*ListSnapchainResponse, error) {
	result := &ListSnapchainResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSnapchainUri(VERSION_V2)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("volumeId", util.StringValue(request.VolumeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSnapshotShare
//
// PARAMS:
//   - request: the arguments to ListSnapshotShare
//
// RETURNS:
//   - ListSnapshotShareResponse: The return type of the ListSnapshotShare interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSnapshotShare(request *ListSnapshotShareRequest) (*ListSnapshotShareResponse, error) {
	result := &ListSnapshotShareResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListSnapshotShareUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListSnapshots
//
// PARAMS:
//   - request: the arguments to ListSnapshots
//
// RETURNS:
//   - ListSnapshotsResponse: The return type of the ListSnapshots interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListSnapshots(request *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	result := &ListSnapshotsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListSnapshotsUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("volumeId", util.StringValue(request.VolumeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTask
//
// PARAMS:
//   - request: the arguments to ListTask
//
// RETURNS:
//   - ListTaskResponse: The return type of the ListTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTask(request *ListTaskRequest) (*ListTaskResponse, error) {
	result := &ListTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListTaskUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListVolumeClusters
//
// PARAMS:
//   - request: the arguments to ListVolumeClusters
//
// RETURNS:
//   - ListVolumeClustersResponse: The return type of the ListVolumeClusters interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListVolumeClusters(request *ListVolumeClustersRequest) (*ListVolumeClustersResponse, error) {
	result := &ListVolumeClustersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListVolumeClustersUri(VERSION_V2)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("clusterName", util.StringValue(request.ClusterName)).
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

// ListZones
//
// PARAMS:
//   - request: the arguments to ListZones
//
// RETURNS:
//   - ListZonesResponse: The return type of the ListZones interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListZones() (*ListZonesResponse, error) {
	result := &ListZonesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListZonesUri(VERSION_V2)).
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

// ModifyEhcCluster
//
// PARAMS:
//   - request: the arguments to ModifyEhcCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyEhcCluster(request *ModifyEhcClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyEhcClusterUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// ModifyInstUserOpAuthorizeRuleAttribute
//
// PARAMS:
//   - request: the arguments to ModifyInstUserOpAuthorizeRuleAttribute
//
// RETURNS:
//   - ModifyInstUserOpAuthorizeRuleAttributeResponse: The return type of the ModifyInstUserOpAuthorizeRuleAttribute interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyInstUserOpAuthorizeRuleAttribute(request *ModifyInstUserOpAuthorizeRuleAttributeRequest) (*ModifyInstUserOpAuthorizeRuleAttributeResponse, error) {
	result := &ModifyInstUserOpAuthorizeRuleAttributeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyInstUserOpAuthorizeRuleAttributeUri()).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// ModifyReservedInstances
//
// PARAMS:
//   - request: the arguments to ModifyReservedInstances
//
// RETURNS:
//   - ModifyReservedInstancesResponse: The return type of the ModifyReservedInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ModifyReservedInstances(request *ModifyReservedInstancesRequest) (*ModifyReservedInstancesResponse, error) {
	result := &ModifyReservedInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyReservedInstancesUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// ModifyVolumeDeleteProtectionV2
//
// PARAMS:
//   - request: the arguments to ModifyVolumeDeleteProtectionV2
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyVolumeDeleteProtectionV2(request *ModifyVolumeDeleteProtectionV2Request) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyVolumeDeleteProtectionV2Uri(VERSION_V2)).
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

// PurchaseReservedVolumeCluster
//
// PARAMS:
//   - request: the arguments to PurchaseReservedVolumeCluster
//
// RETURNS:
//   - PurchaseReservedVolumeClusterResponse: The return type of the PurchaseReservedVolumeCluster interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedVolumeCluster(request *PurchaseReservedVolumeClusterRequest) (*PurchaseReservedVolumeClusterResponse, error) {
	result := &PurchaseReservedVolumeClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPurchaseReservedVolumeClusterUri(VERSION_V2, util.StringValue(request.ClusterId))).
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

// RefuseReservedInstanceTransfer
//
// PARAMS:
//   - request: the arguments to RefuseReservedInstanceTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RefuseReservedInstanceTransfer(request *RefuseReservedInstanceTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRefuseReservedInstanceTransferUri(VERSION_V2)).
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

// RemoteCopySnapshot
//
// PARAMS:
//   - request: the arguments to RemoteCopySnapshot
//
// RETURNS:
//   - RemoteCopySnapshotResponse: The return type of the RemoteCopySnapshot interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RemoteCopySnapshot(request *RemoteCopySnapshotRequest) (*RemoteCopySnapshotResponse, error) {
	result := &RemoteCopySnapshotResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRemoteCopySnapshotUri(VERSION_V2, util.StringValue(request.SnapshotId))).
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

// RenameKeypair
//
// PARAMS:
//   - request: the arguments to RenameKeypair
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RenameKeypair(request *RenameKeypairRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRenameKeypairUri(VERSION_V2, util.StringValue(request.KeypairId))).
		WithQueryParam("rename", "").
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

// RenewReservedInstance
//
// PARAMS:
//   - request: the arguments to RenewReservedInstance
//
// RETURNS:
//   - RenewReservedInstanceResponse: The return type of the RenewReservedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RenewReservedInstance(request *RenewReservedInstanceRequest) (*RenewReservedInstanceResponse, error) {
	result := &RenewReservedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRenewReservedInstanceUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReplaceInstanceSecurityGroup
//
// PARAMS:
//   - request: the arguments to ReplaceInstanceSecurityGroup
//
// RETURNS:
//   - ReplaceInstanceSecurityGroupResponse: The return type of the ReplaceInstanceSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ReplaceInstanceSecurityGroup(request *ReplaceInstanceSecurityGroupRequest) (*ReplaceInstanceSecurityGroupResponse, error) {
	result := &ReplaceInstanceSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getReplaceInstanceSecurityGroupUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// ResizeVolumeCluster
//
// PARAMS:
//   - request: the arguments to ResizeVolumeCluster
//
// RETURNS:
//   - ResizeVolumeClusterResponse: The return type of the ResizeVolumeCluster interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeVolumeCluster(request *ResizeVolumeClusterRequest) (*ResizeVolumeClusterResponse, error) {
	result := &ResizeVolumeClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResizeVolumeClusterUri(VERSION_V2, util.StringValue(request.ClusterId))).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RevokeReservedInstanceTransfer
//
// PARAMS:
//   - request: the arguments to RevokeReservedInstanceTransfer
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RevokeReservedInstanceTransfer(request *RevokeReservedInstanceTransferRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRevokeReservedInstanceTransferUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// RevokeSecurityGroupRule
//
// PARAMS:
//   - request: the arguments to RevokeSecurityGroupRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RevokeSecurityGroupRule(request *RevokeSecurityGroupRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRevokeSecurityGroupRuleUri(VERSION_V2, util.StringValue(request.SecurityGroupId))).
		WithQueryParam("revokeRule", "").
		WithQueryParamFilter("sgVersion", util.Int32Value(request.SgVersion)).
		WithBody(request).
		Do()
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

// UnbindInstanceSecurityGroup
//
// PARAMS:
//   - request: the arguments to UnbindInstanceSecurityGroup
//
// RETURNS:
//   - UnbindInstanceSecurityGroupResponse: The return type of the UnbindInstanceSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UnbindInstanceSecurityGroup(request *UnbindInstanceSecurityGroupRequest) (*UnbindInstanceSecurityGroupResponse, error) {
	result := &UnbindInstanceSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindInstanceSecurityGroupUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindReservedInstanceFromTags
//
// PARAMS:
//   - request: the arguments to UnbindReservedInstanceFromTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindReservedInstanceFromTags(request *UnbindReservedInstanceFromTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindReservedInstanceFromTagsUri(VERSION_V2)).
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

// UnbindTagSnapchain
//
// PARAMS:
//   - request: the arguments to UnbindTagSnapchain
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTagSnapchain(request *UnbindTagSnapchainRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTagSnapchainUri(VERSION_V2, util.StringValue(request.ChainId))).
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

// UnbindTagVolumeCluster
//
// PARAMS:
//   - request: the arguments to UnbindTagVolumeCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTagVolumeCluster(request *UnbindTagVolumeClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTagVolumeClusterUri(VERSION_V2, util.StringValue(request.ClusterId))).
		WithQueryParam("unbind", "").
		WithBody(request).
		Do()
}

// UpdateAsp
//
// PARAMS:
//   - request: the arguments to UpdateAsp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAsp(request *UpdateAspRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAspUri(VERSION_V2)).
		WithBody(request).
		Do()
}

// UpdateDeploySet
//
// PARAMS:
//   - request: the arguments to UpdateDeploySet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateDeploySet(request *UpdateDeploySetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateDeploySetUri(VERSION_V2, util.StringValue(request.DeployId))).
		WithQueryParam("modifyAttribute", "").
		WithBody(request).
		Do()
}

// UpdateDeploySetRelation
//
// PARAMS:
//   - request: the arguments to UpdateDeploySetRelation
//
// RETURNS:
//   - UpdateDeploySetRelationResponse: The return type of the UpdateDeploySetRelation interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateDeploySetRelation(request *UpdateDeploySetRelationRequest) (*UpdateDeploySetRelationResponse, error) {
	result := &UpdateDeploySetRelationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateDeploySetRelationUri(VERSION_V2)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
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

// UpdateKeypairDescription
//
// PARAMS:
//   - request: the arguments to UpdateKeypairDescription
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateKeypairDescription(request *UpdateKeypairDescriptionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateKeypairDescriptionUri(VERSION_V2, util.StringValue(request.KeypairId))).
		WithQueryParam("updateDesc", "").
		WithBody(request).
		Do()
}

// UpdateSecurityGroupRule
//
// PARAMS:
//   - request: the arguments to UpdateSecurityGroupRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityGroupRule(request *UpdateSecurityGroupRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateSecurityGroupRuleUri(VERSION_V2)).
		WithQueryParamFilter("sgVersion", util.Int32Value(request.SgVersion)).
		WithBody(request).
		Do()
}
