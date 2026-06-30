package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListUsersWithinTheGroupResponse struct {
	bce.BaseResponse
	Users []*UserModel `json:"users,omitempty"`
}
