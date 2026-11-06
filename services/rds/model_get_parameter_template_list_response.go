package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetParameterTemplateListResponse struct {
	bce.BaseResponse
	List    []*Template `json:"list,omitempty"`
	Page    *int32      `json:"page,omitempty"`
	Perpage *int32      `json:"perpage,omitempty"`
	Total   *int32      `json:"total,omitempty"`
}
