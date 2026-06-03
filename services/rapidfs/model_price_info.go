package rapidfs

type PriceInfo struct {
	Currency   *string  `json:"currency,omitempty"`
	UnitPrice  *float64 `json:"unitPrice,omitempty"`
	ChargeType *string  `json:"chargeType,omitempty"`
	ChargeUnit *string  `json:"chargeUnit,omitempty"`
}
