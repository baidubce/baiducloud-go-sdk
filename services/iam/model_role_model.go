package iam

type RoleModel struct {
	Id                       *string `json:"id,omitempty"`
	Name                     *string `json:"name,omitempty"`
	CreateTime               *string `json:"createTime,omitempty"`
	Description              *string `json:"description,omitempty"`
	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty"`
}
