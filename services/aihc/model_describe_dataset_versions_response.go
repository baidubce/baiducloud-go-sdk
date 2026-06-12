package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDatasetVersionsResponse struct {
	bce.BaseResponse
	TotalCount *int32                 `json:"totalCount,omitempty"`
	Versions   []*DatasetVersionEntry `json:"versions,omitempty"`
}
