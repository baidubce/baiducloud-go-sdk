package eip

type PurchaseReservedEipGroupRequest struct {
	Id          *string  `json:"-"`
	ClientToken *string  `json:"-"`
	Billing     *Billing `json:"billing,omitempty"`
}
