package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAppBlbServerGroupPortResponse struct {
	bce.BaseResponse
	Id     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}
