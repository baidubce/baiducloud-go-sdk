package iam

type AssociateGroupPermissionsRequest struct {
	GroupName  *string `json:"-"`
	PolicyName *string `json:"-"`
	PolicyType *string `json:"-"`
}
