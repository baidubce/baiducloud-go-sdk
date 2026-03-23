package blb

type Price struct {
	ChargeItem    *string `json:"chargeItem,omitempty"`
	DiscountPrice *string `json:"discountPrice,omitempty"`
	OriginalPrice *string `json:"originalPrice,omitempty"`
	ChargeUnit    *string `json:"chargeUnit,omitempty"`
}
