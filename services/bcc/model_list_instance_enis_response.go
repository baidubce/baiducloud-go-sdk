package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListInstanceEnisResponse struct {
	bce.BaseResponse
	Enis []*EniInfo `json:"enis,omitempty"`
}
