package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLLMSessionsResponse struct {
	bce.BaseResponse
	Success  *bool             `json:"success,omitempty"`
	Code     *string           `json:"code,omitempty"`
	Message  *string           `json:"message,omitempty"`
	Sessions []*LLMSessionItem `json:"sessions,omitempty"`
}
