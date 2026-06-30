package iam

type RemoveGroupPermissionsRequest struct {
	GroupName  *string `json:"-"`
	PolicyName *string `json:"-"`
	PolicyType *string `json:"-"`
}
