package iam

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "iam." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_ACCOUNT = "account"

	CONSTANT_SUMMARY = "summary"

	CONSTANT_POLICY = "policy"

	CONSTANT_ROLE = "role"

	CONSTANT_APIKEY = "apikey"

	CONSTANT_CREATE = "create"

	CONSTANT_ENTITY = "entity"

	CONSTANT_SUB_USER = "subUser"

	CONSTANT_IDP = "idp"

	CONSTANT_UPDATE_STATUS = "updateStatus"

	CONSTANT_LIST = "list"

	CONSTANT_ACCESSKEY = "accesskey"

	CONSTANT_LASTUSEDTIME = "lastusedtime"

	CONSTANT_USER = "user"

	CONSTANT_LOGIN_PROFILE = "loginProfile"

	CONSTANT_B_C_E__B_E_A_R_E_R = "BCE-BEARER"

	CONSTANT_TOKEN = "token"

	CONSTANT_GROUP = "group"

	CONSTANT_OPERATION = "operation"

	CONSTANT_MFA = "mfa"

	CONSTANT_SWITCH = "switch"

	CONSTANT_DELETE = "delete"

	CONSTANT_QUERY = "query"

	CONSTANT_GRANT = "grant"

	CONSTANT_UPDATE = "update"

	CONSTANT_DETAIL = "detail"

	CONSTANT_MFA_TYPE = "mfaType"

	CONSTANT_DECRYPT = "decrypt"
)

// Client of iam service is a kind of BceClient, so derived from BceClient
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

func getAddUserToGroupUri(version string, UserName string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName
}
func getAssociateGroupPermissionsUri(version string, GroupName string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getAssociateRolePermissionsUri(version string, RoleName string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + RoleName + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getAssociateUserPermissionsUri(version string, UserName string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getChangeSubUserPasswordUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUB_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getCreateAccessKeyUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_ACCESSKEY
}
func getCreateApikeyPermanentlyValidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP
}
func getCreateRoleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE
}
func getCreateStrategyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY
}
func getCreateUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER
}
func getDecodingApikeyPermanentlyValidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_DECRYPT
}
func getDeleteAccessKeyUri(version string, UserName string, AccessKeyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_ACCESSKEY + bce.URI_PREFIX + AccessKeyId
}
func getDeleteApikeyPermanentlyValidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteGroupUri(version string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName
}
func getDeleteLoginProfileUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_LOGIN_PROFILE
}
func getDeleteRoleUri(version string, RoleName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + RoleName
}
func getDeleteStrategyUri(version string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getDeleteSubUserIdpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUB_USER + bce.URI_PREFIX + CONSTANT_IDP + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteUserUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName
}
func getDisableAccessKeyUri(version string, UserName string, AccessKeyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_ACCESSKEY + bce.URI_PREFIX + AccessKeyId
}
func getEnableAccessKeyUri(version string, UserName string, AccessKeyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_ACCESSKEY + bce.URI_PREFIX + AccessKeyId
}
func getGetLoginProfileUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_LOGIN_PROFILE
}
func getGetSessionApiKeyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_B_C_E__B_E_A_R_E_R + bce.URI_PREFIX + CONSTANT_TOKEN
}
func getGetUserUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName
}
func getListAccessKeyUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_ACCESSKEY
}
func getListAllSubjectsGrantedPermissionsUri(version string, PolicyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyId + bce.URI_PREFIX + CONSTANT_ENTITY
}
func getListGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP
}
func getListRolesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE
}
func getListStrategiesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY
}
func getListThePermissionsOfRolesUri(version string, RoleName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + RoleName + bce.URI_PREFIX + CONSTANT_POLICY
}
func getListThePermissionsOfTheGroupUri(version string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName + bce.URI_PREFIX + CONSTANT_POLICY
}
func getListTheSubjectsGrantedPermissionsUri(version string, PolicyId string, GrantType string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyId + bce.URI_PREFIX + CONSTANT_GRANT + bce.URI_PREFIX + GrantType
}
func getListTheUserSPermissionsUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_POLICY
}
func getListUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER
}
func getListUserGroupsUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_GROUP
}
func getListUsersWithinTheGroupUri(version string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName + bce.URI_PREFIX + CONSTANT_USER
}
func getModifySubUserOperationProtectionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_OPERATION + bce.URI_PREFIX + CONSTANT_MFA + bce.URI_PREFIX + CONSTANT_SWITCH
}
func getObtainAListOfPermanentlyValidApikeysUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_LIST
}
func getQueryApikeyDetailsPermanentlyValidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_DETAIL
}
func getQueryGroupUri(version string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName
}
func getQueryRoleUri(version string, RoleName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + RoleName
}
func getQueryStrategyUri(version string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getQuerySubUserIdpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUB_USER + bce.URI_PREFIX + CONSTANT_IDP + bce.URI_PREFIX + CONSTANT_QUERY
}
func getQuerySummaryOfMainAccountUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCOUNT + bce.URI_PREFIX + CONSTANT_SUMMARY
}
func getQueryTheLastUsageTimeOfAccesskeyUri(version string, AccessKeyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESSKEY + bce.URI_PREFIX + AccessKeyId + bce.URI_PREFIX + CONSTANT_LASTUSEDTIME
}
func getRemoveGroupPermissionsUri(version string, GroupName string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getRemoveRolePermissionsUri(version string, RoleName string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + RoleName + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getRemoveUserFromTheGroupUri(version string, UserName string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName
}
func getRemoveUserPermissionsUri(version string, UserName string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getUnbindSubUserVirtualMfaUri(version string, UserName string, MfaType string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_MFA_TYPE + bce.URI_PREFIX + MfaType
}
func getUpdateApikeyPermanentlyValidUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateGroupUri(version string, GroupName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupName
}
func getUpdateLoginProfileUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName + bce.URI_PREFIX + CONSTANT_LOGIN_PROFILE
}
func getUpdateRoleUri(version string, RoleName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ROLE + bce.URI_PREFIX + RoleName
}
func getUpdateStrategyUri(version string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + PolicyName
}
func getUpdateSubUserIdpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUB_USER + bce.URI_PREFIX + CONSTANT_IDP + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateSubUserIdpStatusUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUB_USER + bce.URI_PREFIX + CONSTANT_IDP + bce.URI_PREFIX + CONSTANT_UPDATE_STATUS
}
func getUpdateUserUri(version string, UserName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + UserName
}
