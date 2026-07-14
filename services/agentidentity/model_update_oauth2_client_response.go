package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateOauth2ClientResponse struct {
	bce.BaseResponse
	Id              *string   `json:"id,omitempty"`
	ClientId        *string   `json:"clientId,omitempty"`
	ClientSecret    *string   `json:"clientSecret,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Description     *string   `json:"description,omitempty"`
	ClientType      *string   `json:"clientType,omitempty"`
	RedirectUris    []*string `json:"redirectUris,omitempty"`
	GrantTypes      []*string `json:"grantTypes,omitempty"`
	Scopes          []*string `json:"scopes,omitempty"`
	AccessTokenTtl  *int32    `json:"accessTokenTtl,omitempty"`
	RefreshTokenTtl *int32    `json:"refreshTokenTtl,omitempty"`
	Enabled         *bool     `json:"enabled,omitempty"`
	LoginUrl        *string   `json:"loginUrl,omitempty"`
	CreatedAt       *string   `json:"createdAt,omitempty"`
}
