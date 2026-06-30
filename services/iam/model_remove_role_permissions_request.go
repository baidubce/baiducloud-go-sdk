package iam

type RemoveRolePermissionsRequest struct {
	RoleName   *string `json:"-"`
	PolicyName *string `json:"-"`
	PolicyType *string `json:"-"`
}
