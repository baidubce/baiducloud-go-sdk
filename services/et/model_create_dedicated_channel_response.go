package et

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDedicatedChannelResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
