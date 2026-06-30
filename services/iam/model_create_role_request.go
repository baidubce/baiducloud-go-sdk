package iam

type CreateRoleRequest struct {
	Name                     *string `json:"name,omitempty"`
	Description              *string `json:"description,omitempty"`
	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty"`
}
