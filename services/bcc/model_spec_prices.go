package bcc

type SpecPrices struct {
	Spec       *string `json:"spec,omitempty"`
	SpecPrice  *string `json:"specPrice,omitempty"`
	Discount   *string `json:"discount,omitempty"`
	TradePrice *string `json:"tradePrice,omitempty"`
	Status     *string `json:"status,omitempty"`
}
