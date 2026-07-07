package agentidentity

type UserPoolDTO struct {
	Id                    *string `json:"id,omitempty"`
	DomainId              *string `json:"domainId,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Description           *string `json:"description,omitempty"`
	UserCount             *int32  `json:"userCount,omitempty"`
	ClientCount           *int32  `json:"clientCount,omitempty"`
	IdpCount              *int32  `json:"idpCount,omitempty"`
	CallbackUrl           *string `json:"callbackUrl,omitempty"`
	DiscoveryUrl          *string `json:"discoveryUrl,omitempty"`
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty"`
	TokenEndpoint         *string `json:"tokenEndpoint,omitempty"`
	UserinfoEndpoint      *string `json:"userinfoEndpoint,omitempty"`
	JwksUrl               *string `json:"jwksUrl,omitempty"`
	Enabled               *bool   `json:"enabled,omitempty"`
	CreatedAt             *string `json:"createdAt,omitempty"`
}
