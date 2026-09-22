package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GePriceForResizeInstanceResponse struct {
	bce.BaseResponse
	Price *float32 `json:"price,omitempty"`
}
