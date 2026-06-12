package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDatasetsResponse struct {
	bce.BaseResponse
	TotalCount *int32     `json:"totalCount,omitempty"`
	Datasets   []*Dataset `json:"datasets,omitempty"`
}
