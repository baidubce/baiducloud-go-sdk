package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetReservedInstancePriceResponse struct {
	bce.BaseResponse
	RequestId     *string        `json:"requestId,omitempty"`
	Spec          *string        `json:"spec,omitempty"`
	CategoryPrice *CatagoryPrice `json:"categoryPrice,omitempty"`
	TradePrice    *TradePrice    `json:"tradePrice,omitempty"`
}
