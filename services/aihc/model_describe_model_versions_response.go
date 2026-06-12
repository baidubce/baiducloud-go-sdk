package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeModelVersionsResponse struct {
	bce.BaseResponse
	TotalCount *int32               `json:"totalCount,omitempty"`
	Versions   []*ModelVersionEntry `json:"versions,omitempty"`
}
