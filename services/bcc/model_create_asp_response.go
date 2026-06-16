package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAspResponse struct {
	bce.BaseResponse
	AspId *string `json:"aspId,omitempty"`
}
