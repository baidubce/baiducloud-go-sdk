package scs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "scs." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V2 = "v2"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_V1 = "v1"

	CONSTANT_BACKUP = "backup"

	CONSTANT_URL = "url"

	CONSTANT_TLS = "tls"

	CONSTANT_BIND_TAG = "bindTag"

	CONSTANT_ACL_USER_ACTIONS = "aclUserActions"

	CONSTANT_AUTHORITY = "authority"

	CONSTANT_LOG = "log"

	CONSTANT_RECYCLER = "recycler"

	CONSTANT_LIST = "list"

	CONSTANT_ENTRANCE = "entrance"

	CONSTANT_DISCONNECT = "disconnect"

	CONSTANT_WHITELIST = "whitelist"

	CONSTANT_POLICY = "policy"

	CONSTANT_TEMPLATE = "template"

	CONSTANT_FLUSH = "flush"

	CONSTANT_GROUP = "group"

	CONSTANT_CREATE = "create"

	CONSTANT_SWAP_DOMAIN = "swapDomain"

	CONSTANT_ZONE = "zone"

	CONSTANT_SYNC_GROUP = "syncGroup"

	CONSTANT_MODIFY_BNS_GROUP = "modifyBnsGroup"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_RESTART = "restart"

	CONSTANT_DEPLOY_SET = "deploySet"

	CONSTANT_RELEASE = "release"

	CONSTANT_RENEW = "renew"

	CONSTANT_APPLY = "apply"

	CONSTANT_MODIFY_PARAMS = "modifyParams"

	CONSTANT_RECOVER__H_T_T_P = "recover HTTP"

	CONSTANT_1_1 = "1.1"

	CONSTANT_QPS = "qps"

	CONSTANT_RENAME = "rename"

	CONSTANT_MODIFY_BANDWIDTH = "modifyBandwidth"

	CONSTANT_SECURITY_GROUP = "securityGroup"

	CONSTANT_UPDATE = "update"

	CONSTANT_DELETE__H_T_T_P = "delete HTTP"

	CONSTANT_SECURITY_IP = "securityIp"

	CONSTANT_MODIFY_POLICY = "modifyPolicy"

	CONSTANT_ADD_PARAMS = "addParams"

	CONSTANT_DELETE_AUTO_SCALING_CONFIG = "deleteAutoScalingConfig"

	CONSTANT_REMOVE_CLUSTER = "removeCluster"

	CONSTANT_CHECK = "check"

	CONSTANT_PRICE = "price"

	CONSTANT_STALE_READABLE = "stale_readable"

	CONSTANT_FORBID_WRITE = "forbidWrite"

	CONSTANT_BLB_STATUS = "blbStatus"

	CONSTANT_UN_BIND_TAG = "unBindTag"

	CONSTANT_TO_PREPAY = "toPrepay"

	CONSTANT_SET_AS_MASTER = "setAsMaster"

	CONSTANT_AZONE_MIGRATION = "azoneMigration"

	CONSTANT_MODIFY_ENTRANCE = "modifyEntrance"

	CONSTANT_STATUS = "status"

	CONSTANT_PARAMETER = "parameter"

	CONSTANT_USAGE = "usage"

	CONSTANT_DELETE = "delete"

	CONSTANT_RENAME_DOMAIN = "renameDomain"

	CONSTANT_RECORD = "record"

	CONSTANT_PROXY_NODE = "proxyNode"

	CONSTANT_TIME_WINDOW = "timeWindow"

	CONSTANT_AUDIT = "audit"

	CONSTANT_SWITCH = "switch"

	CONSTANT_MODIFY_PASSWD = "modifyPasswd"

	CONSTANT_SET_AS_LEADER = "setAsLeader"

	CONSTANT_SWITCH_MASTER_SLAVE = "switchMasterSlave"

	CONSTANT_MODIFY_PASSWORD = "modifyPassword"

	CONSTANT_UPGRADE_PROXY = "upgradeProxy"

	CONSTANT_COMMENT = "comment"

	CONSTANT_CLUSTER_TYPE_CHANGE = "clusterTypeChange"

	CONSTANT_NODETYPES = "nodetypes"

	CONSTANT_CANCEL_TO_POSTPAY = "cancelToPostpay"

	CONSTANT_DELETE_PARAMS = "deleteParams"

	CONSTANT_QUIT = "quit"

	CONSTANT_AUTO_SCALING_CONFIG = "autoScalingConfig"

	CONSTANT_UNBIND = "unbind"

	CONSTANT_SYSTEM = "system"

	CONSTANT_JOIN = "join"

	CONSTANT_TO_POSTPAY = "toPostpay"

	CONSTANT_DELAY_INFO = "delayInfo"

	CONSTANT_UPGRADE_VERSION = "upgradeVersion"

	CONSTANT_SET_AS_SLAVE = "setAsSlave"

	CONSTANT_SYNC_STATUS = "syncStatus"

	CONSTANT_ADD_CLUSTER = "addCluster"

	CONSTANT_CHANGE = "change"

	CONSTANT_BIND = "bind"

	CONSTANT_TDE = "tde"
)

// Client of scs service is a kind of BceClient, so derived from BceClient
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

func getAccountListUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACL_USER_ACTIONS
}
func getAddIpWhitelistUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_IP
}
func getAddParametersToParameterTemplateUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_ADD_PARAMS + bce.URI_PREFIX + TemplateShowId
}
func getApplicationParameterTemplateUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_APPLY + bce.URI_PREFIX + TemplateShowId
}
func getAuditLogSwitchUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_LOG + bce.URI_PREFIX + CONSTANT_AUDIT + bce.URI_PREFIX + CONSTANT_SWITCH
}
func getBatchRestoreInstancesUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_RECOVER__H_T_T_P + bce.URI_PREFIX + CONSTANT_1_1
}
func getBindSecurityGroupUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_BIND
}
func getBindTagsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BIND_TAG
}
func getCancelPrepaidToPostpaidUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_CANCEL_TO_POSTPAY
}
func getChangeAccessPasswordUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_PASSWORD
}
func getChangeAccountPasswordUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACL_USER_ACTIONS + bce.URI_PREFIX + CONSTANT_MODIFY_PASSWD
}
func getChangeConfigurationUri(InstanceId string, ClientToken string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CHANGE
}
func getClearInstanceUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_FLUSH
}
func getClusterStatusCheckUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_STATUS
}
func getClusterTypeUpgradeUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_CLUSTER_TYPE_CHANGE
}
func getCreateAccountUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACL_USER_ACTIONS
}
func getCreateAnInstanceUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateDeploymentSetUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DEPLOY_SET
}
func getCreateEntranceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTRANCE + bce.URI_PREFIX + CONSTANT_CREATE + bce.URI_PREFIX + InstanceId
}
func getCreateHotGroupUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateInstanceWhiteGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_WHITELIST
}
func getCreateParameterTemplateUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateSyncGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + CONSTANT_CREATE
}
func getDeleteAccountUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACL_USER_ACTIONS + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteDeploymentSetUri(DeploySetId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DEPLOY_SET + bce.URI_PREFIX + DeploySetId
}
func getDeleteInstanceWhiteGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_WHITELIST
}
func getDeleteInstancesUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_DELETE__H_T_T_P + bce.URI_PREFIX + CONSTANT_1_1
}
func getDeleteIpWhitelistUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_IP
}
func getDeleteManualBackupUri(InstanceId string, BatchId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + BatchId
}
func getDeleteMemoryScalingConfigUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DELETE_AUTO_SCALING_CONFIG
}
func getDeleteParameterTemplateUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_DELETE + bce.URI_PREFIX + TemplateShowId
}
func getDeleteSyncGroupUri(version string, SyncGroupShowId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + SyncGroupShowId
}
func getDisconnectEntranceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENTRANCE + bce.URI_PREFIX + CONSTANT_DISCONNECT + bce.URI_PREFIX + InstanceId
}
func getDomainNameExchangeUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_SWAP_DOMAIN
}
func getGePriceForResizeInstanceUri(ClientToken string, InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PRICE
}
func getGetApplicationParameterTemplateRecordsUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + TemplateShowId
}
func getGetAvailableZonesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ZONE
}
func getGetBackUpUrlUri(InstanceId string, BackupId int32) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + BackupId + bce.URI_PREFIX + CONSTANT_URL
}
func getGetBackUpUsageUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_USAGE
}
func getGetBackupListUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP
}
func getGetBackupStrategyUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_POLICY
}
func getGetClusterBlbStatusUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BLB_STATUS
}
func getGetDeploymentSetListUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DEPLOY_SET
}
func getGetHotGroupDetailUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getGetHotGroupListUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetInstanceDetailUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getGetInstanceListUri(Marker string, MaxKeys string, InstanceIds string, VnetIp string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getGetInstanceSpecListUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_NODETYPES
}
func getGetInstanceWhiteGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_WHITELIST
}
func getGetParameterListUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PARAMETER
}
func getGetParameterTemplateListUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetPriceForCreateInstanceUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PRICE
}
func getGetRecycleListUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetSubnetListUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getGetSyncGroupStatusUri(version string, SyncGroupShowId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + SyncGroupShowId + bce.URI_PREFIX + CONSTANT_SYNC_STATUS
}
func getGetSystemParameterListUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_SYSTEM
}
func getGetTimeWindowUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TIME_WINDOW
}
func getGetTlsCertUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TLS
}
func getHotGroupAddClusterUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_JOIN
}
func getHotGroupChangeMasterRoleUri(GroupId string, InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_SET_AS_LEADER + bce.URI_PREFIX + InstanceId
}
func getHotGroupForbidWriteUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_FORBID_WRITE
}
func getHotGroupModifyNameUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getHotGroupPreCheckUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_CHECK
}
func getHotGroupRemoveClusterUri(GroupId string, InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_QUIT + bce.URI_PREFIX + InstanceId
}
func getHotGroupSetFlowControlRulesUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_QPS
}
func getHotGroupStaleReadableUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_STALE_READABLE
}
func getHotGroupSyncStatusUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_SYNC_STATUS
}
func getInstanceVersionUpgradeUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_UPGRADE_VERSION
}
func getLogDetailsUri(InstanceId string, LogId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_LOG + bce.URI_PREFIX + LogId
}
func getLogListUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_LOG + bce.URI_PREFIX + CONSTANT_LIST
}
func getManualBackupUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP
}
func getManuallyModifyBandwidthUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_BANDWIDTH
}
func getMasterSlaveSwitchUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SWITCH_MASTER_SLAVE
}
func getModifyBackupCommentUri(InstanceId string, BatchId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + BatchId + bce.URI_PREFIX + CONSTANT_COMMENT
}
func getModifyDeploymentSetUri(DeploySetId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DEPLOY_SET + bce.URI_PREFIX + DeploySetId
}
func getModifyEntranceUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AZONE_MIGRATION + bce.URI_PREFIX + CONSTANT_MODIFY_ENTRANCE
}
func getModifyInstanceDomainNameUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_RENAME_DOMAIN
}
func getModifyInstanceNameUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_RENAME
}
func getModifyParameterTemplateNameUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_RENAME + bce.URI_PREFIX + TemplateShowId
}
func getModifyParametersUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PARAMETER
}
func getModifyReplicationZoneUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AZONE_MIGRATION
}
func getModifySyncGroupNameUri(version string, SyncGroupShowId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + SyncGroupShowId
}
func getModifyTimeWindowUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TIME_WINDOW
}
func getParameterTemplateDeleteParametersUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_DELETE_PARAMS + bce.URI_PREFIX + TemplateShowId
}
func getParameterTemplateDetailsUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + TemplateShowId
}
func getParameterTemplateModifyParametersUri(TemplateShowId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_MODIFY_PARAMS + bce.URI_PREFIX + TemplateShowId
}
func getPostPaidToPrepaidUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TO_PREPAY
}
func getPrepaidToPostpaidUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TO_POSTPAY
}
func getProxyNodeReplaceUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROXY_NODE
}
func getProxyVersionUpgradeOrRestartUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_UPGRADE_PROXY
}
func getQueryIpWhitelistUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_IP
}
func getQueryMemoryScalingConfigUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AUTO_SCALING_CONFIG
}
func getReleaseHotGroupUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_RELEASE
}
func getReleaseInstanceUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_1_1
}
func getRenewInstanceUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RENEW
}
func getRestartInstanceUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_RESTART
}
func getSetBackupPolicyUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_MODIFY_POLICY
}
func getSetClusterAsMasterUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SET_AS_MASTER
}
func getSetClusterAsSlaveUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SET_AS_SLAVE
}
func getSetMemoryScalingConfigUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AUTO_SCALING_CONFIG
}
func getSetPermissionsUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_ACL_USER_ACTIONS + bce.URI_PREFIX + CONSTANT_AUTHORITY
}
func getSyncGroupAddInstanceUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_ADD_CLUSTER
}
func getSyncGroupDelayInfoUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_DELAY_INFO
}
func getSyncGroupDetailUri(version string, SyncGroupShowId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + SyncGroupShowId
}
func getSyncGroupListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + CONSTANT_LIST
}
func getSyncGroupModifyBnsgroupUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_MODIFY_BNS_GROUP
}
func getSyncGroupPreCheckUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + CONSTANT_CHECK
}
func getSyncGroupRemoveInstanceUri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SYNC_GROUP + bce.URI_PREFIX + GroupId + bce.URI_PREFIX + CONSTANT_REMOVE_CLUSTER
}
func getTdeEncryptionUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TDE
}
func getUnbindSecurityGroupUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_UNBIND
}
func getUnbindTagsUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_UN_BIND_TAG
}
func getUpdateInstanceWhiteGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_WHITELIST
}
func getUpdateSecurityGroupUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateTlsEncryptionUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TLS
}
func getViewSecurityGroupUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
}
