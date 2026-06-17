package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateAlarmPolicyResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}
