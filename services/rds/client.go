package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "rds." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_ACCOUNT = "account"

	CONSTANT_V2 = "v2"

	CONSTANT_TABLE_LEVEL = "tableLevel"

	CONSTANT_UPDATE_DESC = "update_desc"

	CONSTANT_STATUS = "status"

	CONSTANT_BACKUP = "backup"

	CONSTANT_BINLOGS = "binlogs"

	CONSTANT_PARA_TEMPLATE = "paraTemplate"

	CONSTANT_TEMPLATE = "template"

	CONSTANT_DETAIL = "detail"

	CONSTANT_PORT = "port"

	CONSTANT_SLOWLOGS = "slowlogs"

	CONSTANT_V1 = "v1"

	CONSTANT_TASK = "task"

	CONSTANT_DESC = "desc"

	CONSTANT_RENEW = "renew"

	CONSTANT_PASSWORD = "password"

	CONSTANT_SSL = "ssl"

	CONSTANT_STATIC = "static"

	CONSTANT_CA = "ca"

	CONSTANT_AUTO_RESIZE_CONFIG = "autoResizeConfig"

	CONSTANT_ZONE = "zone"

	CONSTANT_DATABASES = "databases"

	CONSTANT_PARAMETER = "parameter"

	CONSTANT_HISTORY = "history"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_SECURITY_IP = "securityIp"

	CONSTANT_ERRORLOGS = "errorlogs"

	CONSTANT_DUPLICATE = "duplicate"

	CONSTANT_MAINTAINTIME = "maintaintime"

	CONSTANT_TAGS = "tags"

	CONSTANT_RECOVERY_TO_SOURCE_INSTANCE_BY_SNAPSHOT = "recoveryToSourceInstanceBySnapshot"

	CONSTANT_DISK_AUTO_RESIZE = "diskAutoResize"

	CONSTANT_CONFIG = "config"

	CONSTANT_CREATE = "create"

	CONSTANT_PROBE = "probe"

	CONSTANT_LIST = "list"

	CONSTANT_PG = "pg"

	CONSTANT_LOG_LIST = "logList"

	CONSTANT_PRIVILEGES = "privileges"

	CONSTANT_RECOVERY_TO_SOURCE_INSTANCE_BY_DATETIME = "recoveryToSourceInstanceByDatetime"

	CONSTANT_TARGET = "target"

	CONSTANT_REGION = "region"

	CONSTANT_REMARK = "remark"

	CONSTANT_DOWNLOAD = "download"

	CONSTANT_AUTO_EXPANSION = "autoExpansion"
)

// Client of rds service is a kind of BceClient, so derived from BceClient
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

func getAddAccountApiAtTableGranularityUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_TABLE_LEVEL + bce.URI_PREFIX + CONSTANT_CREATE
}
func getBatchConfigurationOfInstancesUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getBindingTagsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TAGS
}
func getChangeAccountPasswordUri(version string, InstanceId string, AccountName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + AccountName + bce.URI_PREFIX + CONSTANT_PASSWORD
}
func getCheckTheWhitelistUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_IP
}
func getClusterStatusCheckUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_STATUS
}
func getCopyParameterTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_PARA_TEMPLATE + bce.URI_PREFIX + CONSTANT_DUPLICATE + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getCreateADatabaseUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DATABASES
}
func getCreateAPrimaryInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateAProxyInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateAccountUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT
}
func getCreateReadOnlyInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDeleteAccountUri(version string, InstanceId string, AccountName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + AccountName
}
func getDeleteTheDatabaseUri(version string, InstanceId string, DbName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DATABASES + bce.URI_PREFIX + DbName
}
func getDeleteTheSpecifiedBackupSetUri(version string, InstanceId string, SnapshotId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + SnapshotId
}
func getDialTestInterfaceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROBE
}
func getDoesTheInstanceSupportEnablingAutomaticExpansionUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AUTO_EXPANSION
}
func getErrorLogListUri(version string, InstanceId string, Datetime string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ERRORLOGS + bce.URI_PREFIX + CONSTANT_LOG_LIST + bce.URI_PREFIX + Datetime
}
func getGetParameterTemplateDetailsUri(version string, TemplateId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_PARA_TEMPLATE + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_DETAIL + bce.URI_PREFIX + TemplateId
}
func getGetParameterTemplateListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_PARA_TEMPLATE
}
func getGetPgLogDetailsUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PG + bce.URI_PREFIX + CONSTANT_DOWNLOAD
}
func getGetPgLogListUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PG + bce.URI_PREFIX + CONSTANT_LIST
}
func getInstanceEnablesAutomaticRenewalUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getManuallyInitiateFullPhysicalBackupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP
}
func getModifyAccountNotesUri(version string, InstanceId string, AccountName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + AccountName + bce.URI_PREFIX + CONSTANT_DESC
}
func getModifyAccountPermissionsUri(version string, InstanceId string, AccountName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + AccountName + bce.URI_PREFIX + CONSTANT_PRIVILEGES
}
func getModifyBackupStrategyUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getModifyConfigurationParametersUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PARAMETER
}
func getModifyDatabasePortUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PORT
}
func getModifyDescriptionInformationUri(version string, InstanceId string, DbName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DATABASES + bce.URI_PREFIX + DbName + bce.URI_PREFIX + CONSTANT_REMARK
}
func getObtainBinlogInformationUri(version string, InstanceId string, LogId string, DownloadValidTimeInSec string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BINLOGS + bce.URI_PREFIX + LogId + bce.URI_PREFIX + DownloadValidTimeInSec
}
func getObtainSslEncryptionCertificateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_SSL + bce.URI_PREFIX + CONSTANT_STATIC + bce.URI_PREFIX + CONSTANT_CA
}
func getObtainSslEncryptionInformationUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_SSL + bce.URI_PREFIX + InstanceId
}
func getQueryAccountDetailsUri(version string, InstanceId string, AccountName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + AccountName
}
func getQueryAccountListUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT
}
func getQueryAndModifyHistoryUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PARAMETER + bce.URI_PREFIX + CONSTANT_HISTORY
}
func getQueryBackupDetailsUri(version string, InstanceId string, BackupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + BackupId
}
func getQueryBackupListUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP
}
func getQueryInstanceListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getQueryParameterListUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PARAMETER
}
func getQuerySubnetListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getQueryTaskListUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TASK
}
func getQueryTheDetailsOfTheSpecifiedInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getQueryTheListOfAvailableZonesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ZONE
}
func getReleaseInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getRenewalExampleUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RENEW
}
func getRestartTheInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getRestoreLibraryTablesBasedOnBackupSetsUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_RECOVERY_TO_SOURCE_INSTANCE_BY_SNAPSHOT
}
func getRestoreTheDatabaseTableByTimePointUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_RECOVERY_TO_SOURCE_INSTANCE_BY_DATETIME
}
func getRetrieveDatabaseListUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DATABASES
}
func getRetrieveInstanceAutoScalingConfigUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AUTO_RESIZE_CONFIG
}
func getRetrieveTheBinlogListUri(version string, InstanceId string, Datetime string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BINLOGS + bce.URI_PREFIX + Datetime
}
func getSetEncryptionPolicyApiUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getSlowLogListUri(version string, InstanceId string, Datetime string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SLOWLOGS + bce.URI_PREFIX + CONSTANT_LOG_LIST + bce.URI_PREFIX + Datetime
}
func getTableGranularityQueryUserDetailsApiUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_TABLE_LEVEL + bce.URI_PREFIX + CONSTANT_DETAIL
}
func getTableGranularityQueryUserListApiUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_TABLE_LEVEL + bce.URI_PREFIX + CONSTANT_LIST
}
func getTableGranularityUpdateUserDescriptionApiUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_TABLE_LEVEL + bce.URI_PREFIX + CONSTANT_UPDATE_DESC
}
func getTargetRegionMappingListApiUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_TARGET + bce.URI_PREFIX + CONSTANT_REGION
}
func getTransformationExampleUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateAutomaticStorageExpansionConfigurationUri(version string, InstanceId string, Action string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DISK_AUTO_RESIZE + bce.URI_PREFIX + CONSTANT_CONFIG + bce.URI_PREFIX + Action
}
func getUpdateInstanceConnectionInformationUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateInstanceNameUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateInstancePublicNetworkAccessStatusUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateInstanceSynchronizationModeUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateTimeWindowUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MAINTAINTIME
}
func getUpdateWhitelistUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_IP
}
