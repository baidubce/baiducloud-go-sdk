package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateBlbResponse struct {
	bce.BaseResponse
	Address *string `json:"address,omitempty"`
	Name    *string `json:"name,omitempty"`
	BlbId   *string `json:"blbId,omitempty"`
	Desc    *string `json:"desc,omitempty"`
}
