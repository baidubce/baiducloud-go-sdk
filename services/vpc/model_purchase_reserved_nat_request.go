package vpc

type PurchaseReservedNatRequest struct {
	NatId       *string  `json:"-"`
	ClientToken *string  `json:"-"`
	Billing     *Billing `json:"billing,omitempty"`
}
