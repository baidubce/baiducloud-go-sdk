package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListLogStoreViewResponse struct {
	bce.BaseResponse
	Order      *string         `json:"order,omitempty"`
	OrderBy    *string         `json:"orderBy,omitempty"`
	PageNo     *int32          `json:"pageNo,omitempty"`
	PageSize   *int32          `json:"pageSize,omitempty"`
	TotalCount *int32          `json:"totalCount,omitempty"`
	Views      []*LogStoreView `json:"views,omitempty"`
}
