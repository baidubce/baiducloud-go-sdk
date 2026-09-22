package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SyncGroupListResponse struct {
	bce.BaseResponse
	TotalCount *int32      `json:"totalCount,omitempty"`
	PageNo     *int32      `json:"pageNo,omitempty"`
	PageSize   *int32      `json:"pageSize,omitempty"`
	Result     []*ListItem `json:"result,omitempty"`
}
