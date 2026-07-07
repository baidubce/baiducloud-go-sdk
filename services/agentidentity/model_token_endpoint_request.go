package agentidentity

type TokenEndpointRequest struct {
	UserPoolId   *string `json:"-"`
	GrantType    *string `json:"grant_type,omitempty"`
	Code         *string `json:"code,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	ClientId     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	RedirectUri  *string `json:"redirect_uri,omitempty"`
}
