package agentidentity

type DeleteUserRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Id         *string `json:"id,omitempty"`
}
