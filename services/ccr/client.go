package ccr

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "ccr." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V1 = "v1"

	CONSTANT_INSTANCES = "instances"

	CONSTANT_CREDENTIAL = "credential"

	CONSTANT_PROJECTS = "projects"

	CONSTANT_ACCELERATORS = "accelerators"

	CONSTANT_POLICIES = "policies"

	CONSTANT_FILTERS = "filters"

	CONSTANT_SYNCS = "syncs"

	CONSTANT_REPOSITORIES = "repositories"

	CONSTANT_TAGS = "tags"

	CONSTANT_BUILDHISTORY = "buildhistory"

	CONSTANT_TRIGGERS = "triggers"

	CONSTANT_EXECUTIONS = "executions"

	CONSTANT_PUBLICLINKS = "publiclinks"

	CONSTANT_ROBOTS = "robots"

	CONSTANT_REPLICATIONS = "replications"

	CONSTANT_SCANOVERVIEW = "scanoverview"

	CONSTANT_TARGETS = "targets"

	CONSTANT_TASKS = "tasks"

	CONSTANT_SCAN = "scan"

	CONSTANT_PRIVATELINKS = "privatelinks"

	CONSTANT_CHARTS = "charts"

	CONSTANT_VERSIONS = "versions"

	CONSTANT_SECRET = "secret"

	CONSTANT_LOG = "log"

	CONSTANT_WHITELIST = "whitelist"

	CONSTANT_USERS = "users"

	CONSTANT_PROFILE = "profile"

	CONSTANT_ENABLE = "enable"

	CONSTANT_DOWNLOAD = "download"

	CONSTANT_JOBS = "jobs"

	CONSTANT_RETRY = "retry"
)

// Client of ccr service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getAddPublicNetworkWhitelistUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PUBLICLINKS + bce.URI_PREFIX + CONSTANT_WHITELIST
}
func getAddVpcLinkUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PRIVATELINKS
}
func getCreateAcceleratorFilterUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES
}
func getCreateImageMigrationRuleUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_REPLICATIONS
}
func getCreateInstanceSyncUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SYNCS
}
func getCreateProjectUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS
}
func getCreateRobotAccountUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ROBOTS
}
func getCreateTemporaryPasswordUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CREDENTIAL
}
func getCreateTriggerUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES
}
func getDeleteAcceleratorFilterUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId
}
func getDeleteAcceleratorFiltersUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES
}
func getDeleteChartUri(InstanceId string, ProjectName string, ChartName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS + bce.URI_PREFIX + ChartName
}
func getDeleteChartVersionUri(InstanceId string, ProjectName string, ChartName string, ChartVersion string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS + bce.URI_PREFIX + ChartName + bce.URI_PREFIX + CONSTANT_VERSIONS + bce.URI_PREFIX + ChartVersion
}
func getDeleteChartVersionsUri(InstanceId string, ProjectName string, ChartName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS + bce.URI_PREFIX + ChartName + bce.URI_PREFIX + CONSTANT_VERSIONS
}
func getDeleteChartsUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS
}
func getDeleteImageMigrationRuleUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_REPLICATIONS + bce.URI_PREFIX + PolicyId
}
func getDeleteInstanceSyncUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SYNCS + bce.URI_PREFIX + PolicyId
}
func getDeleteProjectUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName
}
func getDeleteProjectsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS
}
func getDeletePublicNetworkWhitelistUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PUBLICLINKS + bce.URI_PREFIX + CONSTANT_WHITELIST
}
func getDeleteRepositoriesUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES
}
func getDeleteRepositoryUri(InstanceId string, ProjectName string, RepositoryName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName
}
func getDeleteRobotAccountUri(InstanceId string, RobotID string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ROBOTS + bce.URI_PREFIX + RobotID
}
func getDeleteTagUri(InstanceId string, ProjectName string, RepositoryName string, TagName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS + bce.URI_PREFIX + TagName
}
func getDeleteTagsUri(InstanceId string, ProjectName string, RepositoryName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS
}
func getDeleteTriggerUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId
}
func getDeleteTriggersUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES
}
func getDeleteVpcLinkUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PRIVATELINKS
}
func getDownloadChartUri(InstanceId string, ProjectName string, Filename string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS + bce.URI_PREFIX + CONSTANT_DOWNLOAD + bce.URI_PREFIX + Filename
}
func getExecuteImageMigrationUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS
}
func getExecuteInstanceSyncUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS
}
func getGetAcceleratorFilterDetailUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId
}
func getGetImageMigrationExecutionRecordDetailUri(InstanceId string, ExecutionId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId
}
func getGetImageMigrationRuleDetailUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_REPLICATIONS + bce.URI_PREFIX + PolicyId
}
func getGetImageMigrationTaskLogsUri(InstanceId string, ExecutionId string, TaskId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId + bce.URI_PREFIX + CONSTANT_TASKS + bce.URI_PREFIX + TaskId + bce.URI_PREFIX + CONSTANT_LOG
}
func getGetInstanceDetailUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId
}
func getGetInstanceSyncDetailUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SYNCS + bce.URI_PREFIX + PolicyId
}
func getGetInstanceSyncExecutionDetailUri(InstanceId string, ExecutionId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId
}
func getGetInstanceSyncTaskLogsUri(InstanceId string, ExecutionId string, TaskId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId + bce.URI_PREFIX + CONSTANT_TASKS + bce.URI_PREFIX + TaskId + bce.URI_PREFIX + CONSTANT_LOG
}
func getGetProjectDetailUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName
}
func getGetPublicNetworkConfigUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PUBLICLINKS
}
func getGetRepositoryUri(InstanceId string, ProjectName string, RepositoryName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName
}
func getGetTagBuildHistoryUri(InstanceId string, ProjectName string, RepositoryName string, TagName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS + bce.URI_PREFIX + TagName + bce.URI_PREFIX + CONSTANT_BUILDHISTORY
}
func getGetTagDetailUri(InstanceId string, ProjectName string, RepositoryName string, TagName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS + bce.URI_PREFIX + TagName
}
func getGetTagsScanOverviewUri(InstanceId string, ProjectName string, RepositoryName string, TagName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS + bce.URI_PREFIX + TagName + bce.URI_PREFIX + CONSTANT_SCANOVERVIEW
}
func getGetTriggerDetailUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId
}
func getGetUserDetailUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_USERS + bce.URI_PREFIX + CONSTANT_PROFILE
}
func getListAcceleratorFiltersUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES
}
func getListChartVersionsUri(InstanceId string, ProjectName string, ChartName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS + bce.URI_PREFIX + ChartName + bce.URI_PREFIX + CONSTANT_VERSIONS
}
func getListChartsUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_CHARTS
}
func getListImageMigrationRecordsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS
}
func getListImageMigrationRulesUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_REPLICATIONS
}
func getListImageMigrationTaskRecordsUri(InstanceId string, ExecutionId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId + bce.URI_PREFIX + CONSTANT_TASKS
}
func getListInstanceSyncRecordsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS
}
func getListInstanceSyncTaskRecordsUri(InstanceId string, ExecutionId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId + bce.URI_PREFIX + CONSTANT_TASKS
}
func getListInstanceSyncsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SYNCS
}
func getListInstancesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES
}
func getListProjectsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS
}
func getListRepositoriesUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES
}
func getListRobotAccountsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ROBOTS
}
func getListTagsUri(InstanceId string, ProjectName string, RepositoryName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS
}
func getListTriggerTasksUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId + bce.URI_PREFIX + CONSTANT_JOBS
}
func getListTriggersUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES
}
func getListVpcLinksUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PRIVATELINKS
}
func getReExecuteTriggerTaskUri(InstanceId string, PolicyId string, JobId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId + bce.URI_PREFIX + CONSTANT_JOBS + bce.URI_PREFIX + JobId + bce.URI_PREFIX + CONSTANT_RETRY
}
func getRefreshRobotAccountKeyUri(InstanceId string, RobotID string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ROBOTS + bce.URI_PREFIX + RobotID + bce.URI_PREFIX + CONSTANT_SECRET
}
func getResetPasswordUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CREDENTIAL
}
func getStopImageMigrationUri(InstanceId string, ExecutionId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId
}
func getStopInstanceSyncUri(InstanceId string, ExecutionId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_EXECUTIONS + bce.URI_PREFIX + ExecutionId
}
func getTestAcceleratorFilterUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + CONSTANT_FILTERS
}
func getTestTriggerTargetAddressUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + CONSTANT_TARGETS
}
func getToggleAcceleratorFilterUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId + bce.URI_PREFIX + CONSTANT_ENABLE
}
func getToggleTriggerUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId + bce.URI_PREFIX + CONSTANT_ENABLE
}
func getTriggerTagScanUri(InstanceId string, ProjectName string, RepositoryName string, TagName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName + bce.URI_PREFIX + CONSTANT_TAGS + bce.URI_PREFIX + TagName + bce.URI_PREFIX + CONSTANT_SCAN
}
func getUpdateAcceleratorFilterUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCELERATORS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId
}
func getUpdateImageMigrationRuleUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_REPLICATIONS + bce.URI_PREFIX + PolicyId
}
func getUpdateInstanceNameUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId
}
func getUpdateInstanceSyncUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SYNCS + bce.URI_PREFIX + PolicyId
}
func getUpdateInstanceTagsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TAGS
}
func getUpdateProjectUri(InstanceId string, ProjectName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName
}
func getUpdatePublicNetworkUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PUBLICLINKS
}
func getUpdateRepositoryUri(InstanceId string, ProjectName string, RepositoryName string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROJECTS + bce.URI_PREFIX + ProjectName + bce.URI_PREFIX + CONSTANT_REPOSITORIES + bce.URI_PREFIX + RepositoryName
}
func getUpdateRobotAccountUri(InstanceId string, RobotID string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ROBOTS + bce.URI_PREFIX + RobotID
}
func getUpdateTriggerUri(InstanceId string, PolicyId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCES + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TRIGGERS + bce.URI_PREFIX + CONSTANT_POLICIES + bce.URI_PREFIX + PolicyId
}
