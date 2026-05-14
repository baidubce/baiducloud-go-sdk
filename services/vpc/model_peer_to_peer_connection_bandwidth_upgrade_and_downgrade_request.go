package vpc

type PeerToPeerConnectionBandwidthUpgradeAndDowngradeRequest struct {
	PeerConnId         *string `json:"-"`
	ClientToken        *string `json:"-"`
	NewBandwidthInMbps *int32  `json:"newBandwidthInMbps,omitempty"`
}
