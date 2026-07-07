package agentidentity

type UpdateOauth2ClientRequest struct {
	UserPoolId      *string   `json:"userPoolId,omitempty"`
	Id              *string   `json:"id,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Description     *string   `json:"description,omitempty"`
	RedirectUris    []*string `json:"redirectUris,omitempty"`
	GrantTypes      []*string `json:"grantTypes,omitempty"`
	Scopes          []*string `json:"scopes,omitempty"`
	AccessTokenTtl  *int32    `json:"accessTokenTtl,omitempty"`
	RefreshTokenTtl *int32    `json:"refreshTokenTtl,omitempty"`
}
