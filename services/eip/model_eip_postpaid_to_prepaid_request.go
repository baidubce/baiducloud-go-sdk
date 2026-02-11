package eip

type EipPostpaidToPrepaidRequest struct {
	Eip            *string `json:"-"`
	ClientToken    *string `json:"-"`
	PurchaseLength *int32  `json:"purchaseLength,omitempty"`
	BandWidth      *int32  `json:"bandWidth,omitempty"`
}
