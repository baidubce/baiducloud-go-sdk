package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDatasetResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
