package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListSnapchainResponse struct {
	bce.BaseResponse
	OrderBy     *string           `json:"orderBy,omitempty"`
	TotalCount  *int32            `json:"totalCount,omitempty"`
	PageSize    *int32            `json:"pageSize,omitempty"`
	PageNo      *int32            `json:"pageNo,omitempty"`
	IsTruncated *bool             `json:"isTruncated,omitempty"`
	Snapchains  []*SnapchainModel `json:"snapchains,omitempty"`
}
