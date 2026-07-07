package agentidentity

type GetUserRequest struct {
	UserPoolId *string `json:"userPoolId,omitempty"`
	Id         *string `json:"id,omitempty"`
	Username   *string `json:"username,omitempty"`
}
