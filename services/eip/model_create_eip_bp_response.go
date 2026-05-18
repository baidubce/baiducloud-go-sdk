package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEipBpResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
