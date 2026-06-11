package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListSharedUserResponse struct {
	bce.BaseResponse
	Users []*SharedUser `json:"users,omitempty"`
}
