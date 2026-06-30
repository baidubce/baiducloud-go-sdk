package iam

type RemoveUserPermissionsRequest struct {
	UserName   *string `json:"-"`
	PolicyName *string `json:"-"`
	PolicyType *string `json:"-"`
}
