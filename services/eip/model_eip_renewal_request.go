package eip

type EipRenewalRequest struct {
	Eip         *string  `json:"-"`
	ClientToken *string  `json:"-"`
	Billing     *Billing `json:"billing,omitempty"`
}
