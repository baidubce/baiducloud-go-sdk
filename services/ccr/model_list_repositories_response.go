package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRepositoriesResponse struct {
	bce.BaseResponse
	Total    *int32        `json:"total,omitempty"`
	PageNo   *int32        `json:"pageNo,omitempty"`
	PageSize *int32        `json:"pageSize,omitempty"`
	Items    []*Repository `json:"items,omitempty"`
}
