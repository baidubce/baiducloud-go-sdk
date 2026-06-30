package iam

type AddUserToGroupRequest struct {
	UserName  *string `json:"-"`
	GroupName *string `json:"-"`
}
