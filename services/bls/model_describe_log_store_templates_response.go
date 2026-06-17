package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLogStoreTemplatesResponse struct {
	bce.BaseResponse
	Success    *bool       `json:"success,omitempty"`
	Code       *string     `json:"code,omitempty"`
	Message    *string     `json:"message,omitempty"`
	Templates  []*Template `json:"templates,omitempty"`
	PageNo     *int32      `json:"pageNo,omitempty"`
	PageSize   *int32      `json:"pageSize,omitempty"`
	TotalCount *int32      `json:"totalCount,omitempty"`
}
