package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDbStatementResponse struct {
	bce.BaseResponse
	Success    *bool              `json:"success,omitempty"`
	Code       *string            `json:"code,omitempty"`
	Message    *string            `json:"message,omitempty"`
	Statements []*StatementDetail `json:"statements,omitempty"`
}
