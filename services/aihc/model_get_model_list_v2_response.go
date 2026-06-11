package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetModelListV2Response struct {
	bce.BaseResponse
	TotalCount *int32   `json:"totalCount,omitempty"`
	Models     []*Model `json:"models,omitempty"`
}
