package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListAlarmExecutionsResponse struct {
	bce.BaseResponse
	Success *bool                `json:"success,omitempty"`
	Code    *string              `json:"code,omitempty"`
	Message *string              `json:"message,omitempty"`
	Result  *ExecutionListResult `json:"result,omitempty"`
}
