package agentidentity

type GetOauth2ClientRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Id         *string `json:"id,omitempty"`
}
