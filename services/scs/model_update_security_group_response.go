package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateSecurityGroupResponse struct {
	bce.BaseResponse
	Success *bool `json:"success,omitempty"`
}
