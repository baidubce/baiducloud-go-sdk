package agentidentity

type UpdateUserRequest struct {
	UserPoolId  *string `json:"userPoolId,omitempty"`
	Id          *string `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Description *string `json:"description,omitempty"`
}
