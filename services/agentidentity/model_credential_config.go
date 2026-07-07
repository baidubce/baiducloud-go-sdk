package agentidentity

type CredentialConfig struct {
	ApiKey                *string `json:"apiKey,omitempty"`
	ConfigType            *string `json:"configType,omitempty"`
	Provider              *string `json:"provider,omitempty"`
	DiscoveryUrl          *string `json:"discoveryUrl,omitempty"`
	Issuer                *string `json:"issuer,omitempty"`
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty"`
	TokenEndpoint         *string `json:"tokenEndpoint,omitempty"`
	UserinfoEndpoint      *string `json:"userinfoEndpoint,omitempty"`
	JwksUri               *string `json:"jwksUri,omitempty"`
	ClientId              *string `json:"clientId,omitempty"`
	ClientSecret          *string `json:"clientSecret,omitempty"`
	Scopes                *string `json:"scopes,omitempty"`
	RedirectUri           *string `json:"redirectUri,omitempty"`
	TokenEncryptionKeyId  *string `json:"tokenEncryptionKeyId,omitempty"`
	RoleArn               *string `json:"roleArn,omitempty"`
	ExternalId            *string `json:"externalId,omitempty"`
	DurationSeconds       *int32  `json:"durationSeconds,omitempty"`
}
