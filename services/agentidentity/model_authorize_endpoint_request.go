package agentidentity

type AuthorizeEndpointRequest struct {
	UserPoolId   *string `json:"-"`
	ClientId     *string `json:"-"`
	RedirectUri  *string `json:"-"`
	ResponseType *string `json:"-"`
	Scope        *string `json:"-"`
	State        *string `json:"-"`
	Nonce        *string `json:"-"`
}
