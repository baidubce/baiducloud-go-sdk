package agentidentity

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "agentidentity." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_AGENT_IDENTITY = "agent-identity"

	CONSTANT_WORKLOAD_ACCESS_TOKEN = "workload-access-token"

	CONSTANT_AGENT = "agent"

	CONSTANT_LIST = "list"

	CONSTANT_USER_POOL = "user-pool"

	CONSTANT_USER = "user"

	CONSTANT_RESET_PASSWORD = "resetPassword"

	CONSTANT_OAUTH2_CLIENT = "oauth2-client"

	CONSTANT_GET = "get"

	CONSTANT_UPDATE = "update"

	CONSTANT_CREDENTIAL = "credential"

	CONSTANT_APIKEY = "apikey"

	CONSTANT_BATCH = "batch"

	CONSTANT_IDP_CONFIG = "idp-config"

	CONSTANT_CREATE = "create"

	CONSTANT_DELETE = "delete"

	CONSTANT_CREDENTIAL_PROVIDER = "credential-provider"

	CONSTANT_INBOUND = "inbound"

	CONSTANT_TOKEN = "token"

	CONSTANT_WORKLOAD_ACCESS_TOKEN_FOR_USER = "workload-access-token-for-user"

	CONSTANT_OAUTH2 = "oauth2"

	CONSTANT_CALLBACK = "callback"

	CONSTANT_WELL_KNOWN = ".well-known"

	CONSTANT_OPENID_CONFIGURATION = "openid-configuration"

	CONSTANT_ENABLE = "enable"

	CONSTANT_USERINFO = "userinfo"

	CONSTANT_AUTHORIZE = "authorize"

	CONSTANT_COMPLETE_AUTH = "complete-auth"

	CONSTANT_DISABLE = "disable"
)

// Client of agentidentity service is a kind of BceClient, so derived from BceClient
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

func getAuthorizeEndpointUri(version string, UserPoolId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_INBOUND + bce.URI_PREFIX + UserPoolId + bce.URI_PREFIX + CONSTANT_AUTHORIZE
}
func getBatchAcquisitionOfUsersUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_BATCH
}
func getBatchGetResourceApiKeyUri(version string, XBceWorkloadAccessToken string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL + bce.URI_PREFIX + CONSTANT_APIKEY + bce.URI_PREFIX + CONSTANT_BATCH
}
func getCompleteOauth2sessionUri(version string, XBceWorkloadAccessToken string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_OAUTH2 + bce.URI_PREFIX + CONSTANT_COMPLETE_AUTH
}
func getCreateAgentUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_AGENT + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateCredentialProviderUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL_PROVIDER + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateIdpConfigurationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateOauth2ClientUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_OAUTH2_CLIENT + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_CREATE
}
func getCreateUserPoolUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_CREATE
}
func getDeleteAgentUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_AGENT + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteCredentialProviderUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL_PROVIDER + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteIdpConfigurationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteOauth2ClientUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_OAUTH2_CLIENT + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteUserPoolUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDisableIdpConfigurationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_DISABLE
}
func getEnableIdpConfigurationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_ENABLE
}
func getGetAgentUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_AGENT + bce.URI_PREFIX + CONSTANT_GET
}
func getGetCredentialProviderUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL_PROVIDER + bce.URI_PREFIX + CONSTANT_GET
}
func getGetIdpConfigurationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_GET
}
func getGetOauth2ClientUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_OAUTH2_CLIENT + bce.URI_PREFIX + CONSTANT_GET
}
func getGetResourceApikeyUri(version string, XBceWorkloadAccessToken string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL + bce.URI_PREFIX + CONSTANT_APIKEY
}
func getGetResourceOauth2tokenUri(version string, XBceWorkloadAccessToken string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL + bce.URI_PREFIX + CONSTANT_OAUTH2
}
func getGetUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_GET
}
func getGetUserPoolUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_GET
}
func getGetWATForUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_WORKLOAD_ACCESS_TOKEN_FOR_USER
}
func getGetWorkloadAccessTokenUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_WORKLOAD_ACCESS_TOKEN
}
func getListAgentsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_AGENT + bce.URI_PREFIX + CONSTANT_LIST
}
func getListCredentialProvidersUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL_PROVIDER + bce.URI_PREFIX + CONSTANT_LIST
}
func getListIdpConfigurationsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_LIST
}
func getListOauth2ClientsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_OAUTH2_CLIENT + bce.URI_PREFIX + CONSTANT_LIST
}
func getListUserPoolsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_LIST
}
func getListUsersUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_LIST
}
func getOIDCDiscoveryUri(version string, UserPoolId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_INBOUND + bce.URI_PREFIX + UserPoolId + bce.URI_PREFIX + CONSTANT_WELL_KNOWN + bce.URI_PREFIX + CONSTANT_OPENID_CONFIGURATION
}
func getOauth2idpCallbackUri(version string, ProviderId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_OAUTH2 + bce.URI_PREFIX + CONSTANT_CALLBACK + bce.URI_PREFIX + ProviderId
}
func getResetPasswordUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_RESET_PASSWORD
}
func getTokenEndpointUri(version string, UserPoolId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_INBOUND + bce.URI_PREFIX + UserPoolId + bce.URI_PREFIX + CONSTANT_TOKEN
}
func getUpdateAgentUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_AGENT + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateCredentialProviderUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_CREDENTIAL_PROVIDER + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateIdpConfigurationUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_IDP_CONFIG + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateOauth2ClientUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_OAUTH2_CLIENT + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateUserUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUpdateUserPoolUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_USER_POOL + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getUserinfoEndpointUri(version string, UserPoolId string, Authorization string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AGENT_IDENTITY + bce.URI_PREFIX + CONSTANT_INBOUND + bce.URI_PREFIX + UserPoolId + bce.URI_PREFIX + CONSTANT_USERINFO
}
