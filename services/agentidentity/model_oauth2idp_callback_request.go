package agentidentity

type Oauth2idpCallbackRequest struct {
	ProviderId *string `json:"-"`
	Code       *string `json:"-"`
	State      *string `json:"-"`
}
