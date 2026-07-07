package agentidentity

type UpdateIdpConfigurationRequest struct {
	UserPoolId            *string   `json:"userPoolId,omitempty"`
	Id                    *string   `json:"id,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	ClientId              *string   `json:"clientId,omitempty"`
	ClientSecret          *string   `json:"clientSecret,omitempty"`
	AuthorizationEndpoint *string   `json:"authorizationEndpoint,omitempty"`
	TokenEndpoint         *string   `json:"tokenEndpoint,omitempty"`
	UserinfoEndpoint      *string   `json:"userinfoEndpoint,omitempty"`
	Scopes                []*string `json:"scopes,omitempty"`
	UserIdClaim           *string   `json:"userIdClaim,omitempty"`
	DisplayNameClaim      *string   `json:"displayNameClaim,omitempty"`
	AutoCreateUser        *bool     `json:"autoCreateUser,omitempty"`
}
