package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListUserResponse struct {
	bce.BaseResponse
	Users []*UserModel `json:"users,omitempty"`
}
