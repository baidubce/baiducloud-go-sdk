package vdb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "vdb." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VDB = "vdb"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_PRICE = "price"

	CONSTANT_ACCOUNT = "account"

	CONSTANT_RESET_PASSWORD = "resetPassword"

	CONSTANT_MODIFY_PUBLIC_ACCESS = "modifyPublicAccess"

	CONSTANT_DETAIL = "detail"

	CONSTANT_DESCRIBE_INSTANCE_CONFIGS = "describeInstanceConfigs"

	CONSTANT_RECYCLER = "recycler"

	CONSTANT_DELETE = "delete"

	CONSTANT_MODIFY_NAME = "modifyName"

	CONSTANT_MODIFY_INSTANCE_CONFIG = "modifyInstanceConfig"

	CONSTANT_LIST = "list"

	CONSTANT_PASSWORD = "password"

	CONSTANT_BACKUP = "backup"

	CONSTANT_SET_COMMENT = "setComment"

	CONSTANT_MODIFY_DOMAIN = "modifyDomain"

	CONSTANT_RECOVER = "recover"

	CONSTANT_BIND_EIP = "bindEip"

	CONSTANT_MANUAL_BACKUP = "manualBackup"

	CONSTANT_SECURITY = "security"

	CONSTANT_MODIFY_T_L_S = "modifyTLS"

	CONSTANT_CREATE = "create"

	CONSTANT_RESIZE = "resize"

	CONSTANT_GET_T_L_S_CERTIFICATE = "getTLSCertificate"

	CONSTANT_FREE_QUOTA = "freeQuota"

	CONSTANT_DELETE_RECORD = "deleteRecord"

	CONSTANT_LIST_RECORDS = "listRecords"

	CONSTANT_GET_CONFIG = "getConfig"

	CONSTANT_GET_NODE_SPEC_LIST = "getNodeSpecList"

	CONSTANT_UNBIND_EIP = "unbindEip"

	CONSTANT_QUOTA = "quota"

	CONSTANT_ZONE = "zone"

	CONSTANT_GET_T_L_S_INFO = "getTLSInfo"

	CONSTANT_SET_CONFIG = "setConfig"
)

// Client of vdb service is a kind of BceClient, so derived from BceClient
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

func getAccountListUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_LIST
}
func getBindEipUsingPOSTUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_BIND_EIP
}
func getCreateInstanceUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_CREATE
}
func getDeleteInstanceUsingDELETEUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteRecordUsingDELETEUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_DELETE_RECORD
}
func getDeleteRecyclerInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteinstanceusingdelete1Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDescribeInstanceConfigsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DESCRIBE_INSTANCE_CONFIGS
}
func getDescribeInstanceConfigsUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DESCRIBE_INSTANCE_CONFIGS
}
func getGetConfigUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_GET_CONFIG
}
func getGetFreeInstanceQuotaUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_FREE_QUOTA
}
func getGetFreeInstanceQuotaUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_FREE_QUOTA
}
func getGetInstanceListUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetNodeSpecListUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_GET_NODE_SPEC_LIST
}
func getGetPriceUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_PRICE
}
func getGetQuotaUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_QUOTA
}
func getGetTLSCertificateUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_GET_T_L_S_CERTIFICATE
}
func getGetTLSInfoUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_GET_T_L_S_INFO
}
func getGetinstancelistusingget1Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_LIST
}
func getInstanceDetailUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_DETAIL
}
func getListRecordsUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_LIST_RECORDS
}
func getManualBackupUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_MANUAL_BACKUP
}
func getModifyInstanceConfigUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_MODIFY_INSTANCE_CONFIG
}
func getModifyInstanceConfigUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_MODIFY_INSTANCE_CONFIG
}
func getModifyPasswordUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_RESET_PASSWORD
}
func getModifyPublicAccessUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_PUBLIC_ACCESS
}
func getModifyPublicAccessUsingPUTUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_PUBLIC_ACCESS
}
func getModifyTLSUsingPUTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_SECURITY + bce.URI_PREFIX + CONSTANT_MODIFY_T_L_S
}
func getPasswordUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_PASSWORD
}
func getRecoverInstanceUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_RECYCLER + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RECOVER
}
func getRecoverUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_RECOVER
}
func getResizeInstanceUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RESIZE
}
func getSetCommentUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_SET_COMMENT
}
func getSetConfigUsingPOSTUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_BACKUP + bce.URI_PREFIX + CONSTANT_SET_CONFIG
}
func getUnbindEipUsingPOSTUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_UNBIND_EIP
}
func getUpdateInstanceDomainUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_DOMAIN
}
func getUpdateInstanceDomainUsingPOSTUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_DOMAIN
}
func getUpdateInstanceNameUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_NAME
}
func getUpdateInstanceNameUsingPOSTUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MODIFY_NAME
}
func getZoneListUsingGETUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VDB + bce.URI_PREFIX + CONSTANT_ZONE
}
