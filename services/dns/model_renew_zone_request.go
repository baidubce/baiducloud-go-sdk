package dns

type RenewZoneRequest struct {
	Name        *string          `json:"-"`
	Action      *string          `json:"-"`
	ClientToken *string          `json:"-"`
	Billing     *BillingForRenew `json:"billing,omitempty"`
}
