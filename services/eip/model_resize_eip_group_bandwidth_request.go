package eip

type ResizeEipGroupBandwidthRequest struct {
	Id              *string `json:"-"`
	ClientToken     *string `json:"-"`
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
}
