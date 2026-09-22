package scs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AccountList
//
// PARAMS:
//   - request: the arguments to AccountList
//
// RETURNS:
//   - AccountListResponse: The return type of the AccountList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AccountList(request *AccountListRequest) (*AccountListResponse, error) {
	result := &AccountListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAccountListUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddIpWhitelist
//
// PARAMS:
//   - request: the arguments to AddIpWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddIpWhitelist(request *AddIpWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAddIpWhitelistUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// AddParametersToParameterTemplate
//
// PARAMS:
//   - request: the arguments to AddParametersToParameterTemplate
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddParametersToParameterTemplate(request *AddParametersToParameterTemplateRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddParametersToParameterTemplateUri(util.StringValue(request.TemplateShowId))).
		WithBody(request).
		Do()
}

// ApplicationParameterTemplate
//
// PARAMS:
//   - request: the arguments to ApplicationParameterTemplate
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ApplicationParameterTemplate(request *ApplicationParameterTemplateRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApplicationParameterTemplateUri(util.StringValue(request.TemplateShowId))).
		WithBody(request).
		Do()
}

// AuditLogSwitch
//
// PARAMS:
//   - request: the arguments to AuditLogSwitch
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuditLogSwitch(request *AuditLogSwitchRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAuditLogSwitchUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// BatchRestoreInstances
//
// PARAMS:
//   - request: the arguments to BatchRestoreInstances
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BatchRestoreInstances(request *BatchRestoreInstancesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchRestoreInstancesUri()).
		WithBody(request).
		Do()
}

// BindSecurityGroup
//
// PARAMS:
//   - request: the arguments to BindSecurityGroup
//
// RETURNS:
//   - BindSecurityGroupResponse: The return type of the BindSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BindSecurityGroup(request *BindSecurityGroupRequest) (*BindSecurityGroupResponse, error) {
	result := &BindSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBindSecurityGroupUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindTags
//
// PARAMS:
//   - request: the arguments to BindTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindTags(request *BindTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBindTagsUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CancelPrepaidToPostpaid
//
// PARAMS:
//   - request: the arguments to CancelPrepaidToPostpaid
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CancelPrepaidToPostpaid(request *CancelPrepaidToPostpaidRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelPrepaidToPostpaidUri()).
		WithBody(request).
		Do()
}

// ChangeAccessPassword
//
// PARAMS:
//   - request: the arguments to ChangeAccessPassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ChangeAccessPassword(request *ChangeAccessPasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getChangeAccessPasswordUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ChangeAccountPassword
//
// PARAMS:
//   - request: the arguments to ChangeAccountPassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ChangeAccountPassword(request *ChangeAccountPasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getChangeAccountPasswordUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ChangeConfiguration
//
// PARAMS:
//   - request: the arguments to ChangeConfiguration
//
// RETURNS:
//   - ChangeConfigurationResponse: The return type of the ChangeConfiguration interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ChangeConfiguration(request *ChangeConfigurationRequest) (*ChangeConfigurationResponse, error) {
	result := &ChangeConfigurationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getChangeConfigurationUri(util.StringValue(request.InstanceId), util.StringValue(request.ClientToken))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ClearInstance
//
// PARAMS:
//   - request: the arguments to ClearInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ClearInstance(request *ClearInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClearInstanceUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ClusterStatusCheck
//
// PARAMS:
//   - request: the arguments to ClusterStatusCheck
//
// RETURNS:
//   - ClusterStatusCheckResponse: The return type of the ClusterStatusCheck interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ClusterStatusCheck(request *ClusterStatusCheckRequest) (*ClusterStatusCheckResponse, error) {
	result := &ClusterStatusCheckResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterStatusCheckUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ClusterTypeUpgrade
//
// PARAMS:
//   - request: the arguments to ClusterTypeUpgrade
//
// RETURNS:
//   - ClusterTypeUpgradeResponse: The return type of the ClusterTypeUpgrade interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ClusterTypeUpgrade(request *ClusterTypeUpgradeRequest) (*ClusterTypeUpgradeResponse, error) {
	result := &ClusterTypeUpgradeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClusterTypeUpgradeUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAccount
//
// PARAMS:
//   - request: the arguments to CreateAccount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAccount(request *CreateAccountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAccountUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateAnInstance
//
// PARAMS:
//   - request: the arguments to CreateAnInstance
//
// RETURNS:
//   - CreateAnInstanceResponse: The return type of the CreateAnInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAnInstance(request *CreateAnInstanceRequest) (*CreateAnInstanceResponse, error) {
	result := &CreateAnInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAnInstanceUri()).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateDeploymentSet
//
// PARAMS:
//   - request: the arguments to CreateDeploymentSet
//
// RETURNS:
//   - CreateDeploymentSetResponse: The return type of the CreateDeploymentSet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateDeploymentSet(request *CreateDeploymentSetRequest) (*CreateDeploymentSetResponse, error) {
	result := &CreateDeploymentSetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateDeploymentSetUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateEntrance
//
// PARAMS:
//   - request: the arguments to CreateEntrance
//
// RETURNS:
//   - CreateEntranceResponse: The return type of the CreateEntrance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEntrance(request *CreateEntranceRequest) (*CreateEntranceResponse, error) {
	result := &CreateEntranceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateEntranceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateHotGroup
//
// PARAMS:
//   - request: the arguments to CreateHotGroup
//
// RETURNS:
//   - CreateHotGroupResponse: The return type of the CreateHotGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateHotGroup(request *CreateHotGroupRequest) (*CreateHotGroupResponse, error) {
	result := &CreateHotGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateHotGroupUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateInstanceWhiteGroup
//
// PARAMS:
//   - request: the arguments to CreateInstanceWhiteGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceWhiteGroup(request *CreateInstanceWhiteGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceWhiteGroupUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateParameterTemplate
//
// PARAMS:
//   - request: the arguments to CreateParameterTemplate
//
// RETURNS:
//   - CreateParameterTemplateResponse: The return type of the CreateParameterTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateParameterTemplate(request *CreateParameterTemplateRequest) (*CreateParameterTemplateResponse, error) {
	result := &CreateParameterTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateParameterTemplateUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateSyncGroup
//
// PARAMS:
//   - request: the arguments to CreateSyncGroup
//
// RETURNS:
//   - CreateSyncGroupResponse: The return type of the CreateSyncGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSyncGroup(request *CreateSyncGroupRequest) (*CreateSyncGroupResponse, error) {
	result := &CreateSyncGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateSyncGroupUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAccount
//
// PARAMS:
//   - request: the arguments to DeleteAccount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAccount(request *DeleteAccountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAccountUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeleteDeploymentSet
//
// PARAMS:
//   - request: the arguments to DeleteDeploymentSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteDeploymentSet(request *DeleteDeploymentSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteDeploymentSetUri(util.StringValue(request.DeploySetId))).
		Do()
}

// DeleteInstanceWhiteGroup
//
// PARAMS:
//   - request: the arguments to DeleteInstanceWhiteGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceWhiteGroup(request *DeleteInstanceWhiteGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteInstanceWhiteGroupUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("groupName", util.StringValue(request.GroupName)).
		Do()
}

// DeleteInstances
//
// PARAMS:
//   - request: the arguments to DeleteInstances
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstances(request *DeleteInstancesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteInstancesUri()).
		WithBody(request).
		Do()
}

// DeleteIpWhitelist
//
// PARAMS:
//   - request: the arguments to DeleteIpWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIpWhitelist(request *DeleteIpWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteIpWhitelistUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeleteManualBackup
//
// PARAMS:
//   - request: the arguments to DeleteManualBackup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteManualBackup(request *DeleteManualBackupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteManualBackupUri(util.StringValue(request.InstanceId), util.StringValue(request.BatchId))).
		Do()
}

// DeleteMemoryScalingConfig
//
// PARAMS:
//   - request: the arguments to DeleteMemoryScalingConfig
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteMemoryScalingConfig(request *DeleteMemoryScalingConfigRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteMemoryScalingConfigUri(util.StringValue(request.InstanceId))).
		Do()
}

// DeleteParameterTemplate
//
// PARAMS:
//   - request: the arguments to DeleteParameterTemplate
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteParameterTemplate(request *DeleteParameterTemplateRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteParameterTemplateUri(util.StringValue(request.TemplateShowId))).
		Do()
}

// DeleteSyncGroup
//
// PARAMS:
//   - request: the arguments to DeleteSyncGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSyncGroup(request *DeleteSyncGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteSyncGroupUri(VERSION_V1, util.StringValue(request.SyncGroupShowId))).
		Do()
}

// DisconnectEntrance
//
// PARAMS:
//   - request: the arguments to DisconnectEntrance
//
// RETURNS:
//   - DisconnectEntranceResponse: The return type of the DisconnectEntrance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DisconnectEntrance(request *DisconnectEntranceRequest) (*DisconnectEntranceResponse, error) {
	result := &DisconnectEntranceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDisconnectEntranceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DomainNameExchange
//
// PARAMS:
//   - request: the arguments to DomainNameExchange
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DomainNameExchange(request *DomainNameExchangeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDomainNameExchangeUri()).
		WithBody(request).
		Do()
}

// GePriceForResizeInstance
//
// PARAMS:
//   - request: the arguments to GePriceForResizeInstance
//
// RETURNS:
//   - GePriceForResizeInstanceResponse: The return type of the GePriceForResizeInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GePriceForResizeInstance(request *GePriceForResizeInstanceRequest) (*GePriceForResizeInstanceResponse, error) {
	result := &GePriceForResizeInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGePriceForResizeInstanceUri(util.StringValue(request.ClientToken), util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetApplicationParameterTemplateRecords
//
// PARAMS:
//   - request: the arguments to GetApplicationParameterTemplateRecords
//
// RETURNS:
//   - GetApplicationParameterTemplateRecordsResponse: The return type of the GetApplicationParameterTemplateRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetApplicationParameterTemplateRecords(request *GetApplicationParameterTemplateRecordsRequest) (*GetApplicationParameterTemplateRecordsResponse, error) {
	result := &GetApplicationParameterTemplateRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetApplicationParameterTemplateRecordsUri(util.StringValue(request.TemplateShowId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAvailableZones
//
// PARAMS:
//   - request: the arguments to GetAvailableZones
//
// RETURNS:
//   - GetAvailableZonesResponse: The return type of the GetAvailableZones interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAvailableZones() (*GetAvailableZonesResponse, error) {
	result := &GetAvailableZonesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAvailableZonesUri()).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBackUpUrl
//
// PARAMS:
//   - request: the arguments to GetBackUpUrl
//
// RETURNS:
//   - GetBackUpUrlResponse: The return type of the GetBackUpUrl interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackUpUrl(request *GetBackUpUrlRequest) (*GetBackUpUrlResponse, error) {
	result := &GetBackUpUrlResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetBackUpUrlUri(util.StringValue(request.InstanceId), util.Int32Value(request.BackupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBackUpUsage
//
// PARAMS:
//   - request: the arguments to GetBackUpUsage
//
// RETURNS:
//   - GetBackUpUsageResponse: The return type of the GetBackUpUsage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackUpUsage(request *GetBackUpUsageRequest) (*GetBackUpUsageResponse, error) {
	result := &GetBackUpUsageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetBackUpUsageUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBackupList
//
// PARAMS:
//   - request: the arguments to GetBackupList
//
// RETURNS:
//   - GetBackupListResponse: The return type of the GetBackupList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupList(request *GetBackupListRequest) (*GetBackupListResponse, error) {
	result := &GetBackupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetBackupListUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBackupStrategy
//
// PARAMS:
//   - request: the arguments to GetBackupStrategy
//
// RETURNS:
//   - GetBackupStrategyResponse: The return type of the GetBackupStrategy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetBackupStrategy(request *GetBackupStrategyRequest) (*GetBackupStrategyResponse, error) {
	result := &GetBackupStrategyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetBackupStrategyUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetClusterBlbStatus
//
// PARAMS:
//   - request: the arguments to GetClusterBlbStatus
//
// RETURNS:
//   - GetClusterBlbStatusResponse: The return type of the GetClusterBlbStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetClusterBlbStatus(request *GetClusterBlbStatusRequest) (*GetClusterBlbStatusResponse, error) {
	result := &GetClusterBlbStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetClusterBlbStatusUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDeploymentSetList
//
// PARAMS:
//   - request: the arguments to GetDeploymentSetList
//
// RETURNS:
//   - GetDeploymentSetListResponse: The return type of the GetDeploymentSetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetDeploymentSetList(request *GetDeploymentSetListRequest) (*GetDeploymentSetListResponse, error) {
	result := &GetDeploymentSetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetDeploymentSetListUri()).
		WithQueryParamFilter("maxKeys", "1").
		WithQueryParamFilter("marker", "-1").
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHotGroupDetail
//
// PARAMS:
//   - request: the arguments to GetHotGroupDetail
//
// RETURNS:
//   - GetHotGroupDetailResponse: The return type of the GetHotGroupDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetHotGroupDetail(request *GetHotGroupDetailRequest) (*GetHotGroupDetailResponse, error) {
	result := &GetHotGroupDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetHotGroupDetailUri(util.StringValue(request.GroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHotGroupList
//
// PARAMS:
//   - request: the arguments to GetHotGroupList
//
// RETURNS:
//   - GetHotGroupListResponse: The return type of the GetHotGroupList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetHotGroupList(request *GetHotGroupListRequest) (*GetHotGroupListResponse, error) {
	result := &GetHotGroupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetHotGroupListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceDetail
//
// PARAMS:
//   - request: the arguments to GetInstanceDetail
//
// RETURNS:
//   - GetInstanceDetailResponse: The return type of the GetInstanceDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(request *GetInstanceDetailRequest) (*GetInstanceDetailResponse, error) {
	result := &GetInstanceDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceDetailUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceList
//
// PARAMS:
//   - request: the arguments to GetInstanceList
//
// RETURNS:
//   - GetInstanceListResponse: The return type of the GetInstanceList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceList(request *GetInstanceListRequest) (*GetInstanceListResponse, error) {
	result := &GetInstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceListUri(util.StringValue(request.Marker), util.StringValue(request.MaxKeys), util.StringValue(request.InstanceIds), util.StringValue(request.VnetIp))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceSpecList
//
// PARAMS:
//   - request: the arguments to GetInstanceSpecList
//
// RETURNS:
//   - GetInstanceSpecListResponse: The return type of the GetInstanceSpecList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceSpecList() (*GetInstanceSpecListResponse, error) {
	result := &GetInstanceSpecListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceSpecListUri()).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceWhiteGroup
//
// PARAMS:
//   - request: the arguments to GetInstanceWhiteGroup
//
// RETURNS:
//   - GetInstanceWhiteGroupResponse: The return type of the GetInstanceWhiteGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceWhiteGroup(request *GetInstanceWhiteGroupRequest) (*GetInstanceWhiteGroupResponse, error) {
	result := &GetInstanceWhiteGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceWhiteGroupUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("groupName", util.StringValue(request.GroupName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetParameterList
//
// PARAMS:
//   - request: the arguments to GetParameterList
//
// RETURNS:
//   - GetParameterListResponse: The return type of the GetParameterList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetParameterList(request *GetParameterListRequest) (*GetParameterListResponse, error) {
	result := &GetParameterListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetParameterListUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetParameterTemplateList
//
// PARAMS:
//   - request: the arguments to GetParameterTemplateList
//
// RETURNS:
//   - GetParameterTemplateListResponse: The return type of the GetParameterTemplateList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetParameterTemplateList(request *GetParameterTemplateListRequest) (*GetParameterTemplateListResponse, error) {
	result := &GetParameterTemplateListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetParameterTemplateListUri()).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPriceForCreateInstance
//
// PARAMS:
//   - request: the arguments to GetPriceForCreateInstance
//
// RETURNS:
//   - GetPriceForCreateInstanceResponse: The return type of the GetPriceForCreateInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPriceForCreateInstance(request *GetPriceForCreateInstanceRequest) (*GetPriceForCreateInstanceResponse, error) {
	result := &GetPriceForCreateInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetPriceForCreateInstanceUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRecycleList
//
// PARAMS:
//   - request: the arguments to GetRecycleList
//
// RETURNS:
//   - GetRecycleListResponse: The return type of the GetRecycleList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRecycleList(request *GetRecycleListRequest) (*GetRecycleListResponse, error) {
	result := &GetRecycleListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRecycleListUri()).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSubnetList
//
// PARAMS:
//   - request: the arguments to GetSubnetList
//
// RETURNS:
//   - GetSubnetListResponse: The return type of the GetSubnetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSubnetList(request *GetSubnetListRequest) (*GetSubnetListResponse, error) {
	result := &GetSubnetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSubnetListUri()).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSyncGroupStatus
//
// PARAMS:
//   - request: the arguments to GetSyncGroupStatus
//
// RETURNS:
//   - GetSyncGroupStatusResponse: The return type of the GetSyncGroupStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSyncGroupStatus(request *GetSyncGroupStatusRequest) (*GetSyncGroupStatusResponse, error) {
	result := &GetSyncGroupStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSyncGroupStatusUri(VERSION_V1, util.StringValue(request.SyncGroupShowId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSystemParameterList
//
// PARAMS:
//   - request: the arguments to GetSystemParameterList
//
// RETURNS:
//   - GetSystemParameterListResponse: The return type of the GetSystemParameterList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSystemParameterList(request *GetSystemParameterListRequest) (*GetSystemParameterListResponse, error) {
	result := &GetSystemParameterListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSystemParameterListUri()).
		WithQueryParamFilter("engine", util.StringValue(request.Engine)).
		WithQueryParamFilter("engineVersion", util.StringValue(request.EngineVersion)).
		WithQueryParamFilter("clusterType", util.StringValue(request.ClusterType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTimeWindow
//
// PARAMS:
//   - request: the arguments to GetTimeWindow
//
// RETURNS:
//   - GetTimeWindowResponse: The return type of the GetTimeWindow interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTimeWindow(request *GetTimeWindowRequest) (*GetTimeWindowResponse, error) {
	result := &GetTimeWindowResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTimeWindowUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTlsCert
//
// PARAMS:
//   - request: the arguments to GetTlsCert
//
// RETURNS:
//   - GetTlsCertResponse: The return type of the GetTlsCert interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTlsCert(request *GetTlsCertRequest) (*GetTlsCertResponse, error) {
	result := &GetTlsCertResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTlsCertUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HotGroupAddCluster
//
// PARAMS:
//   - request: the arguments to HotGroupAddCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupAddCluster(request *HotGroupAddClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHotGroupAddClusterUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// HotGroupChangeMasterRole
//
// PARAMS:
//   - request: the arguments to HotGroupChangeMasterRole
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupChangeMasterRole(request *HotGroupChangeMasterRoleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHotGroupChangeMasterRoleUri(util.StringValue(request.GroupId), util.StringValue(request.InstanceId))).
		Do()
}

// HotGroupForbidWrite
//
// PARAMS:
//   - request: the arguments to HotGroupForbidWrite
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupForbidWrite(request *HotGroupForbidWriteRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHotGroupForbidWriteUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// HotGroupModifyName
//
// PARAMS:
//   - request: the arguments to HotGroupModifyName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupModifyName(request *HotGroupModifyNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHotGroupModifyNameUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// HotGroupPreCheck
//
// PARAMS:
//   - request: the arguments to HotGroupPreCheck
//
// RETURNS:
//   - HotGroupPreCheckResponse: The return type of the HotGroupPreCheck interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HotGroupPreCheck(request *HotGroupPreCheckRequest) (*HotGroupPreCheckResponse, error) {
	result := &HotGroupPreCheckResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHotGroupPreCheckUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HotGroupRemoveCluster
//
// PARAMS:
//   - request: the arguments to HotGroupRemoveCluster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupRemoveCluster(request *HotGroupRemoveClusterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHotGroupRemoveClusterUri(util.StringValue(request.GroupId), util.StringValue(request.InstanceId))).
		Do()
}

// HotGroupSetFlowControlRules
//
// PARAMS:
//   - request: the arguments to HotGroupSetFlowControlRules
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupSetFlowControlRules(request *HotGroupSetFlowControlRulesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHotGroupSetFlowControlRulesUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// HotGroupStaleReadable
//
// PARAMS:
//   - request: the arguments to HotGroupStaleReadable
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) HotGroupStaleReadable(request *HotGroupStaleReadableRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHotGroupStaleReadableUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// HotGroupSyncStatus
//
// PARAMS:
//   - request: the arguments to HotGroupSyncStatus
//
// RETURNS:
//   - HotGroupSyncStatusResponse: The return type of the HotGroupSyncStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HotGroupSyncStatus(request *HotGroupSyncStatusRequest) (*HotGroupSyncStatusResponse, error) {
	result := &HotGroupSyncStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getHotGroupSyncStatusUri(util.StringValue(request.GroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InstanceVersionUpgrade
//
// PARAMS:
//   - request: the arguments to InstanceVersionUpgrade
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) InstanceVersionUpgrade(request *InstanceVersionUpgradeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceVersionUpgradeUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// LogDetails
//
// PARAMS:
//   - request: the arguments to LogDetails
//
// RETURNS:
//   - LogDetailsResponse: The return type of the LogDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) LogDetails(request *LogDetailsRequest) (*LogDetailsResponse, error) {
	result := &LogDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getLogDetailsUri(util.StringValue(request.InstanceId), util.StringValue(request.LogId))).
		WithQueryParamFilter("validSeconds", util.StringValue(request.ValidSeconds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LogList
//
// PARAMS:
//   - request: the arguments to LogList
//
// RETURNS:
//   - LogListResponse: The return type of the LogList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) LogList(request *LogListRequest) (*LogListResponse, error) {
	result := &LogListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getLogListUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("fileType", "runlog").
		WithQueryParamFilter("fileType", util.StringValue(request.FileType)).
		WithQueryParamFilter("startTime", util.StringValue(request.StartTime)).
		WithQueryParamFilter("endTime", util.StringValue(request.EndTime)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ManualBackup
//
// PARAMS:
//   - request: the arguments to ManualBackup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ManualBackup(request *ManualBackupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getManualBackupUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ManuallyModifyBandwidth
//
// PARAMS:
//   - request: the arguments to ManuallyModifyBandwidth
//
// RETURNS:
//   - ManuallyModifyBandwidthResponse: The return type of the ManuallyModifyBandwidth interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ManuallyModifyBandwidth(request *ManuallyModifyBandwidthRequest) (*ManuallyModifyBandwidthResponse, error) {
	result := &ManuallyModifyBandwidthResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getManuallyModifyBandwidthUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MasterSlaveSwitch
//
// PARAMS:
//   - request: the arguments to MasterSlaveSwitch
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) MasterSlaveSwitch(request *MasterSlaveSwitchRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getMasterSlaveSwitchUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyBackupComment
//
// PARAMS:
//   - request: the arguments to ModifyBackupComment
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupComment(request *ModifyBackupCommentRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyBackupCommentUri(util.StringValue(request.InstanceId), util.StringValue(request.BatchId))).
		WithBody(request).
		Do()
}

// ModifyDeploymentSet
//
// PARAMS:
//   - request: the arguments to ModifyDeploymentSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyDeploymentSet(request *ModifyDeploymentSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyDeploymentSetUri(util.StringValue(request.DeploySetId))).
		WithBody(request).
		Do()
}

// ModifyEntrance
//
// PARAMS:
//   - request: the arguments to ModifyEntrance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyEntrance(request *ModifyEntranceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyEntranceUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyInstanceDomainName
//
// PARAMS:
//   - request: the arguments to ModifyInstanceDomainName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceDomainName(request *ModifyInstanceDomainNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyInstanceDomainNameUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyInstanceName
//
// PARAMS:
//   - request: the arguments to ModifyInstanceName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyInstanceNameUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyParameterTemplateName
//
// PARAMS:
//   - request: the arguments to ModifyParameterTemplateName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyParameterTemplateName(request *ModifyParameterTemplateNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyParameterTemplateNameUri(util.StringValue(request.TemplateShowId))).
		WithBody(request).
		Do()
}

// ModifyParameters
//
// PARAMS:
//   - request: the arguments to ModifyParameters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyParameters(request *ModifyParametersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyParametersUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyReplicationZone
//
// PARAMS:
//   - request: the arguments to ModifyReplicationZone
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyReplicationZone(request *ModifyReplicationZoneRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyReplicationZoneUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifySyncGroupName
//
// PARAMS:
//   - request: the arguments to ModifySyncGroupName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifySyncGroupName(request *ModifySyncGroupNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifySyncGroupNameUri(VERSION_V1, util.StringValue(request.SyncGroupShowId))).
		WithBody(request).
		Do()
}

// ModifyTimeWindow
//
// PARAMS:
//   - request: the arguments to ModifyTimeWindow
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTimeWindow(request *ModifyTimeWindowRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTimeWindowUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ParameterTemplateDeleteParameters
//
// PARAMS:
//   - request: the arguments to ParameterTemplateDeleteParameters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ParameterTemplateDeleteParameters(request *ParameterTemplateDeleteParametersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getParameterTemplateDeleteParametersUri(util.StringValue(request.TemplateShowId))).
		WithBody(request).
		Do()
}

// ParameterTemplateDetails
//
// PARAMS:
//   - request: the arguments to ParameterTemplateDetails
//
// RETURNS:
//   - ParameterTemplateDetailsResponse: The return type of the ParameterTemplateDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ParameterTemplateDetails(request *ParameterTemplateDetailsRequest) (*ParameterTemplateDetailsResponse, error) {
	result := &ParameterTemplateDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getParameterTemplateDetailsUri(util.StringValue(request.TemplateShowId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParameterTemplateModifyParameters
//
// PARAMS:
//   - request: the arguments to ParameterTemplateModifyParameters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ParameterTemplateModifyParameters(request *ParameterTemplateModifyParametersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getParameterTemplateModifyParametersUri(util.StringValue(request.TemplateShowId))).
		WithBody(request).
		Do()
}

// PostPaidToPrepaid
//
// PARAMS:
//   - request: the arguments to PostPaidToPrepaid
//
// RETURNS:
//   - PostPaidToPrepaidResponse: The return type of the PostPaidToPrepaid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PostPaidToPrepaid(request *PostPaidToPrepaidRequest) (*PostPaidToPrepaidResponse, error) {
	result := &PostPaidToPrepaidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPostPaidToPrepaidUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PrepaidToPostpaid
//
// PARAMS:
//   - request: the arguments to PrepaidToPostpaid
//
// RETURNS:
//   - PrepaidToPostpaidResponse: The return type of the PrepaidToPostpaid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PrepaidToPostpaid(request *PrepaidToPostpaidRequest) (*PrepaidToPostpaidResponse, error) {
	result := &PrepaidToPostpaidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPrepaidToPostpaidUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProxyNodeReplace
//
// PARAMS:
//   - request: the arguments to ProxyNodeReplace
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ProxyNodeReplace(request *ProxyNodeReplaceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getProxyNodeReplaceUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ProxyVersionUpgradeOrRestart
//
// PARAMS:
//   - request: the arguments to ProxyVersionUpgradeOrRestart
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ProxyVersionUpgradeOrRestart(request *ProxyVersionUpgradeOrRestartRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getProxyVersionUpgradeOrRestartUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// QueryIpWhitelist
//
// PARAMS:
//   - request: the arguments to QueryIpWhitelist
//
// RETURNS:
//   - QueryIpWhitelistResponse: The return type of the QueryIpWhitelist interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryIpWhitelist(request *QueryIpWhitelistRequest) (*QueryIpWhitelistResponse, error) {
	result := &QueryIpWhitelistResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryIpWhitelistUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryMemoryScalingConfig
//
// PARAMS:
//   - request: the arguments to QueryMemoryScalingConfig
//
// RETURNS:
//   - QueryMemoryScalingConfigResponse: The return type of the QueryMemoryScalingConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryMemoryScalingConfig(request *QueryMemoryScalingConfigRequest) (*QueryMemoryScalingConfigResponse, error) {
	result := &QueryMemoryScalingConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryMemoryScalingConfigUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseHotGroup
//
// PARAMS:
//   - request: the arguments to ReleaseHotGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseHotGroup(request *ReleaseHotGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseHotGroupUri(util.StringValue(request.GroupId))).
		Do()
}

// ReleaseInstance
//
// PARAMS:
//   - request: the arguments to ReleaseInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReleaseInstance(request *ReleaseInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseInstanceUri(util.StringValue(request.InstanceId))).
		Do()
}

// RenewInstance
//
// PARAMS:
//   - request: the arguments to RenewInstance
//
// RETURNS:
//   - RenewInstanceResponse: The return type of the RenewInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RenewInstance(request *RenewInstanceRequest) (*RenewInstanceResponse, error) {
	result := &RenewInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRenewInstanceUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RestartInstance
//
// PARAMS:
//   - request: the arguments to RestartInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestartInstance(request *RestartInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRestartInstanceUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// SetBackupPolicy
//
// PARAMS:
//   - request: the arguments to SetBackupPolicy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetBackupPolicy(request *SetBackupPolicyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSetBackupPolicyUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// SetClusterAsMaster
//
// PARAMS:
//   - request: the arguments to SetClusterAsMaster
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetClusterAsMaster(request *SetClusterAsMasterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSetClusterAsMasterUri(util.StringValue(request.InstanceId))).
		Do()
}

// SetClusterAsSlave
//
// PARAMS:
//   - request: the arguments to SetClusterAsSlave
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetClusterAsSlave(request *SetClusterAsSlaveRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSetClusterAsSlaveUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// SetMemoryScalingConfig
//
// PARAMS:
//   - request: the arguments to SetMemoryScalingConfig
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetMemoryScalingConfig(request *SetMemoryScalingConfigRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSetMemoryScalingConfigUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// SetPermissions
//
// PARAMS:
//   - request: the arguments to SetPermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetPermissions(request *SetPermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSetPermissionsUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// SyncGroupAddInstance
//
// PARAMS:
//   - request: the arguments to SyncGroupAddInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SyncGroupAddInstance(request *SyncGroupAddInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupAddInstanceUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// SyncGroupDelayInfo
//
// PARAMS:
//   - request: the arguments to SyncGroupDelayInfo
//
// RETURNS:
//   - SyncGroupDelayInfoResponse: The return type of the SyncGroupDelayInfo interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SyncGroupDelayInfo(request *SyncGroupDelayInfoRequest) (*SyncGroupDelayInfoResponse, error) {
	result := &SyncGroupDelayInfoResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupDelayInfoUri(util.StringValue(request.GroupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SyncGroupDetail
//
// PARAMS:
//   - request: the arguments to SyncGroupDetail
//
// RETURNS:
//   - SyncGroupDetailResponse: The return type of the SyncGroupDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SyncGroupDetail(request *SyncGroupDetailRequest) (*SyncGroupDetailResponse, error) {
	result := &SyncGroupDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupDetailUri(VERSION_V1, util.StringValue(request.SyncGroupShowId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SyncGroupList
//
// PARAMS:
//   - request: the arguments to SyncGroupList
//
// RETURNS:
//   - SyncGroupListResponse: The return type of the SyncGroupList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SyncGroupList(request *SyncGroupListRequest) (*SyncGroupListResponse, error) {
	result := &SyncGroupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSyncGroupListUri(VERSION_V1)).
		WithQueryParamFilter("page", "1").
		WithQueryParamFilter("pageSize", "10").
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SyncGroupModifyBnsgroup
//
// PARAMS:
//   - request: the arguments to SyncGroupModifyBnsgroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SyncGroupModifyBnsgroup(request *SyncGroupModifyBnsgroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSyncGroupModifyBnsgroupUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// SyncGroupPreCheck
//
// PARAMS:
//   - request: the arguments to SyncGroupPreCheck
//
// RETURNS:
//   - SyncGroupPreCheckResponse: The return type of the SyncGroupPreCheck interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SyncGroupPreCheck(request *SyncGroupPreCheckRequest) (*SyncGroupPreCheckResponse, error) {
	result := &SyncGroupPreCheckResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupPreCheckUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SyncGroupRemoveInstance
//
// PARAMS:
//   - request: the arguments to SyncGroupRemoveInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SyncGroupRemoveInstance(request *SyncGroupRemoveInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSyncGroupRemoveInstanceUri(util.StringValue(request.GroupId))).
		WithBody(request).
		Do()
}

// TdeEncryption
//
// PARAMS:
//   - request: the arguments to TdeEncryption
//
// RETURNS:
//   - TdeEncryptionResponse: The return type of the TdeEncryption interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TdeEncryption(request *TdeEncryptionRequest) (*TdeEncryptionResponse, error) {
	result := &TdeEncryptionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTdeEncryptionUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindSecurityGroup
//
// PARAMS:
//   - request: the arguments to UnbindSecurityGroup
//
// RETURNS:
//   - UnbindSecurityGroupResponse: The return type of the UnbindSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UnbindSecurityGroup(request *UnbindSecurityGroupRequest) (*UnbindSecurityGroupResponse, error) {
	result := &UnbindSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUnbindSecurityGroupUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindTags
//
// PARAMS:
//   - request: the arguments to UnbindTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindTags(request *UnbindTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUnbindTagsUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// UpdateInstanceWhiteGroup
//
// PARAMS:
//   - request: the arguments to UpdateInstanceWhiteGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceWhiteGroup(request *UpdateInstanceWhiteGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceWhiteGroupUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// UpdateSecurityGroup
//
// PARAMS:
//   - request: the arguments to UpdateSecurityGroup
//
// RETURNS:
//   - UpdateSecurityGroupResponse: The return type of the UpdateSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityGroup(request *UpdateSecurityGroupRequest) (*UpdateSecurityGroupResponse, error) {
	result := &UpdateSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateSecurityGroupUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTlsEncryption
//
// PARAMS:
//   - request: the arguments to UpdateTlsEncryption
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTlsEncryption(request *UpdateTlsEncryptionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateTlsEncryptionUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ViewSecurityGroup
//
// PARAMS:
//   - request: the arguments to ViewSecurityGroup
//
// RETURNS:
//   - ViewSecurityGroupResponse: The return type of the ViewSecurityGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ViewSecurityGroup(request *ViewSecurityGroupRequest) (*ViewSecurityGroupResponse, error) {
	result := &ViewSecurityGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getViewSecurityGroupUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
