package eip

type BandwidthPackageInquiryRequest struct {
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
	Count           *int32  `json:"count,omitempty"`
	Type            *string `json:"type,omitempty"`
}
