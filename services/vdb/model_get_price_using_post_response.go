package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPriceUsingPOSTResponse struct {
	bce.BaseResponse
	CatalogPrice     *float64                   `json:"catalogPrice,omitempty"`
	Discount         *float64                   `json:"discount,omitempty"`
	DiscountType     *string                    `json:"discountType,omitempty"`
	Price            *float64                   `json:"price,omitempty"`
	PriceDetails     []*VdbPricingQueryResponse `json:"priceDetails,omitempty"`
	RealCatalogPrice *float64                   `json:"realCatalogPrice,omitempty"`
}
