package iam

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	IAM_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	IAM_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_AddUserToGroup(t *testing.T) {
	addUserToGroupRequest := &AddUserToGroupRequest{
		UserName:  util.PtrString(""),
		GroupName: util.PtrString(""),
	}
	err := IAM_CLIENT.AddUserToGroup(addUserToGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AssociateGroupPermissions(t *testing.T) {
	associateGroupPermissionsRequest := &AssociateGroupPermissionsRequest{
		GroupName:  util.PtrString(""),
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.AssociateGroupPermissions(associateGroupPermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AssociateRolePermissions(t *testing.T) {
	associateRolePermissionsRequest := &AssociateRolePermissionsRequest{
		RoleName:   util.PtrString(""),
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.AssociateRolePermissions(associateRolePermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AssociateUserPermissions(t *testing.T) {
	associateUserPermissionsRequest := &AssociateUserPermissionsRequest{
		UserName:   util.PtrString(""),
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.AssociateUserPermissions(associateUserPermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ChangeSubUserPassword(t *testing.T) {
	changeSubUserPasswordRequest := &ChangeSubUserPasswordRequest{
		UserName: util.PtrString(""),
		Password: util.PtrString(""),
	}
	err := IAM_CLIENT.ChangeSubUserPassword(changeSubUserPasswordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAccessKey(t *testing.T) {
	createAccessKeyRequest := &CreateAccessKeyRequest{
		UserName: util.PtrString(""),
	}
	err := IAM_CLIENT.CreateAccessKey(createAccessKeyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateApikeyPermanentlyValid(t *testing.T) {
	createApikeyPermanentlyValidRequest := &CreateApikeyPermanentlyValidRequest{
		UserId: util.PtrString(""),
		Acl:    util.PtrString(""),
		Name:   util.PtrString(""),
	}
	err := IAM_CLIENT.CreateApikeyPermanentlyValid(createApikeyPermanentlyValidRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateGroup(t *testing.T) {
	createGroupRequest := &CreateGroupRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := IAM_CLIENT.CreateGroup(createGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateRole(t *testing.T) {
	createRoleRequest := &CreateRoleRequest{
		Name:                     util.PtrString(""),
		Description:              util.PtrString(""),
		AssumeRolePolicyDocument: util.PtrString(""),
	}
	err := IAM_CLIENT.CreateRole(createRoleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateStrategy(t *testing.T) {
	createStrategyRequest := &CreateStrategyRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Document:    util.PtrString(""),
	}
	err := IAM_CLIENT.CreateStrategy(createStrategyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateUser(t *testing.T) {
	createUserRequest := &CreateUserRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := IAM_CLIENT.CreateUser(createUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DecodingApikeyPermanentlyValid(t *testing.T) {
	decodingApikeyPermanentlyValidRequest := &DecodingApikeyPermanentlyValidRequest{
		UserId: util.PtrString(""),
		Id:     util.PtrString(""),
	}
	err := IAM_CLIENT.DecodingApikeyPermanentlyValid(decodingApikeyPermanentlyValidRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAccessKey(t *testing.T) {
	deleteAccessKeyRequest := &DeleteAccessKeyRequest{
		UserName:    util.PtrString(""),
		AccessKeyId: util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteAccessKey(deleteAccessKeyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteApikeyPermanentlyValid(t *testing.T) {
	deleteApikeyPermanentlyValidRequest := &DeleteApikeyPermanentlyValidRequest{
		UserId: util.PtrString(""),
		Id:     util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteApikeyPermanentlyValid(deleteApikeyPermanentlyValidRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteGroup(t *testing.T) {
	deleteGroupRequest := &DeleteGroupRequest{
		GroupName: util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteGroup(deleteGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteLoginProfile(t *testing.T) {
	deleteLoginProfileRequest := &DeleteLoginProfileRequest{
		UserName: util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteLoginProfile(deleteLoginProfileRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRole(t *testing.T) {
	deleteRoleRequest := &DeleteRoleRequest{
		RoleName: util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteRole(deleteRoleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteStrategy(t *testing.T) {
	deleteStrategyRequest := &DeleteStrategyRequest{
		PolicyName: util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteStrategy(deleteStrategyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSubUserIdp(t *testing.T) {
	err := IAM_CLIENT.DeleteSubUserIdp()
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteUser(t *testing.T) {
	deleteUserRequest := &DeleteUserRequest{
		UserName: util.PtrString(""),
	}
	err := IAM_CLIENT.DeleteUser(deleteUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableAccessKey(t *testing.T) {
	disableAccessKeyRequest := &DisableAccessKeyRequest{
		UserName:    util.PtrString(""),
		AccessKeyId: util.PtrString(""),
	}
	err := IAM_CLIENT.DisableAccessKey(disableAccessKeyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableAccessKey(t *testing.T) {
	enableAccessKeyRequest := &EnableAccessKeyRequest{
		UserName:    util.PtrString(""),
		AccessKeyId: util.PtrString(""),
	}
	err := IAM_CLIENT.EnableAccessKey(enableAccessKeyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetLoginProfile(t *testing.T) {
	getLoginProfileRequest := &GetLoginProfileRequest{
		UserName: util.PtrString(""),
	}
	err := IAM_CLIENT.GetLoginProfile(getLoginProfileRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetSessionApiKey(t *testing.T) {
	getSessionApiKeyRequest := &GetSessionApiKeyRequest{
		ExpireInSeconds: util.PtrInt32(int32(0)),
		Acl:             util.PtrString(""),
	}
	err := IAM_CLIENT.GetSessionApiKey(getSessionApiKeyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetUser(t *testing.T) {
	getUserRequest := &GetUserRequest{
		UserName: util.PtrString(""),
	}
	err := IAM_CLIENT.GetUser(getUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAccessKey(t *testing.T) {
	listAccessKeyRequest := &ListAccessKeyRequest{
		UserName: util.PtrString(""),
	}
	result := &ListAccessKeyResponse{}
	result, err := IAM_CLIENT.ListAccessKey(listAccessKeyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAllSubjectsGrantedPermissions(t *testing.T) {
	listAllSubjectsGrantedPermissionsRequest := &ListAllSubjectsGrantedPermissionsRequest{
		PolicyId: util.PtrString(""),
	}
	result := &ListAllSubjectsGrantedPermissionsResponse{}
	result, err := IAM_CLIENT.ListAllSubjectsGrantedPermissions(listAllSubjectsGrantedPermissionsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListGroups(t *testing.T) {
	result := &ListGroupsResponse{}
	result, err := IAM_CLIENT.ListGroups()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListRoles(t *testing.T) {
	result := &ListRolesResponse{}
	result, err := IAM_CLIENT.ListRoles()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListStrategies(t *testing.T) {
	listStrategiesRequest := &ListStrategiesRequest{
		PolicyType: util.PtrString(""),
		NameFilter: util.PtrString(""),
	}
	result := &ListStrategiesResponse{}
	result, err := IAM_CLIENT.ListStrategies(listStrategiesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListThePermissionsOfRoles(t *testing.T) {
	listThePermissionsOfRolesRequest := &ListThePermissionsOfRolesRequest{
		RoleName: util.PtrString(""),
	}
	result := &ListThePermissionsOfRolesResponse{}
	result, err := IAM_CLIENT.ListThePermissionsOfRoles(listThePermissionsOfRolesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListThePermissionsOfTheGroup(t *testing.T) {
	listThePermissionsOfTheGroupRequest := &ListThePermissionsOfTheGroupRequest{
		GroupName: util.PtrString(""),
	}
	result := &ListThePermissionsOfTheGroupResponse{}
	result, err := IAM_CLIENT.ListThePermissionsOfTheGroup(listThePermissionsOfTheGroupRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTheSubjectsGrantedPermissions(t *testing.T) {
	listTheSubjectsGrantedPermissionsRequest := &ListTheSubjectsGrantedPermissionsRequest{
		PolicyId:  util.PtrString(""),
		GrantType: util.PtrString(""),
	}
	result := &ListTheSubjectsGrantedPermissionsResponse{}
	result, err := IAM_CLIENT.ListTheSubjectsGrantedPermissions(listTheSubjectsGrantedPermissionsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTheUserSPermissions(t *testing.T) {
	listTheUserSPermissionsRequest := &ListTheUserSPermissionsRequest{
		UserName: util.PtrString(""),
	}
	result := &ListTheUserSPermissionsResponse{}
	result, err := IAM_CLIENT.ListTheUserSPermissions(listTheUserSPermissionsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListUser(t *testing.T) {
	result := &ListUserResponse{}
	result, err := IAM_CLIENT.ListUser()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListUserGroups(t *testing.T) {
	listUserGroupsRequest := &ListUserGroupsRequest{
		UserName: util.PtrString(""),
	}
	result := &ListUserGroupsResponse{}
	result, err := IAM_CLIENT.ListUserGroups(listUserGroupsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListUsersWithinTheGroup(t *testing.T) {
	listUsersWithinTheGroupRequest := &ListUsersWithinTheGroupRequest{
		GroupName: util.PtrString(""),
	}
	result := &ListUsersWithinTheGroupResponse{}
	result, err := IAM_CLIENT.ListUsersWithinTheGroup(listUsersWithinTheGroupRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifySubUserOperationProtection(t *testing.T) {
	modifySubUserOperationProtectionRequest := &ModifySubUserOperationProtectionRequest{
		UserName:   util.PtrString(""),
		EnabledMfa: util.PtrBool(false),
		MfaType:    util.PtrString(""),
	}
	err := IAM_CLIENT.ModifySubUserOperationProtection(modifySubUserOperationProtectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ObtainAListOfPermanentlyValidApikeys(t *testing.T) {
	obtainAListOfPermanentlyValidApikeysRequest := &ObtainAListOfPermanentlyValidApikeysRequest{
		UserId:   util.PtrString(""),
		Service:  []*string{},
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	err := IAM_CLIENT.ObtainAListOfPermanentlyValidApikeys(obtainAListOfPermanentlyValidApikeysRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryApikeyDetailsPermanentlyValid(t *testing.T) {
	queryApikeyDetailsPermanentlyValidRequest := &QueryApikeyDetailsPermanentlyValidRequest{
		UserId: util.PtrString(""),
		Id:     util.PtrString(""),
	}
	err := IAM_CLIENT.QueryApikeyDetailsPermanentlyValid(queryApikeyDetailsPermanentlyValidRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryGroup(t *testing.T) {
	queryGroupRequest := &QueryGroupRequest{
		GroupName: util.PtrString(""),
	}
	err := IAM_CLIENT.QueryGroup(queryGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryRole(t *testing.T) {
	queryRoleRequest := &QueryRoleRequest{
		RoleName: util.PtrString(""),
	}
	err := IAM_CLIENT.QueryRole(queryRoleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryStrategy(t *testing.T) {
	queryStrategyRequest := &QueryStrategyRequest{
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.QueryStrategy(queryStrategyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySubUserIdp(t *testing.T) {
	result := &QuerySubUserIdpResponse{}
	result, err := IAM_CLIENT.QuerySubUserIdp()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySummaryOfMainAccount(t *testing.T) {
	result := &QuerySummaryOfMainAccountResponse{}
	result, err := IAM_CLIENT.QuerySummaryOfMainAccount()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTheLastUsageTimeOfAccesskey(t *testing.T) {
	queryTheLastUsageTimeOfAccesskeyRequest := &QueryTheLastUsageTimeOfAccesskeyRequest{
		AccessKeyId: util.PtrString(""),
	}
	result := &QueryTheLastUsageTimeOfAccesskeyResponse{}
	result, err := IAM_CLIENT.QueryTheLastUsageTimeOfAccesskey(queryTheLastUsageTimeOfAccesskeyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveGroupPermissions(t *testing.T) {
	removeGroupPermissionsRequest := &RemoveGroupPermissionsRequest{
		GroupName:  util.PtrString(""),
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.RemoveGroupPermissions(removeGroupPermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveRolePermissions(t *testing.T) {
	removeRolePermissionsRequest := &RemoveRolePermissionsRequest{
		RoleName:   util.PtrString(""),
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.RemoveRolePermissions(removeRolePermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveUserFromTheGroup(t *testing.T) {
	removeUserFromTheGroupRequest := &RemoveUserFromTheGroupRequest{
		UserName:  util.PtrString(""),
		GroupName: util.PtrString(""),
	}
	err := IAM_CLIENT.RemoveUserFromTheGroup(removeUserFromTheGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveUserPermissions(t *testing.T) {
	removeUserPermissionsRequest := &RemoveUserPermissionsRequest{
		UserName:   util.PtrString(""),
		PolicyName: util.PtrString(""),
		PolicyType: util.PtrString(""),
	}
	err := IAM_CLIENT.RemoveUserPermissions(removeUserPermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindSubUserVirtualMfa(t *testing.T) {
	unbindSubUserVirtualMfaRequest := &UnbindSubUserVirtualMfaRequest{
		UserName: util.PtrString(""),
		MfaType:  util.PtrString(""),
	}
	err := IAM_CLIENT.UnbindSubUserVirtualMfa(unbindSubUserVirtualMfaRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateApikeyPermanentlyValid(t *testing.T) {
	updateApikeyPermanentlyValidRequest := &UpdateApikeyPermanentlyValidRequest{
		UserId: util.PtrString(""),
		Id:     util.PtrString(""),
		Acl:    util.PtrString(""),
	}
	err := IAM_CLIENT.UpdateApikeyPermanentlyValid(updateApikeyPermanentlyValidRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateGroup(t *testing.T) {
	updateGroupRequest := &UpdateGroupRequest{
		GroupName:   util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := IAM_CLIENT.UpdateGroup(updateGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateLoginProfile(t *testing.T) {
	updateLoginProfileRequest := &UpdateLoginProfileRequest{
		UserName:          util.PtrString(""),
		Password:          util.PtrString(""),
		NeedResetPassword: util.PtrBool(false),
		EnabledLogin:      util.PtrBool(false),
		EnabledLoginMfa:   util.PtrBool(false),
		LoginMfaType:      util.PtrString(""),
		ThirdPartyType:    util.PtrString(""),
		ThirdPartyAccount: util.PtrString(""),
	}
	err := IAM_CLIENT.UpdateLoginProfile(updateLoginProfileRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRole(t *testing.T) {
	updateRoleRequest := &UpdateRoleRequest{
		RoleName:                 util.PtrString(""),
		Name:                     util.PtrString(""),
		Description:              util.PtrString(""),
		AssumeRolePolicyDocument: util.PtrString(""),
	}
	err := IAM_CLIENT.UpdateRole(updateRoleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateStrategy(t *testing.T) {
	updateStrategyRequest := &UpdateStrategyRequest{
		PolicyName:  util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Document:    util.PtrString(""),
	}
	err := IAM_CLIENT.UpdateStrategy(updateStrategyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSubUserIdp(t *testing.T) {
	updateSubUserIdpRequest := &UpdateSubUserIdpRequest{
		FileName:        util.PtrString(""),
		EncodeMetadata:  util.PtrString(""),
		AuxiliaryDomain: util.PtrString(""),
	}
	result := &UpdateSubUserIdpResponse{}
	result, err := IAM_CLIENT.UpdateSubUserIdp(updateSubUserIdpRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSubUserIdpStatus(t *testing.T) {
	updateSubUserIdpStatusRequest := &UpdateSubUserIdpStatusRequest{
		Status: util.PtrString(""),
	}
	result := &UpdateSubUserIdpStatusResponse{}
	result, err := IAM_CLIENT.UpdateSubUserIdpStatus(updateSubUserIdpStatusRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateUser(t *testing.T) {
	updateUserRequest := &UpdateUserRequest{
		UserName:    util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Enabled:     util.PtrBool(false),
	}
	err := IAM_CLIENT.UpdateUser(updateUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
