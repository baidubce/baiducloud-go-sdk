package sts

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AssumeRoleResponse struct {
	bce.BaseResponse
	Credential *string `json:"credential,omitempty"`
}
