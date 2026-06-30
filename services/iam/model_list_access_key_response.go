package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListAccessKeyResponse struct {
	bce.BaseResponse
	AccessKeys []*AccessKey `json:"accessKeys,omitempty"`
}
