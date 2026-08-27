package iam

type CreateRoleRequest struct {
	Name                     *string `json:"name,omitempty"`
	Description              *string `json:"description,omitempty"`
	GrantType                *string `json:"grantType,omitempty"`
	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty"`
}
