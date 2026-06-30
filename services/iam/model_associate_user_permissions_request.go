package iam

type AssociateUserPermissionsRequest struct {
	UserName   *string `json:"-"`
	PolicyName *string `json:"-"`
	PolicyType *string `json:"-"`
}
