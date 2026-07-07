package agentidentity

type DeleteOauth2ClientRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Id         *string `json:"id,omitempty"`
}
