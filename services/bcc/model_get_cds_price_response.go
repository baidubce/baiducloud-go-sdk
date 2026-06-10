package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetCdsPriceResponse struct {
	bce.BaseResponse
	Price []*CdsPrices `json:"price,omitempty"`
}
