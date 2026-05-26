package vpc

type UpdatePeerConnBandwidthRequest struct {
	PeerConnId         *string `json:"-"`
	ClientToken        *string `json:"-"`
	NewBandwidthInMbps *int32  `json:"newBandwidthInMbps,omitempty"`
}
