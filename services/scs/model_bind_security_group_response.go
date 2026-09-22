package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BindSecurityGroupResponse struct {
	bce.BaseResponse
	Success *bool `json:"success,omitempty"`
}
