package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAppBlbResponse struct {
	bce.BaseResponse
	Address *string `json:"address,omitempty"`
	Name    *string `json:"name,omitempty"`
	Desc    *string `json:"desc,omitempty"`
	BlbId   *string `json:"blbId,omitempty"`
}
