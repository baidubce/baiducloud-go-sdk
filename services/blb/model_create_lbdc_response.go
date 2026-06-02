package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateLbdcResponse struct {
	bce.BaseResponse
	Id      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	BlbType *string `json:"type,omitempty"`
	Desc    *string `json:"desc,omitempty"`
}
