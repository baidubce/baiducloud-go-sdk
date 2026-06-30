package iam

type AssociateRolePermissionsRequest struct {
	RoleName   *string `json:"-"`
	PolicyName *string `json:"-"`
	PolicyType *string `json:"-"`
}
