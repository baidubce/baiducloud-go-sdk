package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListFilesetResponse struct {
	bce.BaseResponse
	RequestId     *string         `json:"requestId,omitempty"`
	Result        []*FilesetModel `json:"result,omitempty"`
	PageNo        *int32          `json:"pageNo,omitempty"`
	PageSize      *int32          `json:"pageSize,omitempty"`
	TotalCount    *int32          `json:"totalCount,omitempty"`
	MaxFilesetNum *int32          `json:"maxFilesetNum,omitempty"`
	MaxFilesQuota *int64          `json:"maxFilesQuota,omitempty"`
	MinFilesQuota *int64          `json:"minFilesQuota,omitempty"`
}
