package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryCsnBpPriceResponse struct {
	bce.BaseResponse
	Price *string `json:"price,omitempty"`
}
