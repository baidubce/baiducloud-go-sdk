package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VOLUME = "volume"

	CONSTANT_CLUSTER = "cluster"

	CONSTANT_TAG = "tag"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_EHC = "ehc"

	CONSTANT_CREATE = "create"

	CONSTANT_IMAGE = "image"

	CONSTANT_REGION = "region"

	CONSTANT_DESCRIBE_REGIONS = "describeRegions"

	CONSTANT_SNAPSHOT = "snapshot"

	CONSTANT_DISK = "disk"

	CONSTANT_QUOTA = "quota"

	CONSTANT_RESERVED = "reserved"

	CONSTANT_GET_PRICE = "getPrice"

	CONSTANT_RECOVERY = "recovery"

	CONSTANT_DEPLOYSET = "deployset"

	CONSTANT_GET_AVAILABLE_IMAGES_BY_SPEC = "getAvailableImagesBySpec"

	CONSTANT_ROLE = "role"

	CONSTANT_MODIFY = "modify"

	CONSTANT_DELETE_PROTECTION = "deleteProtection"

	CONSTANT_BCC = "bcc"

	CONSTANT_SECURITY_GROUP = "securityGroup"

	CONSTANT_PRICE = "price"

	CONSTANT_BATCH_CREATE_AUTO_RENEW_RULES = "batchCreateAutoRenewRules"

	CONSTANT_TRANSFER = "transfer"

	CONSTANT_ASP = "asp"

	CONSTANT_UPDATE = "update"

	CONSTANT_CHAIN = "chain"

	CONSTANT_IN = "in"

	CONSTANT_LIST = "list"

	CONSTANT_AUTO_RENEW = "autoRenew"

	CONSTANT_KEYPAIR = "keypair"

	CONSTANT_VNC = "vnc"

	CONSTANT_UN_SHARE = "unShare"

	CONSTANT_RECYCLE = "recycle"

	CONSTANT_ACCEPT = "accept"

	CONSTANT_SHARED_USERS = "sharedUsers"

	CONSTANT_NO_CHARGE = "noCharge"

	CONSTANT_REBUILD = "rebuild"

	CONSTANT_INSTANCE_BATCH_BY_SPEC = "instanceBatchBySpec"

	CONSTANT_MODIFY_RELATED_DELETE_POLICY = "modifyRelatedDeletePolicy"

	CONSTANT_DEL_RELATION = "delRelation"

	CONSTANT_FLAVOR_SPEC = "flavorSpec"

	CONSTANT_REVOKE = "revoke"

	CONSTANT_CANCEL_BID_ORDER = "cancelBidOrder"

	CONSTANT_BID_FLAVOR = "bidFlavor"

	CONSTANT_INSTANCE_BY_SPEC = "instanceBySpec"

	CONSTANT_BATCH_DELETE = "batchDelete"

	CONSTANT_BATCH_REFUND_RESOURCE = "batchRefundResource"

	CONSTANT_REFUSE = "refuse"

	CONSTANT_LIST_BY_INSTANCE_ID = "listByInstanceId"

	CONSTANT_OS = "os"

	CONSTANT_SECURITYGROUP = "securitygroup"

	CONSTANT_UNBIND = "unbind"

	CONSTANT_BATCH_ACTION = "batchAction"

	CONSTANT_TASK = "task"

	CONSTANT_DETAIL = "detail"

	CONSTANT_BIND = "bind"

	CONSTANT_RULE = "rule"

	CONSTANT_ATTRIBUTE = "attribute"

	CONSTANT_GET_USERDATA = "getUserdata"

	CONSTANT_RESERVED_INSTANCE = "reservedInstance"

	CONSTANT_FLAVOR_ZONES = "flavorZones"

	CONSTANT_ZONE = "zone"

	CONSTANT_CANCEL_AUTO_RENEW = "cancelAutoRenew"

	CONSTANT_BATCH_DEL_IP = "batchDelIp"

	CONSTANT_PROGRESS = "progress"

	CONSTANT_BATCH = "batch"

	CONSTANT_CHARGING = "charging"

	CONSTANT_BID_PRICE = "bidPrice"

	CONSTANT_SNAPSHOT_SHARE = "snapshotShare"

	CONSTANT_RENAME = "rename"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_CHANGE_SUBNET = "changeSubnet"

	CONSTANT_REPLACE = "replace"

	CONSTANT_VPC = "vpc"

	CONSTANT_CHANGE_VPC = "changeVpc"

	CONSTANT_BATCH_DELETE_AUTO_RENEW_RULES = "batchDeleteAutoRenewRules"

	CONSTANT_DELETE = "delete"

	CONSTANT_ADD_IPV6 = "addIpv6"

	CONSTANT_DELETION_PROTECTION = "deletionProtection"

	CONSTANT_OUT = "out"

	CONSTANT_SHARE = "share"

	CONSTANT_UPDATE_RELATION = "updateRelation"

	CONSTANT_DEL_IPV6 = "delIpv6"

	CONSTANT_ENI = "eni"

	CONSTANT_BATCH_ADD_IP = "batchAddIp"

	CONSTANT_REMOTE_COPY = "remote_copy"

	CONSTANT_IMPORT = "import"

	CONSTANT_RESCUE = "rescue"

	CONSTANT_MODE = "mode"

	CONSTANT_EXIT = "exit"

	CONSTANT_RENEW = "renew"

	CONSTANT_ENTER = "enter"
)

// Client of bcc service is a kind of BceClient, so derived from BceClient
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

func getAcceptReservedInstanceTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TRANSFER + bce.URI_PREFIX + CONSTANT_ACCEPT
}
func getAddIpv6Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ADD_IPV6
}
func getAttachAspUri(version string, AspId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP + bce.URI_PREFIX + AspId
}
func getAttachKeypairUri(version string, KeypairId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR + bce.URI_PREFIX + KeypairId
}
func getAttachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getAuthorizeSecurityGroupRuleUri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getAuthorizeServerEventUri() string {
	return bce.URI_PREFIX
}
func getAutoReleaseInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getAutoRenewReservedInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_AUTO_RENEW
}
func getAutoRenewVolumeClusterUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_AUTO_RENEW
}
func getBatchAddIpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_ADD_IP
}
func getBatchChangeToPostpaidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH + bce.URI_PREFIX + CONSTANT_CHARGING
}
func getBatchChangeToPrepaidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH + bce.URI_PREFIX + CONSTANT_CHARGING
}
func getBatchDeleteIpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_DEL_IP
}
func getBatchRefundResourceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_REFUND_RESOURCE
}
func getBatchStartInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_ACTION
}
func getBatchStopInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_ACTION
}
func getBindInstanceSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITYGROUP + bce.URI_PREFIX + CONSTANT_BIND
}
func getBindInstanceToSecurityGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getBindInstanceToTagsUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindReservedInstanceToTagsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BCC + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindRoleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ROLE
}
func getBindTagImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindTagSnapchainUri(version string, ChainId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_CHAIN + bce.URI_PREFIX + ChainId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindTagVolumeClusterUri(version string, ClusterId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterId + bce.URI_PREFIX + CONSTANT_TAG
}
func getCancelAutoRenewReservedInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_CANCEL_AUTO_RENEW
}
func getCancelAutoRenewVolumeClusterUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_CANCEL_AUTO_RENEW
}
func getCancelBidOrderUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_CANCEL_BID_ORDER
}
func getCancelRemoteCopyImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getCancelSnapshotShareUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_UN_SHARE
}
func getChangeToPrepaidUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getChangeVpcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_CHANGE_VPC
}
func getCheckServerEventUri() string {
	return bce.URI_PREFIX
}
func getCreateAspUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP
}
func getCreateAuthorizationRuleUri() string {
	return bce.URI_PREFIX
}
func getCreateAutoRenewRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_CREATE_AUTO_RENEW_RULES
}
func getCreateBidInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE_BY_SPEC
}
func getCreateDeploySetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateEhcClusterUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_EHC + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE
}
func getCreateInstanceBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE_BY_SPEC
}
func getCreateKeypairUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR
}
func getCreateReservedInstanceTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TRANSFER + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateReservedInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
}
func getCreateSnapshotUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT
}
func getCreateSnapshotShareUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_SHARE
}
func getCreateVolumeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getCreateVolumeClusterUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER
}
func getDelIpv6Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEL_IPV6
}
func getDeleteAspUri(version string, AspId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP + bce.URI_PREFIX + AspId
}
func getDeleteAutoRenewRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_DELETE_AUTO_RENEW_RULES
}
func getDeleteDeploySetUri(version string, DeployId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + DeployId
}
func getDeleteEhcClusterUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_EHC + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getDeleteInstUserOpAuthorizeRuleUri() string {
	return bce.URI_PREFIX
}
func getDeleteInstanceDeploySetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + CONSTANT_DEL_RELATION
}
func getDeleteKeypairUri(version string, KeypairId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR + bce.URI_PREFIX + KeypairId
}
func getDeletePrepayInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteRecycledInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RECYCLE + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getDeleteSecurityGroupUri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getDeleteSecurityGroupRuleUri(version string, SecurityGroupRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + SecurityGroupRuleId
}
func getDeleteSnapshotUri(version string, SnapshotId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + SnapshotId
}
func getDeletesInstanceDeploySetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + CONSTANT_DEL_RELATION
}
func getDescribeAuthorizeRulesUri() string {
	return bce.URI_PREFIX
}
func getDescribePlannedEventRecordsUri() string {
	return bce.URI_PREFIX
}
func getDescribePlannedEventsUri() string {
	return bce.URI_PREFIX
}
func getDescribeRegionsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_REGION + bce.URI_PREFIX + CONSTANT_DESCRIBE_REGIONS
}
func getDescribeSnapshotsUsageUri(version string, Action string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + Action
}
func getDescribeUnplannedEventRecordsUri() string {
	return bce.URI_PREFIX
}
func getDescribeUnplannedEventsUri() string {
	return bce.URI_PREFIX
}
func getDetachAspUri(version string, AspId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP + bce.URI_PREFIX + AspId
}
func getDetachKeypairUri(version string, KeypairId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR + bce.URI_PREFIX + KeypairId
}
func getDetachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getEhcClusterListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_EHC + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_LIST
}
func getEnterRescueModeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESCUE + bce.URI_PREFIX + CONSTANT_MODE + bce.URI_PREFIX + CONSTANT_ENTER
}
func getExitRescueModeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESCUE + bce.URI_PREFIX + CONSTANT_MODE + bce.URI_PREFIX + CONSTANT_EXIT
}
func getGetAspUri(version string, AspId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP + bce.URI_PREFIX + AspId
}
func getGetAvailableImagesBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_GET_AVAILABLE_IMAGES_BY_SPEC
}
func getGetBidInstancePriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BID_PRICE
}
func getGetCdsPriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_GET_PRICE
}
func getGetDeploySetUri(version string, Id string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + Id
}
func getGetDiskQuotaUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_DISK + bce.URI_PREFIX + CONSTANT_QUOTA
}
func getGetImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getGetInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getGetInstanceNoChargeListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_NO_CHARGE
}
func getGetInstanceUserDataInfoUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ATTRIBUTE + bce.URI_PREFIX + CONSTANT_GET_USERDATA
}
func getGetInstanceVncUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_VNC
}
func getGetPriceBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_PRICE
}
func getGetReservedInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetReservedInstancePriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RESERVED_INSTANCE + bce.URI_PREFIX + CONSTANT_PRICE
}
func getGetRoleListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetSnapshotUri(version string, SnapshotId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + SnapshotId
}
func getGetTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_DETAIL
}
func getGetVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getGetVolumeClusterUri(version string, ClusterId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterId
}
func getGetVolumeResizeProgressUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_PROGRESS + bce.URI_PREFIX + VolumeId
}
func getGetZoneBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_FLAVOR_ZONES
}
func getImportImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_IMPORT
}
func getImportKeypairUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR
}
func getInstanceBatchResizeBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE_BATCH_BY_SPEC
}
func getInstanceDeletionProtectionUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DELETION_PROTECTION
}
func getInstanceRecoveryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RECOVERY
}
func getKeypairDetailUri(version string, KeypairId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR + bce.URI_PREFIX + KeypairId
}
func getListAspsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP
}
func getListAvailableResizeSpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListBidFlavorUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BID_FLAVOR
}
func getListDeploySetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + CONSTANT_LIST
}
func getListFlavorSpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_FLAVOR_SPEC
}
func getListImagesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE
}
func getListInstanceByIdsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_LIST_BY_INSTANCE_ID
}
func getListInstanceEnisUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ENI + bce.URI_PREFIX + InstanceId
}
func getListInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListKeypairUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR
}
func getListOsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_OS
}
func getListRecycleInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RECYCLE + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListReservedInstanceTransferInUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TRANSFER + bce.URI_PREFIX + CONSTANT_IN + bce.URI_PREFIX + CONSTANT_LIST
}
func getListReservedInstanceTransferOutUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TRANSFER + bce.URI_PREFIX + CONSTANT_OUT + bce.URI_PREFIX + CONSTANT_LIST
}
func getListSecurityGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP
}
func getListSharedUserUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_SHARED_USERS
}
func getListSnapchainUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_CHAIN
}
func getListSnapshotShareUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_SNAPSHOT_SHARE + bce.URI_PREFIX + CONSTANT_LIST
}
func getListSnapshotsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT
}
func getListTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_LIST
}
func getListVolumeClustersUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER
}
func getListVolumesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getListZonesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ZONE
}
func getModifyCdsAttributeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getModifyEhcClusterUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_EHC + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_MODIFY
}
func getModifyInstUserOpAuthorizeRuleAttributeUri() string {
	return bce.URI_PREFIX
}
func getModifyInstanceAttributesUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getModifyInstanceDescUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getModifyInstanceHostnameUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getModifyInstancePasswordUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getModifyRelatedDeletePolicyUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_RELATED_DELETE_POLICY
}
func getModifyReservedInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_MODIFY
}
func getModifySnapshotAttributeUri(version string, SnapshotId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + SnapshotId
}
func getModifyVolumeChargeTypeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getModifyVolumeDeleteProtectionV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_MODIFY + bce.URI_PREFIX + CONSTANT_DELETE_PROTECTION
}
func getPurchaseReservedInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getPurchaseReservedVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getPurchaseReservedVolumeClusterUri(version string, ClusterId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterId
}
func getRebootInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getRebuildBatchInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_REBUILD
}
func getRebuildInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getRefuseReservedInstanceTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TRANSFER + bce.URI_PREFIX + CONSTANT_REFUSE
}
func getReleaseInstanceByPostUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getReleaseMultipleInstanceByPostUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_DELETE
}
func getReleaseVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getRemoteCopyImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getRemoteCopySnapshotUri(version string, SnapshotId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_REMOTE_COPY + bce.URI_PREFIX + SnapshotId
}
func getRenameImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_RENAME
}
func getRenameKeypairUri(version string, KeypairId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR + bce.URI_PREFIX + KeypairId
}
func getRenameVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getRenewReservedInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_RENEW
}
func getReplaceInstanceSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITYGROUP + bce.URI_PREFIX + CONSTANT_REPLACE
}
func getResizeInstanceBySpecUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE_BY_SPEC + bce.URI_PREFIX + InstanceId
}
func getResizeVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getResizeVolumeClusterUri(version string, ClusterId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterId
}
func getRevokeReservedInstanceTransferUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TRANSFER + bce.URI_PREFIX + CONSTANT_REVOKE
}
func getRevokeSecurityGroupRuleUri(version string, SecurityGroupId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + SecurityGroupId
}
func getRollbackVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getShareImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getStartInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getStopInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUnShareImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getUnbindInstanceFromSecurityGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUnbindInstanceFromTagsUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindInstanceSecurityGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITYGROUP + bce.URI_PREFIX + CONSTANT_UNBIND
}
func getUnbindReservedInstanceFromTagsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BCC + bce.URI_PREFIX + CONSTANT_RESERVED + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindRoleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ROLE
}
func getUnbindTagImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindTagSnapchainUri(version string, ChainId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SNAPSHOT + bce.URI_PREFIX + CONSTANT_CHAIN + bce.URI_PREFIX + ChainId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindTagVolumeClusterUri(version string, ClusterId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUpdateAspUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ASP + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateDeploySetUri(version string, DeployId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + DeployId
}
func getUpdateDeploySetRelationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + CONSTANT_UPDATE_RELATION
}
func getUpdateInstanceSubnetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_CHANGE_SUBNET
}
func getUpdateKeypairDescriptionUri(version string, KeypairId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_KEYPAIR + bce.URI_PREFIX + KeypairId
}
func getUpdateSecurityGroupRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SECURITY_GROUP + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + CONSTANT_UPDATE
}
