package iam

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddUserToGroup
//
// PARAMS:
//   - request: the arguments to AddUserToGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAddUserToGroupUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.GroupName))).
		Do()
}

// AssociateGroupPermissions
//
// PARAMS:
//   - request: the arguments to AssociateGroupPermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AssociateGroupPermissions(request *AssociateGroupPermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAssociateGroupPermissionsUri(VERSION_V1, util.StringValue(request.GroupName), util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// AssociateRolePermissions
//
// PARAMS:
//   - request: the arguments to AssociateRolePermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AssociateRolePermissions(request *AssociateRolePermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAssociateRolePermissionsUri(VERSION_V1, util.StringValue(request.RoleName), util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// AssociateUserPermissions
//
// PARAMS:
//   - request: the arguments to AssociateUserPermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AssociateUserPermissions(request *AssociateUserPermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAssociateUserPermissionsUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// ChangeSubUserPassword
//
// PARAMS:
//   - request: the arguments to ChangeSubUserPassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ChangeSubUserPassword(request *ChangeSubUserPasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getChangeSubUserPasswordUri(VERSION_V1, util.StringValue(request.UserName))).
		WithBody(request).
		Do()
}

// CreateAccessKey
//
// PARAMS:
//   - request: the arguments to CreateAccessKey
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAccessKey(request *CreateAccessKeyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAccessKeyUri(VERSION_V1, util.StringValue(request.UserName))).
		Do()
}

// CreateApikeyPermanentlyValid
//
// PARAMS:
//   - request: the arguments to CreateApikeyPermanentlyValid
//
// RETURNS:
//   - CreateApikeyPermanentlyValidResponse: The return type of the CreateApikeyPermanentlyValid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateApikeyPermanentlyValid(request *CreateApikeyPermanentlyValidRequest) (*CreateApikeyPermanentlyValidResponse, error) {
	result := &CreateApikeyPermanentlyValidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateApikeyPermanentlyValidUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateGroup
//
// PARAMS:
//   - request: the arguments to CreateGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateGroup(request *CreateGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateGroupUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreateRole
//
// PARAMS:
//   - request: the arguments to CreateRole
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateRole(request *CreateRoleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRoleUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreateStrategy
//
// PARAMS:
//   - request: the arguments to CreateStrategy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateStrategy(request *CreateStrategyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateStrategyUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreateUser
//
// PARAMS:
//   - request: the arguments to CreateUser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateUser(request *CreateUserRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateUserUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// DecodingApikeyPermanentlyValid
//
// PARAMS:
//   - request: the arguments to DecodingApikeyPermanentlyValid
//
// RETURNS:
//   - DecodingApikeyPermanentlyValidResponse: The return type of the DecodingApikeyPermanentlyValid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DecodingApikeyPermanentlyValid(request *DecodingApikeyPermanentlyValidRequest) (*DecodingApikeyPermanentlyValidResponse, error) {
	result := &DecodingApikeyPermanentlyValidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDecodingApikeyPermanentlyValidUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAccessKey
//
// PARAMS:
//   - request: the arguments to DeleteAccessKey
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAccessKeyUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.AccessKeyId))).
		Do()
}

// DeleteApikeyPermanentlyValid
//
// PARAMS:
//   - request: the arguments to DeleteApikeyPermanentlyValid
//
// RETURNS:
//   - DeleteApikeyPermanentlyValidResponse: The return type of the DeleteApikeyPermanentlyValid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteApikeyPermanentlyValid(request *DeleteApikeyPermanentlyValidRequest) (*DeleteApikeyPermanentlyValidResponse, error) {
	result := &DeleteApikeyPermanentlyValidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteApikeyPermanentlyValidUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteGroup
//
// PARAMS:
//   - request: the arguments to DeleteGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteGroup(request *DeleteGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteGroupUri(VERSION_V1, util.StringValue(request.GroupName))).
		Do()
}

// DeleteLoginProfile
//
// PARAMS:
//   - request: the arguments to DeleteLoginProfile
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteLoginProfile(request *DeleteLoginProfileRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteLoginProfileUri(VERSION_V1, util.StringValue(request.UserName))).
		Do()
}

// DeleteRole
//
// PARAMS:
//   - request: the arguments to DeleteRole
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRole(request *DeleteRoleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRoleUri(VERSION_V1, util.StringValue(request.RoleName))).
		Do()
}

// DeleteStrategy
//
// PARAMS:
//   - request: the arguments to DeleteStrategy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteStrategy(request *DeleteStrategyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteStrategyUri(VERSION_V1, util.StringValue(request.PolicyName))).
		Do()
}

// DeleteSubUserIdp
//
// PARAMS:
//   - request: the arguments to DeleteSubUserIdp
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteSubUserIdp() error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteSubUserIdpUri(VERSION_V1)).
		Do()
}

// DeleteUser
//
// PARAMS:
//   - request: the arguments to DeleteUser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteUser(request *DeleteUserRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteUserUri(VERSION_V1, util.StringValue(request.UserName))).
		Do()
}

// DisableAccessKey
//
// PARAMS:
//   - request: the arguments to DisableAccessKey
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DisableAccessKey(request *DisableAccessKeyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getDisableAccessKeyUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.AccessKeyId))).
		WithQueryParam("disable", "").
		Do()
}

// EnableAccessKey
//
// PARAMS:
//   - request: the arguments to EnableAccessKey
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) EnableAccessKey(request *EnableAccessKeyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEnableAccessKeyUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.AccessKeyId))).
		WithQueryParam("enable", "").
		Do()
}

// GetLoginProfile
//
// PARAMS:
//   - request: the arguments to GetLoginProfile
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetLoginProfile(request *GetLoginProfileRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetLoginProfileUri(VERSION_V1, util.StringValue(request.UserName))).
		Do()
}

// GetSessionApiKey
//
// PARAMS:
//   - request: the arguments to GetSessionApiKey
//
// RETURNS:
//   - GetSessionApiKeyResponse: The return type of the GetSessionApiKey interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetSessionApiKey(request *GetSessionApiKeyRequest) (*GetSessionApiKeyResponse, error) {
	result := &GetSessionApiKeyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetSessionApiKeyUri(VERSION_V1)).
		WithQueryParamFilter("expireInSeconds", util.Int32Value(request.ExpireInSeconds)).
		WithQueryParamFilter("acl", util.StringValue(request.Acl)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUser
//
// PARAMS:
//   - request: the arguments to GetUser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetUser(request *GetUserRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetUserUri(VERSION_V1, util.StringValue(request.UserName))).
		Do()
}

// ListAccessKey
//
// PARAMS:
//   - request: the arguments to ListAccessKey
//
// RETURNS:
//   - ListAccessKeyResponse: The return type of the ListAccessKey interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAccessKey(request *ListAccessKeyRequest) (*ListAccessKeyResponse, error) {
	result := &ListAccessKeyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAccessKeyUri(VERSION_V1, util.StringValue(request.UserName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAllSubjectsGrantedPermissions
//
// PARAMS:
//   - request: the arguments to ListAllSubjectsGrantedPermissions
//
// RETURNS:
//   - ListAllSubjectsGrantedPermissionsResponse: The return type of the ListAllSubjectsGrantedPermissions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAllSubjectsGrantedPermissions(request *ListAllSubjectsGrantedPermissionsRequest) (*ListAllSubjectsGrantedPermissionsResponse, error) {
	result := &ListAllSubjectsGrantedPermissionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListAllSubjectsGrantedPermissionsUri(VERSION_V1, util.StringValue(request.PolicyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListGroups
//
// PARAMS:
//   - request: the arguments to ListGroups
//
// RETURNS:
//   - ListGroupsResponse: The return type of the ListGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListGroups() (*ListGroupsResponse, error) {
	result := &ListGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListGroupsUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRoles
//
// PARAMS:
//   - request: the arguments to ListRoles
//
// RETURNS:
//   - ListRolesResponse: The return type of the ListRoles interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRoles() (*ListRolesResponse, error) {
	result := &ListRolesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRolesUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListStrategies
//
// PARAMS:
//   - request: the arguments to ListStrategies
//
// RETURNS:
//   - ListStrategiesResponse: The return type of the ListStrategies interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListStrategies(request *ListStrategiesRequest) (*ListStrategiesResponse, error) {
	result := &ListStrategiesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListStrategiesUri(VERSION_V1)).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		WithQueryParamFilter("nameFilter", util.StringValue(request.NameFilter)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListThePermissionsOfRoles
//
// PARAMS:
//   - request: the arguments to ListThePermissionsOfRoles
//
// RETURNS:
//   - ListThePermissionsOfRolesResponse: The return type of the ListThePermissionsOfRoles interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListThePermissionsOfRoles(request *ListThePermissionsOfRolesRequest) (*ListThePermissionsOfRolesResponse, error) {
	result := &ListThePermissionsOfRolesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListThePermissionsOfRolesUri(VERSION_V1, util.StringValue(request.RoleName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListThePermissionsOfTheGroup
//
// PARAMS:
//   - request: the arguments to ListThePermissionsOfTheGroup
//
// RETURNS:
//   - ListThePermissionsOfTheGroupResponse: The return type of the ListThePermissionsOfTheGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListThePermissionsOfTheGroup(request *ListThePermissionsOfTheGroupRequest) (*ListThePermissionsOfTheGroupResponse, error) {
	result := &ListThePermissionsOfTheGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListThePermissionsOfTheGroupUri(VERSION_V1, util.StringValue(request.GroupName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTheSubjectsGrantedPermissions
//
// PARAMS:
//   - request: the arguments to ListTheSubjectsGrantedPermissions
//
// RETURNS:
//   - ListTheSubjectsGrantedPermissionsResponse: The return type of the ListTheSubjectsGrantedPermissions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTheSubjectsGrantedPermissions(request *ListTheSubjectsGrantedPermissionsRequest) (*ListTheSubjectsGrantedPermissionsResponse, error) {
	result := &ListTheSubjectsGrantedPermissionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTheSubjectsGrantedPermissionsUri(VERSION_V1, util.StringValue(request.PolicyId), util.StringValue(request.GrantType))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListTheUserSPermissions
//
// PARAMS:
//   - request: the arguments to ListTheUserSPermissions
//
// RETURNS:
//   - ListTheUserSPermissionsResponse: The return type of the ListTheUserSPermissions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListTheUserSPermissions(request *ListTheUserSPermissionsRequest) (*ListTheUserSPermissionsResponse, error) {
	result := &ListTheUserSPermissionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListTheUserSPermissionsUri(VERSION_V1, util.StringValue(request.UserName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListUser
//
// PARAMS:
//   - request: the arguments to ListUser
//
// RETURNS:
//   - ListUserResponse: The return type of the ListUser interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListUser() (*ListUserResponse, error) {
	result := &ListUserResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListUserUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListUserGroups
//
// PARAMS:
//   - request: the arguments to ListUserGroups
//
// RETURNS:
//   - ListUserGroupsResponse: The return type of the ListUserGroups interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListUserGroups(request *ListUserGroupsRequest) (*ListUserGroupsResponse, error) {
	result := &ListUserGroupsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListUserGroupsUri(VERSION_V1, util.StringValue(request.UserName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListUsersWithinTheGroup
//
// PARAMS:
//   - request: the arguments to ListUsersWithinTheGroup
//
// RETURNS:
//   - ListUsersWithinTheGroupResponse: The return type of the ListUsersWithinTheGroup interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListUsersWithinTheGroup(request *ListUsersWithinTheGroupRequest) (*ListUsersWithinTheGroupResponse, error) {
	result := &ListUsersWithinTheGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListUsersWithinTheGroupUri(VERSION_V1, util.StringValue(request.GroupName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ModifySubUserOperationProtection
//
// PARAMS:
//   - request: the arguments to ModifySubUserOperationProtection
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifySubUserOperationProtection(request *ModifySubUserOperationProtectionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifySubUserOperationProtectionUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// ObtainAListOfPermanentlyValidApikeys
//
// PARAMS:
//   - request: the arguments to ObtainAListOfPermanentlyValidApikeys
//
// RETURNS:
//   - ObtainAListOfPermanentlyValidApikeysResponse: The return type of the ObtainAListOfPermanentlyValidApikeys interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ObtainAListOfPermanentlyValidApikeys(request *ObtainAListOfPermanentlyValidApikeysRequest) (*ObtainAListOfPermanentlyValidApikeysResponse, error) {
	result := &ObtainAListOfPermanentlyValidApikeysResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getObtainAListOfPermanentlyValidApikeysUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryApikeyDetailsPermanentlyValid
//
// PARAMS:
//   - request: the arguments to QueryApikeyDetailsPermanentlyValid
//
// RETURNS:
//   - QueryApikeyDetailsPermanentlyValidResponse: The return type of the QueryApikeyDetailsPermanentlyValid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryApikeyDetailsPermanentlyValid(request *QueryApikeyDetailsPermanentlyValidRequest) (*QueryApikeyDetailsPermanentlyValidResponse, error) {
	result := &QueryApikeyDetailsPermanentlyValidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryApikeyDetailsPermanentlyValidUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryGroup
//
// PARAMS:
//   - request: the arguments to QueryGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) QueryGroup(request *QueryGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryGroupUri(VERSION_V1, util.StringValue(request.GroupName))).
		Do()
}

// QueryRole
//
// PARAMS:
//   - request: the arguments to QueryRole
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) QueryRole(request *QueryRoleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryRoleUri(VERSION_V1, util.StringValue(request.RoleName))).
		Do()
}

// QueryStrategy
//
// PARAMS:
//   - request: the arguments to QueryStrategy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) QueryStrategy(request *QueryStrategyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryStrategyUri(VERSION_V1, util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// QuerySubUserIdp
//
// PARAMS:
//   - request: the arguments to QuerySubUserIdp
//
// RETURNS:
//   - QuerySubUserIdpResponse: The return type of the QuerySubUserIdp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySubUserIdp() (*QuerySubUserIdpResponse, error) {
	result := &QuerySubUserIdpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySubUserIdpUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySummaryOfMainAccount
//
// PARAMS:
//   - request: the arguments to QuerySummaryOfMainAccount
//
// RETURNS:
//   - QuerySummaryOfMainAccountResponse: The return type of the QuerySummaryOfMainAccount interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySummaryOfMainAccount() (*QuerySummaryOfMainAccountResponse, error) {
	result := &QuerySummaryOfMainAccountResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySummaryOfMainAccountUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheLastUsageTimeOfAccesskey
//
// PARAMS:
//   - request: the arguments to QueryTheLastUsageTimeOfAccesskey
//
// RETURNS:
//   - QueryTheLastUsageTimeOfAccesskeyResponse: The return type of the QueryTheLastUsageTimeOfAccesskey interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheLastUsageTimeOfAccesskey(request *QueryTheLastUsageTimeOfAccesskeyRequest) (*QueryTheLastUsageTimeOfAccesskeyResponse, error) {
	result := &QueryTheLastUsageTimeOfAccesskeyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheLastUsageTimeOfAccesskeyUri(VERSION_V1, util.StringValue(request.AccessKeyId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveGroupPermissions
//
// PARAMS:
//   - request: the arguments to RemoveGroupPermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveGroupPermissions(request *RemoveGroupPermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveGroupPermissionsUri(VERSION_V1, util.StringValue(request.GroupName), util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// RemoveRolePermissions
//
// PARAMS:
//   - request: the arguments to RemoveRolePermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveRolePermissions(request *RemoveRolePermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveRolePermissionsUri(VERSION_V1, util.StringValue(request.RoleName), util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// RemoveUserFromTheGroup
//
// PARAMS:
//   - request: the arguments to RemoveUserFromTheGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveUserFromTheGroup(request *RemoveUserFromTheGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveUserFromTheGroupUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.GroupName))).
		Do()
}

// RemoveUserPermissions
//
// PARAMS:
//   - request: the arguments to RemoveUserPermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RemoveUserPermissions(request *RemoveUserPermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRemoveUserPermissionsUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.PolicyName))).
		WithQueryParamFilter("policyType", util.StringValue(request.PolicyType)).
		Do()
}

// UnbindSubUserVirtualMfa
//
// PARAMS:
//   - request: the arguments to UnbindSubUserVirtualMfa
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindSubUserVirtualMfa(request *UnbindSubUserVirtualMfaRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getUnbindSubUserVirtualMfaUri(VERSION_V1, util.StringValue(request.UserName), util.StringValue(request.MfaType))).
		Do()
}

// UpdateApikeyPermanentlyValid
//
// PARAMS:
//   - request: the arguments to UpdateApikeyPermanentlyValid
//
// RETURNS:
//   - UpdateApikeyPermanentlyValidResponse: The return type of the UpdateApikeyPermanentlyValid interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateApikeyPermanentlyValid(request *UpdateApikeyPermanentlyValidRequest) (*UpdateApikeyPermanentlyValidResponse, error) {
	result := &UpdateApikeyPermanentlyValidResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateApikeyPermanentlyValidUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateGroup
//
// PARAMS:
//   - request: the arguments to UpdateGroup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateGroup(request *UpdateGroupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateGroupUri(VERSION_V1, util.StringValue(request.GroupName))).
		WithBody(request).
		Do()
}

// UpdateLoginProfile
//
// PARAMS:
//   - request: the arguments to UpdateLoginProfile
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateLoginProfileUri(VERSION_V1, util.StringValue(request.UserName))).
		WithBody(request).
		Do()
}

// UpdateRole
//
// PARAMS:
//   - request: the arguments to UpdateRole
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateRole(request *UpdateRoleRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateRoleUri(VERSION_V1, util.StringValue(request.RoleName))).
		WithBody(request).
		Do()
}

// UpdateStrategy
//
// PARAMS:
//   - request: the arguments to UpdateStrategy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateStrategy(request *UpdateStrategyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateStrategyUri(VERSION_V1, util.StringValue(request.PolicyName))).
		WithBody(request).
		Do()
}

// UpdateSubUserIdp
//
// PARAMS:
//   - request: the arguments to UpdateSubUserIdp
//
// RETURNS:
//   - UpdateSubUserIdpResponse: The return type of the UpdateSubUserIdp interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSubUserIdp(request *UpdateSubUserIdpRequest) (*UpdateSubUserIdpResponse, error) {
	result := &UpdateSubUserIdpResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateSubUserIdpUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateSubUserIdpStatus
//
// PARAMS:
//   - request: the arguments to UpdateSubUserIdpStatus
//
// RETURNS:
//   - UpdateSubUserIdpStatusResponse: The return type of the UpdateSubUserIdpStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSubUserIdpStatus(request *UpdateSubUserIdpStatusRequest) (*UpdateSubUserIdpStatusResponse, error) {
	result := &UpdateSubUserIdpStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateSubUserIdpStatusUri(VERSION_V1)).
		WithQueryParamFilter("status", util.StringValue(request.Status)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser
//
// PARAMS:
//   - request: the arguments to UpdateUser
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateUser(request *UpdateUserRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateUserUri(VERSION_V1, util.StringValue(request.UserName))).
		WithBody(request).
		Do()
}
