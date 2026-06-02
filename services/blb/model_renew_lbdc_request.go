package blb

type RenewLbdcRequest struct {
	Id          *string          `json:"-"`
	ClientToken *string          `json:"-"`
	Billing     *BillingForRenew `json:"billing,omitempty"`
}
