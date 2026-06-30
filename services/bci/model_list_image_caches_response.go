package bci

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListImageCachesResponse struct {
	bce.BaseResponse
	TotalCount *int32             `json:"totalCount,omitempty"`
	PageNo     *int32             `json:"pageNo,omitempty"`
	PageSize   *int32             `json:"pageSize,omitempty"`
	Result     []*ImageCacheModel `json:"result,omitempty"`
}
