package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLLMTraceResponse struct {
	bce.BaseResponse
	Success      *bool      `json:"success,omitempty"`
	Code         *string    `json:"code,omitempty"`
	Message      *string    `json:"message,omitempty"`
	Duration     *int32     `json:"duration,omitempty"`
	MinStartTime *int32     `json:"minStartTime,omitempty"`
	MaxEndTime   *int32     `json:"maxEndTime,omitempty"`
	Tokens       *int32     `json:"tokens,omitempty"`
	Input        *string    `json:"input,omitempty"`
	Output       *string    `json:"output,omitempty"`
	RootSpans    []*LLMSpan `json:"rootSpans,omitempty"`
}
