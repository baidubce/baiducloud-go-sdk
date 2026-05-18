package eip

type ResizeEipBpBandwidthRequest struct {
	Id              *string `json:"-"`
	ClientToken     *string `json:"-"`
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
}
