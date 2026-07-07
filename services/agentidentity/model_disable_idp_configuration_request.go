package agentidentity

type DisableIdpConfigurationRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Id         *string `json:"id,omitempty"`
}
