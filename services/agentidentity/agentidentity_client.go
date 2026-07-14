package agentidentity

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AuthorizeEndpoint
//
// PARAMS:
//   - request: the arguments to AuthorizeEndpoint
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AuthorizeEndpoint(request *AuthorizeEndpointRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAuthorizeEndpointUri(VERSION_V1, util.StringValue(request.UserPoolId))).
		WithQueryParamFilter("client_id", util.StringValue(request.ClientId)).
		WithQueryParamFilter("redirect_uri", util.StringValue(request.RedirectUri)).
		WithQueryParamFilter("response_type", util.StringValue(request.ResponseType)).
		WithQueryParamFilter("scope", util.StringValue(request.Scope)).
		WithQueryParamFilter("state", util.StringValue(request.State)).
		WithQueryParamFilter("nonce", util.StringValue(request.Nonce)).
		Do()
}

// BatchAcquisitionOfUsers
//
// PARAMS:
//   - request: the arguments to BatchAcquisitionOfUsers
//
// RETURNS:
//   - BatchAcquisitionOfUsersResponse: The return type of the BatchAcquisitionOfUsers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchAcquisitionOfUsers(request *BatchAcquisitionOfUsersRequest) (*BatchAcquisitionOfUsersResponse, error) {
	result := &BatchAcquisitionOfUsersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBatchAcquisitionOfUsersUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BatchGetResourceApiKey
//
// PARAMS:
//   - request: the arguments to BatchGetResourceApiKey
//
// RETURNS:
//   - BatchGetResourceApiKeyResponse: The return type of the BatchGetResourceApiKey interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchGetResourceApiKey(request *BatchGetResourceApiKeyRequest) (*BatchGetResourceApiKeyResponse, error) {
	result := &BatchGetResourceApiKeyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithHost(util.StringValue(request.XBceWorkloadAccessToken)).
		WithURL(getBatchGetResourceApiKeyUri(VERSION_V1, util.StringValue(request.XBceWorkloadAccessToken))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CompleteOauth2session
//
// PARAMS:
//   - request: the arguments to CompleteOauth2session
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CompleteOauth2session(request *CompleteOauth2sessionRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithHost(util.StringValue(request.XBceWorkloadAccessToken)).
		WithURL(getCompleteOauth2sessionUri(VERSION_V1, util.StringValue(request.XBceWorkloadAccessToken))).
		WithBody(request).
		Do()
}

// CreateAgent
//
// PARAMS:
//   - request: the arguments to CreateAgent
//
// RETURNS:
//   - CreateAgentResponse: The return type of the CreateAgent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAgent(request *CreateAgentRequest) (*CreateAgentResponse, error) {
	result := &CreateAgentResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAgentUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCredentialProvider
//
// PARAMS:
//   - request: the arguments to CreateCredentialProvider
//
// RETURNS:
//   - CreateCredentialProviderResponse: The return type of the CreateCredentialProvider interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCredentialProvider(request *CreateCredentialProviderRequest) (*CreateCredentialProviderResponse, error) {
	result := &CreateCredentialProviderResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateCredentialProviderUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIdpConfiguration
//
// PARAMS:
//   - request: the arguments to CreateIdpConfiguration
//
// RETURNS:
//   - CreateIdpConfigurationResponse: The return type of the CreateIdpConfiguration interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIdpConfiguration(request *CreateIdpConfigurationRequest) (*CreateIdpConfigurationResponse, error) {
	result := &CreateIdpConfigurationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateIdpConfigurationUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOauth2Client
//
// PARAMS:
//   - request: the arguments to CreateOauth2Client
//
// RETURNS:
//   - CreateOauth2ClientResponse: The return type of the CreateOauth2Client interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateOauth2Client(request *CreateOauth2ClientRequest) (*CreateOauth2ClientResponse, error) {
	result := &CreateOauth2ClientResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateOauth2ClientUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser
//
// PARAMS:
//   - request: the arguments to CreateUser
//
// RETURNS:
//   - CreateUserResponse: The return type of the CreateUser interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateUser(request *CreateUserRequest) (*CreateUserResponse, error) {
	result := &CreateUserResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateUserUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUserPool
//
// PARAMS:
//   - request: the arguments to CreateUserPool
//
// RETURNS:
//   - CreateUserPoolResponse: The return type of the CreateUserPool interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateUserPool(request *CreateUserPoolRequest) (*CreateUserPoolResponse, error) {
	result := &CreateUserPoolResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateUserPoolUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAgent
//
// PARAMS:
//   - request: the arguments to DeleteAgent
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAgent(request *DeleteAgentRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteAgentUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// DeleteCredentialProvider
//
// PARAMS:
//   - request: the arguments to DeleteCredentialProvider
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteCredentialProvider(request *DeleteCredentialProviderRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteCredentialProviderUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// DeleteIdpConfiguration
//
// PARAMS:
//   - request: the arguments to DeleteIdpConfiguration
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteIdpConfiguration(request *DeleteIdpConfigurationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteIdpConfigurationUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// DeleteOauth2Client
//
// PARAMS:
//   - request: the arguments to DeleteOauth2Client
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteOauth2Client(request *DeleteOauth2ClientRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteOauth2ClientUri(VERSION_V1)).
		WithBody(request).
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
		WithMethod(http.POST).
		WithURL(getDeleteUserUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// DeleteUserPool
//
// PARAMS:
//   - request: the arguments to DeleteUserPool
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteUserPool(request *DeleteUserPoolRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteUserPoolUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// DisableIdpConfiguration
//
// PARAMS:
//   - request: the arguments to DisableIdpConfiguration
//
// RETURNS:
//   - DisableIdpConfigurationResponse: The return type of the DisableIdpConfiguration interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DisableIdpConfiguration(request *DisableIdpConfigurationRequest) (*DisableIdpConfigurationResponse, error) {
	result := &DisableIdpConfigurationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDisableIdpConfigurationUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EnableIdpConfiguration
//
// PARAMS:
//   - request: the arguments to EnableIdpConfiguration
//
// RETURNS:
//   - EnableIdpConfigurationResponse: The return type of the EnableIdpConfiguration interface.
//   - error: nil if success otherwise the specific error
func (c *Client) EnableIdpConfiguration(request *EnableIdpConfigurationRequest) (*EnableIdpConfigurationResponse, error) {
	result := &EnableIdpConfigurationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEnableIdpConfigurationUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAgent
//
// PARAMS:
//   - request: the arguments to GetAgent
//
// RETURNS:
//   - GetAgentResponse: The return type of the GetAgent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetAgent(request *GetAgentRequest) (*GetAgentResponse, error) {
	result := &GetAgentResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetAgentUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCredentialProvider
//
// PARAMS:
//   - request: the arguments to GetCredentialProvider
//
// RETURNS:
//   - GetCredentialProviderResponse: The return type of the GetCredentialProvider interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetCredentialProvider(request *GetCredentialProviderRequest) (*GetCredentialProviderResponse, error) {
	result := &GetCredentialProviderResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetCredentialProviderUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIdpConfiguration
//
// PARAMS:
//   - request: the arguments to GetIdpConfiguration
//
// RETURNS:
//   - GetIdpConfigurationResponse: The return type of the GetIdpConfiguration interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetIdpConfiguration(request *GetIdpConfigurationRequest) (*GetIdpConfigurationResponse, error) {
	result := &GetIdpConfigurationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetIdpConfigurationUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOauth2Client
//
// PARAMS:
//   - request: the arguments to GetOauth2Client
//
// RETURNS:
//   - GetOauth2ClientResponse: The return type of the GetOauth2Client interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetOauth2Client(request *GetOauth2ClientRequest) (*GetOauth2ClientResponse, error) {
	result := &GetOauth2ClientResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetOauth2ClientUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetResourceApikey
//
// PARAMS:
//   - request: the arguments to GetResourceApikey
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetResourceApikey(request *GetResourceApikeyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithHost(util.StringValue(request.XBceWorkloadAccessToken)).
		WithURL(getGetResourceApikeyUri(VERSION_V1, util.StringValue(request.XBceWorkloadAccessToken))).
		WithBody(request).
		Do()
}

// GetResourceOauth2token
//
// PARAMS:
//   - request: the arguments to GetResourceOauth2token
//
// RETURNS:
//   - GetResourceOauth2tokenResponse: The return type of the GetResourceOauth2token interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetResourceOauth2token(request *GetResourceOauth2tokenRequest) (*GetResourceOauth2tokenResponse, error) {
	result := &GetResourceOauth2tokenResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithHost(util.StringValue(request.XBceWorkloadAccessToken)).
		WithURL(getGetResourceOauth2tokenUri(VERSION_V1, util.StringValue(request.XBceWorkloadAccessToken))).
		WithBody(request).
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
//   - GetUserResponse: The return type of the GetUser interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetUser(request *GetUserRequest) (*GetUserResponse, error) {
	result := &GetUserResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetUserUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserPool
//
// PARAMS:
//   - request: the arguments to GetUserPool
//
// RETURNS:
//   - GetUserPoolResponse: The return type of the GetUserPool interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetUserPool(request *GetUserPoolRequest) (*GetUserPoolResponse, error) {
	result := &GetUserPoolResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetUserPoolUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWATForUser
//
// PARAMS:
//   - request: the arguments to GetWATForUser
//
// RETURNS:
//   - GetWATForUserResponse: The return type of the GetWATForUser interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetWATForUser(request *GetWATForUserRequest) (*GetWATForUserResponse, error) {
	result := &GetWATForUserResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetWATForUserUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWorkloadAccessToken
//
// PARAMS:
//   - request: the arguments to GetWorkloadAccessToken
//
// RETURNS:
//   - GetWorkloadAccessTokenResponse: The return type of the GetWorkloadAccessToken interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetWorkloadAccessToken(request *GetWorkloadAccessTokenRequest) (*GetWorkloadAccessTokenResponse, error) {
	result := &GetWorkloadAccessTokenResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetWorkloadAccessTokenUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListAgents
//
// PARAMS:
//   - request: the arguments to ListAgents
//
// RETURNS:
//   - ListAgentsResponse: The return type of the ListAgents interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListAgents(request *ListAgentsRequest) (*ListAgentsResponse, error) {
	result := &ListAgentsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListAgentsUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCredentialProviders
//
// PARAMS:
//   - request: the arguments to ListCredentialProviders
//
// RETURNS:
//   - ListCredentialProvidersResponse: The return type of the ListCredentialProviders interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListCredentialProviders(request *ListCredentialProvidersRequest) (*ListCredentialProvidersResponse, error) {
	result := &ListCredentialProvidersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListCredentialProvidersUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListIdpConfigurations
//
// PARAMS:
//   - request: the arguments to ListIdpConfigurations
//
// RETURNS:
//   - ListIdpConfigurationsResponse: The return type of the ListIdpConfigurations interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListIdpConfigurations(request *ListIdpConfigurationsRequest) (*ListIdpConfigurationsResponse, error) {
	result := &ListIdpConfigurationsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListIdpConfigurationsUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListOauth2Clients
//
// PARAMS:
//   - request: the arguments to ListOauth2Clients
//
// RETURNS:
//   - ListOauth2ClientsResponse: The return type of the ListOauth2Clients interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListOauth2Clients(request *ListOauth2ClientsRequest) (*ListOauth2ClientsResponse, error) {
	result := &ListOauth2ClientsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListOauth2ClientsUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListUserPools
//
// PARAMS:
//   - request: the arguments to ListUserPools
//
// RETURNS:
//   - ListUserPoolsResponse: The return type of the ListUserPools interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListUserPools(request *ListUserPoolsRequest) (*ListUserPoolsResponse, error) {
	result := &ListUserPoolsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListUserPoolsUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListUsers
//
// PARAMS:
//   - request: the arguments to ListUsers
//
// RETURNS:
//   - ListUsersResponse: The return type of the ListUsers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListUsers(request *ListUsersRequest) (*ListUsersResponse, error) {
	result := &ListUsersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListUsersUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OIDCDiscovery
//
// PARAMS:
//   - request: the arguments to OIDCDiscovery
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) OIDCDiscovery(request *OIDCDiscoveryRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getOIDCDiscoveryUri(VERSION_V1, util.StringValue(request.UserPoolId))).
		Do()
}

// Oauth2idpCallback
//
// PARAMS:
//   - request: the arguments to Oauth2idpCallback
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) Oauth2idpCallback(request *Oauth2idpCallbackRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getOauth2idpCallbackUri(VERSION_V1, util.StringValue(request.ProviderId))).
		WithQueryParamFilter("code", util.StringValue(request.Code)).
		WithQueryParamFilter("state", util.StringValue(request.State)).
		Do()
}

// ResetPassword
//
// PARAMS:
//   - request: the arguments to ResetPassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ResetPassword(request *ResetPasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getResetPasswordUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// TokenEndpoint
//
// PARAMS:
//   - request: the arguments to TokenEndpoint
//
// RETURNS:
//   - TokenEndpointResponse: The return type of the TokenEndpoint interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TokenEndpoint(request *TokenEndpointRequest) (*TokenEndpointResponse, error) {
	result := &TokenEndpointResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTokenEndpointUri(VERSION_V1, util.StringValue(request.UserPoolId))).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAgent
//
// PARAMS:
//   - request: the arguments to UpdateAgent
//
// RETURNS:
//   - UpdateAgentResponse: The return type of the UpdateAgent interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateAgent(request *UpdateAgentRequest) (*UpdateAgentResponse, error) {
	result := &UpdateAgentResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateAgentUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCredentialProvider
//
// PARAMS:
//   - request: the arguments to UpdateCredentialProvider
//
// RETURNS:
//   - UpdateCredentialProviderResponse: The return type of the UpdateCredentialProvider interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateCredentialProvider(request *UpdateCredentialProviderRequest) (*UpdateCredentialProviderResponse, error) {
	result := &UpdateCredentialProviderResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateCredentialProviderUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIdpConfiguration
//
// PARAMS:
//   - request: the arguments to UpdateIdpConfiguration
//
// RETURNS:
//   - UpdateIdpConfigurationResponse: The return type of the UpdateIdpConfiguration interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateIdpConfiguration(request *UpdateIdpConfigurationRequest) (*UpdateIdpConfigurationResponse, error) {
	result := &UpdateIdpConfigurationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateIdpConfigurationUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateOauth2Client
//
// PARAMS:
//   - request: the arguments to UpdateOauth2Client
//
// RETURNS:
//   - UpdateOauth2ClientResponse: The return type of the UpdateOauth2Client interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateOauth2Client(request *UpdateOauth2ClientRequest) (*UpdateOauth2ClientResponse, error) {
	result := &UpdateOauth2ClientResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateOauth2ClientUri(VERSION_V1)).
		WithBody(request).
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
//   - UpdateUserResponse: The return type of the UpdateUser interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateUser(request *UpdateUserRequest) (*UpdateUserResponse, error) {
	result := &UpdateUserResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateUserUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUserPool
//
// PARAMS:
//   - request: the arguments to UpdateUserPool
//
// RETURNS:
//   - UpdateUserPoolResponse: The return type of the UpdateUserPool interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateUserPool(request *UpdateUserPoolRequest) (*UpdateUserPoolResponse, error) {
	result := &UpdateUserPoolResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateUserPoolUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserinfoEndpoint
//
// PARAMS:
//   - request: the arguments to UserinfoEndpoint
//
// RETURNS:
//   - UserinfoEndpointResponse: The return type of the UserinfoEndpoint interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserinfoEndpoint(request *UserinfoEndpointRequest) (*UserinfoEndpointResponse, error) {
	result := &UserinfoEndpointResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithHost(util.StringValue(request.Authorization)).
		WithURL(getUserinfoEndpointUri(VERSION_V1, util.StringValue(request.UserPoolId), util.StringValue(request.Authorization))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
