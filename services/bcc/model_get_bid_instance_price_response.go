package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetBidInstancePriceResponse struct {
	bce.BaseResponse
	Money            *string `json:"money,omitempty"`
	Count            *string `json:"count,omitempty"`
	OriginalMoney    *string `json:"originalMoney,omitempty"`
	PerOriginalMoney *string `json:"perOriginalMoney,omitempty"`
	PerMoney         *string `json:"perMoney,omitempty"`
}
