package bcc

type GetPriceBySpecRequest struct {
	SpecId         *string `json:"specId,omitempty"`
	Spec           *string `json:"spec,omitempty"`
	PaymentTiming  *string `json:"paymentTiming,omitempty"`
	ZoneName       *string `json:"zoneName,omitempty"`
	PurchaseCount  *int32  `json:"purchaseCount,omitempty"`
	PurchaseLength *int32  `json:"purchaseLength,omitempty"`
}
