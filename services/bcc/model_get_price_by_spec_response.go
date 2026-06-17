package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPriceBySpecResponse struct {
	bce.BaseResponse
	Price []*SpecIdPrices `json:"price,omitempty"`
}
