package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListLogStoreResponse struct {
	bce.BaseResponse
	Code       *string           `json:"code,omitempty"`
	Success    *bool             `json:"success,omitempty"`
	Order      *string           `json:"order,omitempty"`
	OrderBy    *string           `json:"orderBy,omitempty"`
	PageNo     *int32            `json:"pageNo,omitempty"`
	PageSize   *int32            `json:"pageSize,omitempty"`
	Result     []*LogStoreDetail `json:"result,omitempty"`
	TotalCount *int32            `json:"totalCount,omitempty"`
}
