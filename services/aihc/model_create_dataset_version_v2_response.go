package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDatasetVersionV2Response struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
