package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetHotGroupListResponse struct {
	bce.BaseResponse
	Result     []*Result `json:"result,omitempty"`
	TotalCount *int32    `json:"totalCount,omitempty"`
	PageNo     *int32    `json:"pageNo,omitempty"`
	PageSize   *int32    `json:"pageSize,omitempty"`
}
