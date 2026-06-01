package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLLMSpansResponse struct {
	bce.BaseResponse
	Success     *bool      `json:"success,omitempty"`
	Code        *string    `json:"code,omitempty"`
	Message     *string    `json:"message,omitempty"`
	Spans       []*LLMSpan `json:"spans,omitempty"`
	NextMarker  *string    `json:"nextMarker,omitempty"`
	IsTruncated *bool      `json:"isTruncated,omitempty"`
}
