package agentidentity

type UserinfoEndpointRequest struct {
	UserPoolId    *string `json:"-"`
	Authorization *string `json:"-"`
}
