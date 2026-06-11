package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetAListOfModelVersionsV2Response struct {
	bce.BaseResponse
	TotalCount *int32               `json:"totalCount,omitempty"`
	Versions   []*ModelVersionEntry `json:"versions,omitempty"`
}
