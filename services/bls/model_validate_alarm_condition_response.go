package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ValidateAlarmConditionResponse struct {
	bce.BaseResponse
	Success *bool                    `json:"success,omitempty"`
	Code    *string                  `json:"code,omitempty"`
	Message *string                  `json:"message,omitempty"`
	Result  *ConditionValidateResult `json:"result,omitempty"`
}
