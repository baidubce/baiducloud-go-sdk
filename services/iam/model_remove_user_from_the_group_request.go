package iam

type RemoveUserFromTheGroupRequest struct {
	UserName  *string `json:"-"`
	GroupName *string `json:"-"`
}
