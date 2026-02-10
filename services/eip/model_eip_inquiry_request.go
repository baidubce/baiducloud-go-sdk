package eip

type EipInquiryRequest struct {
	BandwidthInMbps *int32   `json:"bandwidthInMbps,omitempty"`
	Count           *int32   `json:"count,omitempty"`
	PurchaseType    *string  `json:"purchaseType,omitempty"`
	Billing         *Billing `json:"billing,omitempty"`
}
