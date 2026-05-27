package csn

type QueryCsnBpPriceRequest struct {
	Name        *string  `json:"name,omitempty"`
	Bandwidth   *int32   `json:"bandwidth,omitempty"`
	GeographicA *string  `json:"geographicA,omitempty"`
	GeographicB *string  `json:"geographicB,omitempty"`
	Billing     *Billing `json:"billing,omitempty"`
}
