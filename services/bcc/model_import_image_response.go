package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ImportImageResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
