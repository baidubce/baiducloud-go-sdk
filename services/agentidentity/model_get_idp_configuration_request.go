package agentidentity

type GetIdpConfigurationRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Id         *string `json:"id,omitempty"`
}
