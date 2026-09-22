package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPriceForCreateInstanceResponse struct {
	bce.BaseResponse
	Price        *float32 `json:"price,omitempty"`
	CatalogPrice *float32 `json:"catalogPrice,omitempty"`
}
