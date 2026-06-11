package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateADatasetV2Response struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
