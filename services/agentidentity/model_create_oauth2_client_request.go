package agentidentity

type CreateOauth2ClientRequest struct {
	UserPoolId      *string   `json:"userPoolId,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Description     *string   `json:"description,omitempty"`
	ClientType      *string   `json:"clientType,omitempty"`
	RedirectUris    []*string `json:"redirectUris,omitempty"`
	GrantTypes      []*string `json:"grantTypes,omitempty"`
	Scopes          []*string `json:"scopes,omitempty"`
	AccessTokenTtl  *int32    `json:"accessTokenTtl,omitempty"`
	RefreshTokenTtl *int32    `json:"refreshTokenTtl,omitempty"`
}
