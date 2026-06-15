package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_ATTRIBUTE = "attribute"

	CONSTANT_GET_USERDATA = "getUserdata"

	CONSTANT_IMAGE = "image"

	CONSTANT_TAG = "tag"

	CONSTANT_VOLUME = "volume"

	CONSTANT_BATCH_DEL_IP = "batchDelIp"

	CONSTANT_PROGRESS = "progress"

	CONSTANT_DISK = "disk"

	CONSTANT_QUOTA = "quota"

	CONSTANT_BATCH = "batch"

	CONSTANT_CHARGING = "charging"

	CONSTANT_GET_PRICE = "getPrice"

	CONSTANT_RECOVERY = "recovery"

	CONSTANT_GET_AVAILABLE_IMAGES_BY_SPEC = "getAvailableImagesBySpec"

	CONSTANT_ROLE = "role"

	CONSTANT_RENAME = "rename"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_CHANGE_SUBNET = "changeSubnet"

	CONSTANT_BATCH_CREATE_AUTO_RENEW_RULES = "batchCreateAutoRenewRules"

	CONSTANT_VPC = "vpc"

	CONSTANT_CHANGE_VPC = "changeVpc"

	CONSTANT_BATCH_DELETE_AUTO_RENEW_RULES = "batchDeleteAutoRenewRules"

	CONSTANT_VNC = "vnc"

	CONSTANT_ADD_IPV6 = "addIpv6"

	CONSTANT_DELETION_PROTECTION = "deletionProtection"

	CONSTANT_RECYCLE = "recycle"

	CONSTANT_BATCH_ACTION = "batchAction"

	CONSTANT_SHARED_USERS = "sharedUsers"

	CONSTANT_NO_CHARGE = "noCharge"

	CONSTANT_REBUILD = "rebuild"

	CONSTANT_INSTANCE_BATCH_BY_SPEC = "instanceBatchBySpec"

	CONSTANT_MODIFY_RELATED_DELETE_POLICY = "modifyRelatedDeletePolicy"

	CONSTANT_DEPLOYSET = "deployset"

	CONSTANT_DEL_RELATION = "delRelation"

	CONSTANT_DEL_IPV6 = "delIpv6"

	CONSTANT_ENI = "eni"

	CONSTANT_LIST = "list"

	CONSTANT_BATCH_ADD_IP = "batchAddIp"

	CONSTANT_INSTANCE_BY_SPEC = "instanceBySpec"

	CONSTANT_BATCH_DELETE = "batchDelete"

	CONSTANT_BATCH_REFUND_RESOURCE = "batchRefundResource"

	CONSTANT_LIST_BY_INSTANCE_ID = "listByInstanceId"

	CONSTANT_IMPORT = "import"

	CONSTANT_OS = "os"

	CONSTANT_RESCUE = "rescue"

	CONSTANT_MODE = "mode"

	CONSTANT_EXIT = "exit"

	CONSTANT_ENTER = "enter"

	CONSTANT_DELETE = "delete"
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

func getAddIpv6Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ADD_IPV6
}
func getAttachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getAutoReleaseInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
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
func getBindInstanceToSecurityGroupUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getBindInstanceToTagsUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindRoleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ROLE
}
func getBindTagImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
func getCancelRemoteCopyImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getChangeToPrepaidUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getChangeVpcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VPC + bce.URI_PREFIX + CONSTANT_CHANGE_VPC
}
func getCreateAutoRenewRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_CREATE_AUTO_RENEW_RULES
}
func getCreateImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE
}
func getCreateInstanceBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE_BY_SPEC
}
func getCreateVolumeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getDelIpv6Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEL_IPV6
}
func getDeleteAutoRenewRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_BATCH_DELETE_AUTO_RENEW_RULES
}
func getDeleteImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getDeleteInstanceDeploySetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DEPLOYSET + bce.URI_PREFIX + CONSTANT_DEL_RELATION
}
func getDeletePrepayInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteRecycledInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RECYCLE + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getDetachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getEnterRescueModeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESCUE + bce.URI_PREFIX + CONSTANT_MODE + bce.URI_PREFIX + CONSTANT_ENTER
}
func getExitRescueModeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESCUE + bce.URI_PREFIX + CONSTANT_MODE + bce.URI_PREFIX + CONSTANT_EXIT
}
func getGetAvailableImagesBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_GET_AVAILABLE_IMAGES_BY_SPEC
}
func getGetCdsPriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_GET_PRICE
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
func getGetRoleListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getGetVolumeResizeProgressUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_PROGRESS + bce.URI_PREFIX + VolumeId
}
func getImportImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_IMPORT
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
func getListAvailableResizeSpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
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
func getListOsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_OS
}
func getListRecycleInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_RECYCLE + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListSharedUserUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_SHARED_USERS
}
func getListVolumesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getModifyCdsAttributeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
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
func getModifyVolumeChargeTypeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getPurchaseReservedInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getPurchaseReservedVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
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
func getRenameImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_RENAME
}
func getRenameVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getResizeInstanceBySpecUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE_BY_SPEC + bce.URI_PREFIX + InstanceId
}
func getResizeVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
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
func getUnbindRoleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_ROLE
}
func getUnbindTagImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUpdateInstanceSubnetUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET + bce.URI_PREFIX + CONSTANT_CHANGE_SUBNET
}
