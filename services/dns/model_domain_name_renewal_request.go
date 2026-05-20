package dns

type DomainNameRenewalRequest struct {
	Name        *string            `json:"-"`
	Action      *string            `json:"-"`
	ClientToken *string            `json:"-"`
	Billing     []*BillingForRenew `json:"billing,omitempty"`
}
