package blb

type BlbInquiryRequest struct {
	BlbType          *string  `json:"blbType,omitempty"`
	PerformanceLevel *string  `json:"performanceLevel,omitempty"`
	Count            *int32   `json:"count,omitempty"`
	Billing          *Billing `json:"billing,omitempty"`
}
