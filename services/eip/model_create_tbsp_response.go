package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateTbspResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
