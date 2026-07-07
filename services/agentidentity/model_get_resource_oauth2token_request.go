package agentidentity

type GetResourceOauth2tokenRequest struct {
	XBceWorkloadAccessToken        *string   `json:"-"`
	ResourceCredentialProviderName *string   `json:"resourceCredentialProviderName,omitempty"`
	Scopes                         []*string `json:"scopes,omitempty"`
	Oauth2Flow                     *string   `json:"oauth2Flow,omitempty"`
	ResourceOauth2ReturnUrl        *string   `json:"resourceOauth2ReturnUrl,omitempty"`
	SessionUri                     *string   `json:"sessionUri,omitempty"`
	ForceAuthentication            *bool     `json:"forceAuthentication,omitempty"`
	WorkloadAccessToken            *string   `json:"workloadAccessToken,omitempty"`
}
