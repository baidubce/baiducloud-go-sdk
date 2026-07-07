package agentidentity

type UserPoolIdpConfigDTO struct {
	Id                    *string   `json:"id,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	IdpType               *string   `json:"idpType,omitempty"`
	IdpProvider           *string   `json:"idpProvider,omitempty"`
	ClientId              *string   `json:"clientId,omitempty"`
	ClientSecret          *string   `json:"clientSecret,omitempty"`
	DiscoveryUrl          *string   `json:"discoveryUrl,omitempty"`
	AuthorizationEndpoint *string   `json:"authorizationEndpoint,omitempty"`
	TokenEndpoint         *string   `json:"tokenEndpoint,omitempty"`
	UserinfoEndpoint      *string   `json:"userinfoEndpoint,omitempty"`
	Scopes                []*string `json:"scopes,omitempty"`
	UserIdClaim           *string   `json:"userIdClaim,omitempty"`
	DisplayNameClaim      *string   `json:"displayNameClaim,omitempty"`
	AutoCreateUser        *bool     `json:"autoCreateUser,omitempty"`
	Enabled               *bool     `json:"enabled,omitempty"`
	CallbackUrl           *string   `json:"callbackUrl,omitempty"`
	CreatedAt             *string   `json:"createdAt,omitempty"`
}
