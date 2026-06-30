package iam

type UpdateRoleRequest struct {
	RoleName                 *string `json:"-"`
	Name                     *string `json:"name,omitempty"`
	Description              *string `json:"description,omitempty"`
	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty"`
}
