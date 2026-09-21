package vdb

type VdbPricingQueryResponse struct {
	CatalogPrice     *float64         `json:"catalogPrice,omitempty"`
	ComponentType    *string          `json:"componentType,omitempty"`
	Discount         *float64         `json:"discount,omitempty"`
	DiscountType     *string          `json:"discountType,omitempty"`
	MultiYearResult  *MultiYearResult `json:"multiYearResult,omitempty"`
	OriginDiscount   *float64         `json:"originDiscount,omitempty"`
	OriginPrice      *float64         `json:"originPrice,omitempty"`
	Price            *float64         `json:"price,omitempty"`
	PriceId          *int64           `json:"priceId,omitempty"`
	PriceName        *string          `json:"priceName,omitempty"`
	PriceType        *string          `json:"priceType,omitempty"`
	RealCatalogPrice *float64         `json:"realCatalogPrice,omitempty"`
}
