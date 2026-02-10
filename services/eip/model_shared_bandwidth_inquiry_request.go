package eip

type SharedBandwidthInquiryRequest struct {
	BandwidthInMbps *int32   `json:"bandwidthInMbps,omitempty"`
	Count           *int32   `json:"count,omitempty"`
	IpNum           *int32   `json:"ipNum,omitempty"`
	Billing         *Billing `json:"billing,omitempty"`
}
