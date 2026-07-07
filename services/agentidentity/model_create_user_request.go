package agentidentity

type CreateUserRequest struct {
	UserPoolId         *string `json:"userPoolId,omitempty"`
	Username           *string `json:"username,omitempty"`
	DisplayName        *string `json:"displayName,omitempty"`
	Description        *string `json:"description,omitempty"`
	Password           *string `json:"password,omitempty"`
	ForceResetPassword *bool   `json:"forceResetPassword,omitempty"`
}
