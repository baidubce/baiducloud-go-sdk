package ccr

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// AddPublicNetworkWhitelist
//
// PARAMS:
//   - request: the arguments to AddPublicNetworkWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddPublicNetworkWhitelist(request *AddPublicNetworkWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddPublicNetworkWhitelistUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// AddVpcLink
//
// PARAMS:
//   - request: the arguments to AddVpcLink
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddVpcLink(request *AddVpcLinkRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddVpcLinkUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateAcceleratorFilter
//
// PARAMS:
//   - request: the arguments to CreateAcceleratorFilter
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAcceleratorFilter(request *CreateAcceleratorFilterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAcceleratorFilterUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateImageMigrationRule
//
// PARAMS:
//   - request: the arguments to CreateImageMigrationRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateImageMigrationRule(request *CreateImageMigrationRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateImageMigrationRuleUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateInstanceSync
//
// PARAMS:
//   - request: the arguments to CreateInstanceSync
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceSync(request *CreateInstanceSyncRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceSyncUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateProject
//
// PARAMS:
//   - request: the arguments to CreateProject
//
// RETURNS:
//   - CreateProjectResponse: The return type of the CreateProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateProject(request *CreateProjectRequest) (*CreateProjectResponse, error) {
	result := &CreateProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateProjectUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateRobotAccount
//
// PARAMS:
//   - request: the arguments to CreateRobotAccount
//
// RETURNS:
//   - CreateRobotAccountResponse: The return type of the CreateRobotAccount interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateRobotAccount(request *CreateRobotAccountRequest) (*CreateRobotAccountResponse, error) {
	result := &CreateRobotAccountResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRobotAccountUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTemporaryPassword
//
// PARAMS:
//   - request: the arguments to CreateTemporaryPassword
//
// RETURNS:
//   - CreateTemporaryPasswordResponse: The return type of the CreateTemporaryPassword interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateTemporaryPassword(request *CreateTemporaryPasswordRequest) (*CreateTemporaryPasswordResponse, error) {
	result := &CreateTemporaryPasswordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTemporaryPasswordUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTrigger
//
// PARAMS:
//   - request: the arguments to CreateTrigger
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateTrigger(request *CreateTriggerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateTriggerUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeleteAcceleratorFilter
//
// PARAMS:
//   - request: the arguments to DeleteAcceleratorFilter
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAcceleratorFilter(request *DeleteAcceleratorFilterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAcceleratorFilterUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		Do()
}

// DeleteAcceleratorFilters
//
// PARAMS:
//   - request: the arguments to DeleteAcceleratorFilters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAcceleratorFilters(request *DeleteAcceleratorFiltersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAcceleratorFiltersUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeleteChart
//
// PARAMS:
//   - request: the arguments to DeleteChart
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteChart(request *DeleteChartRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteChartUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.ChartName))).
		Do()
}

// DeleteChartVersion
//
// PARAMS:
//   - request: the arguments to DeleteChartVersion
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteChartVersion(request *DeleteChartVersionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteChartVersionUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.ChartName), util.StringValue(request.ChartVersion))).
		Do()
}

// DeleteChartVersions
//
// PARAMS:
//   - request: the arguments to DeleteChartVersions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteChartVersions(request *DeleteChartVersionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteChartVersionsUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.ChartName))).
		WithBody(request).
		Do()
}

// DeleteCharts
//
// PARAMS:
//   - request: the arguments to DeleteCharts
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCharts(request *DeleteChartsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteChartsUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		WithBody(request).
		Do()
}

// DeleteImageMigrationRule
//
// PARAMS:
//   - request: the arguments to DeleteImageMigrationRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteImageMigrationRule(request *DeleteImageMigrationRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteImageMigrationRuleUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		Do()
}

// DeleteInstanceSync
//
// PARAMS:
//   - request: the arguments to DeleteInstanceSync
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceSync(request *DeleteInstanceSyncRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteInstanceSyncUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		Do()
}

// DeleteProject
//
// PARAMS:
//   - request: the arguments to DeleteProject
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteProject(request *DeleteProjectRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteProjectUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		Do()
}

// DeleteProjects
//
// PARAMS:
//   - request: the arguments to DeleteProjects
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteProjects(request *DeleteProjectsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteProjectsUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeletePublicNetworkWhitelist
//
// PARAMS:
//   - request: the arguments to DeletePublicNetworkWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePublicNetworkWhitelist(request *DeletePublicNetworkWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePublicNetworkWhitelistUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeleteRepositories
//
// PARAMS:
//   - request: the arguments to DeleteRepositories
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRepositories(request *DeleteRepositoriesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRepositoriesUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		WithBody(request).
		Do()
}

// DeleteRepository
//
// PARAMS:
//   - request: the arguments to DeleteRepository
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRepositoryUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName))).
		Do()
}

// DeleteRobotAccount
//
// PARAMS:
//   - request: the arguments to DeleteRobotAccount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRobotAccount(request *DeleteRobotAccountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRobotAccountUri(util.StringValue(request.InstanceId), util.StringValue(request.RobotID))).
		Do()
}

// DeleteTag
//
// PARAMS:
//   - request: the arguments to DeleteTag
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteTag(request *DeleteTagRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTagUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName), util.StringValue(request.TagName))).
		Do()
}

// DeleteTags
//
// PARAMS:
//   - request: the arguments to DeleteTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteTags(request *DeleteTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTagsUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName))).
		WithBody(request).
		Do()
}

// DeleteTrigger
//
// PARAMS:
//   - request: the arguments to DeleteTrigger
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteTrigger(request *DeleteTriggerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTriggerUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		Do()
}

// DeleteTriggers
//
// PARAMS:
//   - request: the arguments to DeleteTriggers
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteTriggers(request *DeleteTriggersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTriggersUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DeleteVpcLink
//
// PARAMS:
//   - request: the arguments to DeleteVpcLink
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteVpcLink(request *DeleteVpcLinkRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteVpcLinkUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// DownloadChart
//
// PARAMS:
//   - request: the arguments to DownloadChart
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DownloadChart(request *DownloadChartRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDownloadChartUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.Filename))).
		Do()
}

// ExecuteImageMigration
//
// PARAMS:
//   - request: the arguments to ExecuteImageMigration
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ExecuteImageMigration(request *ExecuteImageMigrationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getExecuteImageMigrationUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ExecuteInstanceSync
//
// PARAMS:
//   - request: the arguments to ExecuteInstanceSync
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ExecuteInstanceSync(request *ExecuteInstanceSyncRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getExecuteInstanceSyncUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// GetAcceleratorFilterDetail
//
// PARAMS:
//   - request: the arguments to GetAcceleratorFilterDetail
//
// RETURNS:
//   - GetAcceleratorFilterDetailResponse: The return type of the GetAcceleratorFilterDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAcceleratorFilterDetail(request *GetAcceleratorFilterDetailRequest) (*GetAcceleratorFilterDetailResponse, error) {
	result := &GetAcceleratorFilterDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetAcceleratorFilterDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetImageMigrationExecutionRecordDetail
//
// PARAMS:
//   - request: the arguments to GetImageMigrationExecutionRecordDetail
//
// RETURNS:
//   - GetImageMigrationExecutionRecordDetailResponse: The return type of the GetImageMigrationExecutionRecordDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetImageMigrationExecutionRecordDetail(request *GetImageMigrationExecutionRecordDetailRequest) (*GetImageMigrationExecutionRecordDetailResponse, error) {
	result := &GetImageMigrationExecutionRecordDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetImageMigrationExecutionRecordDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetImageMigrationRuleDetail
//
// PARAMS:
//   - request: the arguments to GetImageMigrationRuleDetail
//
// RETURNS:
//   - GetImageMigrationRuleDetailResponse: The return type of the GetImageMigrationRuleDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetImageMigrationRuleDetail(request *GetImageMigrationRuleDetailRequest) (*GetImageMigrationRuleDetailResponse, error) {
	result := &GetImageMigrationRuleDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetImageMigrationRuleDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetImageMigrationTaskLogs
//
// PARAMS:
//   - request: the arguments to GetImageMigrationTaskLogs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetImageMigrationTaskLogs(request *GetImageMigrationTaskLogsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetImageMigrationTaskLogsUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId), util.StringValue(request.TaskId))).
		Do()
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

// GetInstanceSyncDetail
//
// PARAMS:
//   - request: the arguments to GetInstanceSyncDetail
//
// RETURNS:
//   - GetInstanceSyncDetailResponse: The return type of the GetInstanceSyncDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceSyncDetail(request *GetInstanceSyncDetailRequest) (*GetInstanceSyncDetailResponse, error) {
	result := &GetInstanceSyncDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceSyncDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceSyncExecutionDetail
//
// PARAMS:
//   - request: the arguments to GetInstanceSyncExecutionDetail
//
// RETURNS:
//   - GetInstanceSyncExecutionDetailResponse: The return type of the GetInstanceSyncExecutionDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceSyncExecutionDetail(request *GetInstanceSyncExecutionDetailRequest) (*GetInstanceSyncExecutionDetailResponse, error) {
	result := &GetInstanceSyncExecutionDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceSyncExecutionDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceSyncTaskLogs
//
// PARAMS:
//   - request: the arguments to GetInstanceSyncTaskLogs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetInstanceSyncTaskLogs(request *GetInstanceSyncTaskLogsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceSyncTaskLogsUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId), util.StringValue(request.TaskId))).
		Do()
}

// GetProjectDetail
//
// PARAMS:
//   - request: the arguments to GetProjectDetail
//
// RETURNS:
//   - GetProjectDetailResponse: The return type of the GetProjectDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetProjectDetail(request *GetProjectDetailRequest) (*GetProjectDetailResponse, error) {
	result := &GetProjectDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetProjectDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicNetworkConfig
//
// PARAMS:
//   - request: the arguments to GetPublicNetworkConfig
//
// RETURNS:
//   - GetPublicNetworkConfigResponse: The return type of the GetPublicNetworkConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPublicNetworkConfig(request *GetPublicNetworkConfigRequest) (*GetPublicNetworkConfigResponse, error) {
	result := &GetPublicNetworkConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPublicNetworkConfigUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRepository
//
// PARAMS:
//   - request: the arguments to GetRepository
//
// RETURNS:
//   - GetRepositoryResponse: The return type of the GetRepository interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRepository(request *GetRepositoryRequest) (*GetRepositoryResponse, error) {
	result := &GetRepositoryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRepositoryUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTagBuildHistory
//
// PARAMS:
//   - request: the arguments to GetTagBuildHistory
//
// RETURNS:
//   - GetTagBuildHistoryResponse: The return type of the GetTagBuildHistory interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTagBuildHistory(request *GetTagBuildHistoryRequest) (*GetTagBuildHistoryResponse, error) {
	result := &GetTagBuildHistoryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTagBuildHistoryUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName), util.StringValue(request.TagName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTagDetail
//
// PARAMS:
//   - request: the arguments to GetTagDetail
//
// RETURNS:
//   - GetTagDetailResponse: The return type of the GetTagDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTagDetail(request *GetTagDetailRequest) (*GetTagDetailResponse, error) {
	result := &GetTagDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTagDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName), util.StringValue(request.TagName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTagsScanOverview
//
// PARAMS:
//   - request: the arguments to GetTagsScanOverview
//
// RETURNS:
//   - GetTagsScanOverviewResponse: The return type of the GetTagsScanOverview interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTagsScanOverview(request *GetTagsScanOverviewRequest) (*GetTagsScanOverviewResponse, error) {
	result := &GetTagsScanOverviewResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTagsScanOverviewUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName), util.StringValue(request.TagName))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTriggerDetail
//
// PARAMS:
//   - request: the arguments to GetTriggerDetail
//
// RETURNS:
//   - GetTriggerDetailResponse: The return type of the GetTriggerDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTriggerDetail(request *GetTriggerDetailRequest) (*GetTriggerDetailResponse, error) {
	result := &GetTriggerDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTriggerDetailUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserDetail
//
// PARAMS:
//   - request: the arguments to GetUserDetail
//
// RETURNS:
//   - GetUserDetailResponse: The return type of the GetUserDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetUserDetail(request *GetUserDetailRequest) (*GetUserDetailResponse, error) {
	result := &GetUserDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetUserDetailUri()).
		WithQueryParamFilter("userId", util.StringValue(request.UserId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAcceleratorFilters
//
// PARAMS:
//   - request: the arguments to ListAcceleratorFilters
//
// RETURNS:
//   - ListAcceleratorFiltersResponse: The return type of the ListAcceleratorFilters interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAcceleratorFilters(request *ListAcceleratorFiltersRequest) (*ListAcceleratorFiltersResponse, error) {
	result := &ListAcceleratorFiltersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAcceleratorFiltersUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("policyName", util.StringValue(request.PolicyName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListChartVersions
//
// PARAMS:
//   - request: the arguments to ListChartVersions
//
// RETURNS:
//   - ListChartVersionsResponse: The return type of the ListChartVersions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListChartVersions(request *ListChartVersionsRequest) (*ListChartVersionsResponse, error) {
	result := &ListChartVersionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListChartVersionsUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.ChartName))).
		WithQueryParamFilter("chartVersion", util.StringValue(request.ChartVersion)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCharts
//
// PARAMS:
//   - request: the arguments to ListCharts
//
// RETURNS:
//   - ListChartsResponse: The return type of the ListCharts interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListCharts(request *ListChartsRequest) (*ListChartsResponse, error) {
	result := &ListChartsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListChartsUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		WithQueryParamFilter("chartName", util.StringValue(request.ChartName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListImageMigrationRecords
//
// PARAMS:
//   - request: the arguments to ListImageMigrationRecords
//
// RETURNS:
//   - ListImageMigrationRecordsResponse: The return type of the ListImageMigrationRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListImageMigrationRecords(request *ListImageMigrationRecordsRequest) (*ListImageMigrationRecordsResponse, error) {
	result := &ListImageMigrationRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListImageMigrationRecordsUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("policyId", util.StringValue(request.PolicyId)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListImageMigrationRules
//
// PARAMS:
//   - request: the arguments to ListImageMigrationRules
//
// RETURNS:
//   - ListImageMigrationRulesResponse: The return type of the ListImageMigrationRules interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListImageMigrationRules(request *ListImageMigrationRulesRequest) (*ListImageMigrationRulesResponse, error) {
	result := &ListImageMigrationRulesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListImageMigrationRulesUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("policyName", util.StringValue(request.PolicyName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListImageMigrationTaskRecords
//
// PARAMS:
//   - request: the arguments to ListImageMigrationTaskRecords
//
// RETURNS:
//   - ListImageMigrationTaskRecordsResponse: The return type of the ListImageMigrationTaskRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListImageMigrationTaskRecords(request *ListImageMigrationTaskRecordsRequest) (*ListImageMigrationTaskRecordsResponse, error) {
	result := &ListImageMigrationTaskRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListImageMigrationTaskRecordsUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstanceSyncRecords
//
// PARAMS:
//   - request: the arguments to ListInstanceSyncRecords
//
// RETURNS:
//   - ListInstanceSyncRecordsResponse: The return type of the ListInstanceSyncRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstanceSyncRecords(request *ListInstanceSyncRecordsRequest) (*ListInstanceSyncRecordsResponse, error) {
	result := &ListInstanceSyncRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstanceSyncRecordsUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("policyId", util.StringValue(request.PolicyId)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstanceSyncTaskRecords
//
// PARAMS:
//   - request: the arguments to ListInstanceSyncTaskRecords
//
// RETURNS:
//   - ListInstanceSyncTaskRecordsResponse: The return type of the ListInstanceSyncTaskRecords interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstanceSyncTaskRecords(request *ListInstanceSyncTaskRecordsRequest) (*ListInstanceSyncTaskRecordsResponse, error) {
	result := &ListInstanceSyncTaskRecordsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstanceSyncTaskRecordsUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListInstanceSyncs
//
// PARAMS:
//   - request: the arguments to ListInstanceSyncs
//
// RETURNS:
//   - ListInstanceSyncsResponse: The return type of the ListInstanceSyncs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstanceSyncs(request *ListInstanceSyncsRequest) (*ListInstanceSyncsResponse, error) {
	result := &ListInstanceSyncsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListInstanceSyncsUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("policyName", util.StringValue(request.PolicyName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
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
		WithURL(getListInstancesUri()).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("keywordType", util.StringValue(request.KeywordType)).
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithQueryParamFilter("acrossregion", util.StringValue(request.Acrossregion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListProjects
//
// PARAMS:
//   - request: the arguments to ListProjects
//
// RETURNS:
//   - ListProjectsResponse: The return type of the ListProjects interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListProjects(request *ListProjectsRequest) (*ListProjectsResponse, error) {
	result := &ListProjectsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListProjectsUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("projectName", util.StringValue(request.ProjectName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRepositories
//
// PARAMS:
//   - request: the arguments to ListRepositories
//
// RETURNS:
//   - ListRepositoriesResponse: The return type of the ListRepositories interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRepositories(request *ListRepositoriesRequest) (*ListRepositoriesResponse, error) {
	result := &ListRepositoriesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRepositoriesUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		WithQueryParamFilter("repositoryName", util.StringValue(request.RepositoryName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRobotAccounts
//
// PARAMS:
//   - request: the arguments to ListRobotAccounts
//
// RETURNS:
//   - ListRobotAccountsResponse: The return type of the ListRobotAccounts interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRobotAccounts(request *ListRobotAccountsRequest) (*ListRobotAccountsResponse, error) {
	result := &ListRobotAccountsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRobotAccountsUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTags
//
// PARAMS:
//   - request: the arguments to ListTags
//
// RETURNS:
//   - ListTagsResponse: The return type of the ListTags interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTags(request *ListTagsRequest) (*ListTagsResponse, error) {
	result := &ListTagsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTagsUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName))).
		WithQueryParamFilter("tagName", util.StringValue(request.TagName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTriggerTasks
//
// PARAMS:
//   - request: the arguments to ListTriggerTasks
//
// RETURNS:
//   - ListTriggerTasksResponse: The return type of the ListTriggerTasks interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTriggerTasks(request *ListTriggerTasksRequest) (*ListTriggerTasksResponse, error) {
	result := &ListTriggerTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTriggerTasksUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTriggers
//
// PARAMS:
//   - request: the arguments to ListTriggers
//
// RETURNS:
//   - ListTriggersResponse: The return type of the ListTriggers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTriggers(request *ListTriggersRequest) (*ListTriggersResponse, error) {
	result := &ListTriggersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTriggersUri(util.StringValue(request.InstanceId))).
		WithQueryParamFilter("policyName", util.StringValue(request.PolicyName)).
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListVpcLinks
//
// PARAMS:
//   - request: the arguments to ListVpcLinks
//
// RETURNS:
//   - ListVpcLinksResponse: The return type of the ListVpcLinks interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListVpcLinks(request *ListVpcLinksRequest) (*ListVpcLinksResponse, error) {
	result := &ListVpcLinksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListVpcLinksUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReExecuteTriggerTask
//
// PARAMS:
//   - request: the arguments to ReExecuteTriggerTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ReExecuteTriggerTask(request *ReExecuteTriggerTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getReExecuteTriggerTaskUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId), util.StringValue(request.JobId))).
		Do()
}

// RefreshRobotAccountKey
//
// PARAMS:
//   - request: the arguments to RefreshRobotAccountKey
//
// RETURNS:
//   - RefreshRobotAccountKeyResponse: The return type of the RefreshRobotAccountKey interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RefreshRobotAccountKey(request *RefreshRobotAccountKeyRequest) (*RefreshRobotAccountKeyResponse, error) {
	result := &RefreshRobotAccountKeyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRefreshRobotAccountKeyUri(util.StringValue(request.InstanceId), util.StringValue(request.RobotID))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResetPassword
//
// PARAMS:
//   - request: the arguments to ResetPassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResetPassword(request *ResetPasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getResetPasswordUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// StopImageMigration
//
// PARAMS:
//   - request: the arguments to StopImageMigration
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StopImageMigration(request *StopImageMigrationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStopImageMigrationUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId))).
		Do()
}

// StopInstanceSync
//
// PARAMS:
//   - request: the arguments to StopInstanceSync
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StopInstanceSync(request *StopInstanceSyncRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStopInstanceSyncUri(util.StringValue(request.InstanceId), util.StringValue(request.ExecutionId))).
		Do()
}

// TestAcceleratorFilter
//
// PARAMS:
//   - request: the arguments to TestAcceleratorFilter
//
// RETURNS:
//   - TestAcceleratorFilterResponse: The return type of the TestAcceleratorFilter interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TestAcceleratorFilter(request *TestAcceleratorFilterRequest) (*TestAcceleratorFilterResponse, error) {
	result := &TestAcceleratorFilterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTestAcceleratorFilterUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TestTriggerTargetAddress
//
// PARAMS:
//   - request: the arguments to TestTriggerTargetAddress
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) TestTriggerTargetAddress(request *TestTriggerTargetAddressRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTestTriggerTargetAddressUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ToggleAcceleratorFilter
//
// PARAMS:
//   - request: the arguments to ToggleAcceleratorFilter
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ToggleAcceleratorFilter(request *ToggleAcceleratorFilterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getToggleAcceleratorFilterUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithQueryParamFilter("enabled", util.StringValue(request.Enabled)).
		Do()
}

// ToggleTrigger
//
// PARAMS:
//   - request: the arguments to ToggleTrigger
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ToggleTrigger(request *ToggleTriggerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getToggleTriggerUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithQueryParamFilter("enabled", util.StringValue(request.Enabled)).
		Do()
}

// TriggerTagScan
//
// PARAMS:
//   - request: the arguments to TriggerTagScan
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) TriggerTagScan(request *TriggerTagScanRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTriggerTagScanUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName), util.StringValue(request.TagName))).
		Do()
}

// UpdateAcceleratorFilter
//
// PARAMS:
//   - request: the arguments to UpdateAcceleratorFilter
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAcceleratorFilter(request *UpdateAcceleratorFilterRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAcceleratorFilterUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithBody(request).
		Do()
}

// UpdateImageMigrationRule
//
// PARAMS:
//   - request: the arguments to UpdateImageMigrationRule
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateImageMigrationRule(request *UpdateImageMigrationRuleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateImageMigrationRuleUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithBody(request).
		Do()
}

// UpdateInstanceName
//
// PARAMS:
//   - request: the arguments to UpdateInstanceName
//
// RETURNS:
//   - UpdateInstanceNameResponse: The return type of the UpdateInstanceName interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (*UpdateInstanceNameResponse, error) {
	result := &UpdateInstanceNameResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceNameUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateInstanceSync
//
// PARAMS:
//   - request: the arguments to UpdateInstanceSync
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceSync(request *UpdateInstanceSyncRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceSyncUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithBody(request).
		Do()
}

// UpdateInstanceTags
//
// PARAMS:
//   - request: the arguments to UpdateInstanceTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceTags(request *UpdateInstanceTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceTagsUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// UpdateProject
//
// PARAMS:
//   - request: the arguments to UpdateProject
//
// RETURNS:
//   - UpdateProjectResponse: The return type of the UpdateProject interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateProject(request *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	result := &UpdateProjectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateProjectUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdatePublicNetwork
//
// PARAMS:
//   - request: the arguments to UpdatePublicNetwork
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePublicNetwork(request *UpdatePublicNetworkRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdatePublicNetworkUri(util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// UpdateRepository
//
// PARAMS:
//   - request: the arguments to UpdateRepository
//
// RETURNS:
//   - UpdateRepositoryResponse: The return type of the UpdateRepository interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateRepository(request *UpdateRepositoryRequest) (*UpdateRepositoryResponse, error) {
	result := &UpdateRepositoryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRepositoryUri(util.StringValue(request.InstanceId), util.StringValue(request.ProjectName), util.StringValue(request.RepositoryName))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateRobotAccount
//
// PARAMS:
//   - request: the arguments to UpdateRobotAccount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRobotAccount(request *UpdateRobotAccountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRobotAccountUri(util.StringValue(request.InstanceId), util.StringValue(request.RobotID))).
		WithBody(request).
		Do()
}

// UpdateTrigger
//
// PARAMS:
//   - request: the arguments to UpdateTrigger
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTrigger(request *UpdateTriggerRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateTriggerUri(util.StringValue(request.InstanceId), util.StringValue(request.PolicyId))).
		WithBody(request).
		Do()
}
