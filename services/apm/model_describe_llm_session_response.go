package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLLMSessionResponse struct {
	bce.BaseResponse
	Success      *bool           `json:"success,omitempty"`
	Code         *string         `json:"code,omitempty"`
	Message      *string         `json:"message,omitempty"`
	StartTime    *string         `json:"startTime,omitempty"`
	EndTime      *string         `json:"endTime,omitempty"`
	Duration     *int32          `json:"duration,omitempty"`
	UserId       *string         `json:"userId,omitempty"`
	TraceCount   *int32          `json:"traceCount,omitempty"`
	InputTokens  *int32          `json:"inputTokens,omitempty"`
	OutputTokens *int32          `json:"outputTokens,omitempty"`
	TotalTokens  *int32          `json:"totalTokens,omitempty"`
	Traces       []*SessionTrace `json:"traces,omitempty"`
}
