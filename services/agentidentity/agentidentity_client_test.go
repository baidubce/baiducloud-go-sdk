package agentidentity

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
	AGENTIDENTITY_CLIENT *Client
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

	AGENTIDENTITY_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AuthorizeEndpoint(t *testing.T) {
	authorizeEndpointRequest := &AuthorizeEndpointRequest{
		UserPoolId:   util.PtrString(""),
		ClientId:     util.PtrString(""),
		RedirectUri:  util.PtrString(""),
		ResponseType: util.PtrString(""),
		Scope:        util.PtrString(""),
		State:        util.PtrString(""),
		Nonce:        util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.AuthorizeEndpoint(authorizeEndpointRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchAcquisitionOfUsers(t *testing.T) {
	batchAcquisitionOfUsersRequest := &BatchAcquisitionOfUsersRequest{
		UserPoolId: util.PtrString(""),
		Ids:        []*string{},
	}
	result := &BatchAcquisitionOfUsersResponse{}
	result, err := AGENTIDENTITY_CLIENT.BatchAcquisitionOfUsers(batchAcquisitionOfUsersRequest)
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
func TestClient_BatchGetResourceApiKey(t *testing.T) {
	batchGetResourceApiKeyRequest := &BatchGetResourceApiKeyRequest{
		XBceWorkloadAccessToken: util.PtrString(""),
		Names:                   []*string{},
		WorkloadAccessToken:     util.PtrString(""),
	}
	result := &BatchGetResourceApiKeyResponse{}
	result, err := AGENTIDENTITY_CLIENT.BatchGetResourceApiKey(batchGetResourceApiKeyRequest)
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
func TestClient_CompleteOauth2session(t *testing.T) {
	completeOauth2sessionRequest := &CompleteOauth2sessionRequest{
		XBceWorkloadAccessToken: util.PtrString(""),
		SessionUri:              util.PtrString(""),
		UserIdentifier:          nil,
		UserIdentifierUserId:    util.PtrString(""),
		WorkloadAccessToken:     util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.CompleteOauth2session(completeOauth2sessionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAgent(t *testing.T) {
	createAgentRequest := &CreateAgentRequest{
		Name:                            util.PtrString(""),
		Description:                     util.PtrString(""),
		AllowedResourceOauth2ReturnUrls: []*string{},
	}
	result := &CreateAgentResponse{}
	result, err := AGENTIDENTITY_CLIENT.CreateAgent(createAgentRequest)
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
func TestClient_CreateCredentialProvider(t *testing.T) {
	createCredentialProviderRequest := &CreateCredentialProviderRequest{
		Name:              util.PtrString(""),
		AgentidentityType: util.PtrString(""),
		Desc:              util.PtrString(""),
		Credential:        nil,
	}
	result := &CreateCredentialProviderResponse{}
	result, err := AGENTIDENTITY_CLIENT.CreateCredentialProvider(createCredentialProviderRequest)
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
func TestClient_CreateIdpConfiguration(t *testing.T) {
	createIdpConfigurationRequest := &CreateIdpConfigurationRequest{
		UserPoolId:            util.PtrString(""),
		Name:                  util.PtrString(""),
		IdpType:               util.PtrString(""),
		IdpProvider:           util.PtrString(""),
		ClientId:              util.PtrString(""),
		ClientSecret:          util.PtrString(""),
		DiscoveryUrl:          util.PtrString(""),
		AuthorizationEndpoint: util.PtrString(""),
		TokenEndpoint:         util.PtrString(""),
		UserinfoEndpoint:      util.PtrString(""),
		Scopes:                []*string{},
		UserIdClaim:           util.PtrString(""),
		DisplayNameClaim:      util.PtrString(""),
		AutoCreateUser:        util.PtrBool(false),
	}
	result := &CreateIdpConfigurationResponse{}
	result, err := AGENTIDENTITY_CLIENT.CreateIdpConfiguration(createIdpConfigurationRequest)
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
func TestClient_CreateOauth2Client(t *testing.T) {
	createOauth2ClientRequest := &CreateOauth2ClientRequest{
		UserPoolId:      util.PtrString(""),
		Name:            util.PtrString(""),
		Description:     util.PtrString(""),
		ClientType:      util.PtrString(""),
		RedirectUris:    []*string{},
		GrantTypes:      []*string{},
		Scopes:          []*string{},
		AccessTokenTtl:  util.PtrInt32(int32(0)),
		RefreshTokenTtl: util.PtrInt32(int32(0)),
	}
	result := &CreateOauth2ClientResponse{}
	result, err := AGENTIDENTITY_CLIENT.CreateOauth2Client(createOauth2ClientRequest)
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
func TestClient_CreateUser(t *testing.T) {
	createUserRequest := &CreateUserRequest{
		UserPoolId:         util.PtrString(""),
		Username:           util.PtrString(""),
		DisplayName:        util.PtrString(""),
		Description:        util.PtrString(""),
		Password:           util.PtrString(""),
		ForceResetPassword: util.PtrBool(false),
	}
	result := &CreateUserResponse{}
	result, err := AGENTIDENTITY_CLIENT.CreateUser(createUserRequest)
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
func TestClient_CreateUserPool(t *testing.T) {
	createUserPoolRequest := &CreateUserPoolRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateUserPoolResponse{}
	result, err := AGENTIDENTITY_CLIENT.CreateUserPool(createUserPoolRequest)
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
func TestClient_DeleteAgent(t *testing.T) {
	deleteAgentRequest := &DeleteAgentRequest{
		AgentId: util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.DeleteAgent(deleteAgentRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCredentialProvider(t *testing.T) {
	deleteCredentialProviderRequest := &DeleteCredentialProviderRequest{
		CredentialProviderId: util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.DeleteCredentialProvider(deleteCredentialProviderRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIdpConfiguration(t *testing.T) {
	deleteIdpConfigurationRequest := &DeleteIdpConfigurationRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.DeleteIdpConfiguration(deleteIdpConfigurationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteOauth2Client(t *testing.T) {
	deleteOauth2ClientRequest := &DeleteOauth2ClientRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.DeleteOauth2Client(deleteOauth2ClientRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteUser(t *testing.T) {
	deleteUserRequest := &DeleteUserRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.DeleteUser(deleteUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteUserPool(t *testing.T) {
	deleteUserPoolRequest := &DeleteUserPoolRequest{
		Id: util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.DeleteUserPool(deleteUserPoolRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableIdpConfiguration(t *testing.T) {
	disableIdpConfigurationRequest := &DisableIdpConfigurationRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	result := &DisableIdpConfigurationResponse{}
	result, err := AGENTIDENTITY_CLIENT.DisableIdpConfiguration(disableIdpConfigurationRequest)
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
func TestClient_EnableIdpConfiguration(t *testing.T) {
	enableIdpConfigurationRequest := &EnableIdpConfigurationRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	result := &EnableIdpConfigurationResponse{}
	result, err := AGENTIDENTITY_CLIENT.EnableIdpConfiguration(enableIdpConfigurationRequest)
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
func TestClient_GetAgent(t *testing.T) {
	getAgentRequest := &GetAgentRequest{
		AgentId: util.PtrString(""),
	}
	result := &GetAgentResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetAgent(getAgentRequest)
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
func TestClient_GetCredentialProvider(t *testing.T) {
	getCredentialProviderRequest := &GetCredentialProviderRequest{
		CredentialProviderId: util.PtrString(""),
	}
	result := &GetCredentialProviderResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetCredentialProvider(getCredentialProviderRequest)
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
func TestClient_GetIdpConfiguration(t *testing.T) {
	getIdpConfigurationRequest := &GetIdpConfigurationRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	result := &GetIdpConfigurationResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetIdpConfiguration(getIdpConfigurationRequest)
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
func TestClient_GetOauth2Client(t *testing.T) {
	getOauth2ClientRequest := &GetOauth2ClientRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
	}
	result := &GetOauth2ClientResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetOauth2Client(getOauth2ClientRequest)
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
func TestClient_GetResourceApikey(t *testing.T) {
	getResourceApikeyRequest := &GetResourceApikeyRequest{
		XBceWorkloadAccessToken: util.PtrString(""),
		Name:                    util.PtrString(""),
		WorkloadAccessToken:     util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.GetResourceApikey(getResourceApikeyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetResourceOauth2token(t *testing.T) {
	getResourceOauth2tokenRequest := &GetResourceOauth2tokenRequest{
		XBceWorkloadAccessToken:        util.PtrString(""),
		ResourceCredentialProviderName: util.PtrString(""),
		Scopes:                         []*string{},
		Oauth2Flow:                     util.PtrString(""),
		ResourceOauth2ReturnUrl:        util.PtrString(""),
		SessionUri:                     util.PtrString(""),
		ForceAuthentication:            util.PtrBool(false),
		WorkloadAccessToken:            util.PtrString(""),
	}
	result := &GetResourceOauth2tokenResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetResourceOauth2token(getResourceOauth2tokenRequest)
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
func TestClient_GetUser(t *testing.T) {
	getUserRequest := &GetUserRequest{
		UserPoolId: util.PtrString(""),
		Id:         util.PtrString(""),
		Username:   util.PtrString(""),
	}
	result := &GetUserResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetUser(getUserRequest)
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
func TestClient_GetUserPool(t *testing.T) {
	getUserPoolRequest := &GetUserPoolRequest{
		Id: util.PtrString(""),
	}
	result := &GetUserPoolResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetUserPool(getUserPoolRequest)
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
func TestClient_GetWATForUser(t *testing.T) {
	getWATForUserRequest := &GetWATForUserRequest{
		BceUserId:       util.PtrString(""),
		AgentId:         util.PtrString(""),
		AgentName:       util.PtrString(""),
		UserId:          util.PtrString(""),
		SessionId:       util.PtrString(""),
		DurationSeconds: util.PtrInt32(int32(0)),
	}
	result := &GetWATForUserResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetWATForUser(getWATForUserRequest)
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
func TestClient_GetWorkloadAccessToken(t *testing.T) {
	getWorkloadAccessTokenRequest := &GetWorkloadAccessTokenRequest{
		BceUserId:       util.PtrString(""),
		AgentId:         util.PtrString(""),
		AgentName:       util.PtrString(""),
		DurationSeconds: util.PtrInt32(int32(0)),
	}
	result := &GetWorkloadAccessTokenResponse{}
	result, err := AGENTIDENTITY_CLIENT.GetWorkloadAccessToken(getWorkloadAccessTokenRequest)
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
func TestClient_ListAgents(t *testing.T) {
	listAgentsRequest := &ListAgentsRequest{
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
		Keyword:  util.PtrString(""),
	}
	result := &ListAgentsResponse{}
	result, err := AGENTIDENTITY_CLIENT.ListAgents(listAgentsRequest)
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
func TestClient_ListCredentialProviders(t *testing.T) {
	listCredentialProvidersRequest := &ListCredentialProvidersRequest{
		PageNo:            util.PtrInt32(int32(0)),
		PageSize:          util.PtrInt32(int32(0)),
		AgentidentityType: util.PtrString(""),
		Name:              util.PtrString(""),
	}
	result := &ListCredentialProvidersResponse{}
	result, err := AGENTIDENTITY_CLIENT.ListCredentialProviders(listCredentialProvidersRequest)
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
func TestClient_ListIdpConfigurations(t *testing.T) {
	listIdpConfigurationsRequest := &ListIdpConfigurationsRequest{
		UserPoolId: util.PtrString(""),
		Keyword:    util.PtrString(""),
		PageNo:     util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &ListIdpConfigurationsResponse{}
	result, err := AGENTIDENTITY_CLIENT.ListIdpConfigurations(listIdpConfigurationsRequest)
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
func TestClient_ListOauth2Clients(t *testing.T) {
	listOauth2ClientsRequest := &ListOauth2ClientsRequest{
		UserPoolId: util.PtrString(""),
		Keyword:    util.PtrString(""),
		PageNo:     util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &ListOauth2ClientsResponse{}
	result, err := AGENTIDENTITY_CLIENT.ListOauth2Clients(listOauth2ClientsRequest)
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
func TestClient_ListUserPools(t *testing.T) {
	listUserPoolsRequest := &ListUserPoolsRequest{
		Keyword:  util.PtrString(""),
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &ListUserPoolsResponse{}
	result, err := AGENTIDENTITY_CLIENT.ListUserPools(listUserPoolsRequest)
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
func TestClient_ListUsers(t *testing.T) {
	listUsersRequest := &ListUsersRequest{
		UserPoolId: util.PtrString(""),
		Keyword:    util.PtrString(""),
		PageNo:     util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &ListUsersResponse{}
	result, err := AGENTIDENTITY_CLIENT.ListUsers(listUsersRequest)
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
func TestClient_OIDCDiscovery(t *testing.T) {
	oIDCDiscoveryRequest := &OIDCDiscoveryRequest{
		UserPoolId: util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.OIDCDiscovery(oIDCDiscoveryRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Oauth2idpCallback(t *testing.T) {
	oauth2idpCallbackRequest := &Oauth2idpCallbackRequest{
		ProviderId: util.PtrString(""),
		Code:       util.PtrString(""),
		State:      util.PtrString(""),
	}
	err := AGENTIDENTITY_CLIENT.Oauth2idpCallback(oauth2idpCallbackRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResetPassword(t *testing.T) {
	resetPasswordRequest := &ResetPasswordRequest{
		UserPoolId:         util.PtrString(""),
		UserId:             util.PtrString(""),
		NewPassword:        util.PtrString(""),
		Password:           util.PtrString(""),
		ForceResetPassword: util.PtrBool(false),
	}
	err := AGENTIDENTITY_CLIENT.ResetPassword(resetPasswordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TokenEndpoint(t *testing.T) {
	tokenEndpointRequest := &TokenEndpointRequest{
		UserPoolId:   util.PtrString(""),
		GrantType:    util.PtrString(""),
		Code:         util.PtrString(""),
		RefreshToken: util.PtrString(""),
		ClientId:     util.PtrString(""),
		ClientSecret: util.PtrString(""),
		RedirectUri:  util.PtrString(""),
	}
	result := &TokenEndpointResponse{}
	result, err := AGENTIDENTITY_CLIENT.TokenEndpoint(tokenEndpointRequest)
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
func TestClient_UpdateAgent(t *testing.T) {
	updateAgentRequest := &UpdateAgentRequest{
		AgentId:                         util.PtrString(""),
		Description:                     util.PtrString(""),
		AllowedResourceOauth2ReturnUrls: []*string{},
	}
	result := &UpdateAgentResponse{}
	result, err := AGENTIDENTITY_CLIENT.UpdateAgent(updateAgentRequest)
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
func TestClient_UpdateCredentialProvider(t *testing.T) {
	updateCredentialProviderRequest := &UpdateCredentialProviderRequest{
		CredentialProviderId: util.PtrString(""),
		Desc:                 util.PtrString(""),
		Credential:           nil,
	}
	result := &UpdateCredentialProviderResponse{}
	result, err := AGENTIDENTITY_CLIENT.UpdateCredentialProvider(updateCredentialProviderRequest)
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
func TestClient_UpdateIdpConfiguration(t *testing.T) {
	updateIdpConfigurationRequest := &UpdateIdpConfigurationRequest{
		UserPoolId:            util.PtrString(""),
		Id:                    util.PtrString(""),
		Name:                  util.PtrString(""),
		ClientId:              util.PtrString(""),
		ClientSecret:          util.PtrString(""),
		AuthorizationEndpoint: util.PtrString(""),
		TokenEndpoint:         util.PtrString(""),
		UserinfoEndpoint:      util.PtrString(""),
		Scopes:                []*string{},
		UserIdClaim:           util.PtrString(""),
		DisplayNameClaim:      util.PtrString(""),
		AutoCreateUser:        util.PtrBool(false),
	}
	result := &UpdateIdpConfigurationResponse{}
	result, err := AGENTIDENTITY_CLIENT.UpdateIdpConfiguration(updateIdpConfigurationRequest)
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
func TestClient_UpdateOauth2Client(t *testing.T) {
	updateOauth2ClientRequest := &UpdateOauth2ClientRequest{
		UserPoolId:      util.PtrString(""),
		Id:              util.PtrString(""),
		Name:            util.PtrString(""),
		Description:     util.PtrString(""),
		RedirectUris:    []*string{},
		GrantTypes:      []*string{},
		Scopes:          []*string{},
		AccessTokenTtl:  util.PtrInt32(int32(0)),
		RefreshTokenTtl: util.PtrInt32(int32(0)),
	}
	result := &UpdateOauth2ClientResponse{}
	result, err := AGENTIDENTITY_CLIENT.UpdateOauth2Client(updateOauth2ClientRequest)
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
		UserPoolId:  util.PtrString(""),
		Id:          util.PtrString(""),
		DisplayName: util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &UpdateUserResponse{}
	result, err := AGENTIDENTITY_CLIENT.UpdateUser(updateUserRequest)
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
func TestClient_UpdateUserPool(t *testing.T) {
	updateUserPoolRequest := &UpdateUserPoolRequest{
		Id:          util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &UpdateUserPoolResponse{}
	result, err := AGENTIDENTITY_CLIENT.UpdateUserPool(updateUserPoolRequest)
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
func TestClient_UserinfoEndpoint(t *testing.T) {
	userinfoEndpointRequest := &UserinfoEndpointRequest{
		UserPoolId:    util.PtrString(""),
		Authorization: util.PtrString(""),
	}
	result := &UserinfoEndpointResponse{}
	result, err := AGENTIDENTITY_CLIENT.UserinfoEndpoint(userinfoEndpointRequest)
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
