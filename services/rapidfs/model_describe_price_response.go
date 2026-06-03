package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribePriceResponse struct {
	bce.BaseResponse
	PriceInfos []*PriceInfo `json:"priceInfos,omitempty"`
}
