package agentidentity

type ResetPasswordRequest struct {
	UserPoolId         *string `json:"userPoolId,omitempty"`
	UserId             *string `json:"userId,omitempty"`
	NewPassword        *string `json:"newPassword,omitempty"`
	Password           *string `json:"password,omitempty"`
	ForceResetPassword *bool   `json:"forceResetPassword,omitempty"`
}
