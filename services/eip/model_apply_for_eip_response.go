package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ApplyForEipResponse struct {
	bce.BaseResponse
	Eip *string `json:"eip,omitempty"`
}
