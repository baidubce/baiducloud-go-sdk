package bcc

type SpecIdPrices struct {
	SpecId     *string       `json:"specId,omitempty"`
	SpecPrices []*SpecPrices `json:"specPrices,omitempty"`
}
