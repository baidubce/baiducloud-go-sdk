package eip

type EipMoveOutModel struct {
	Eip             *string  `json:"eip,omitempty"`
	BandwidthInMbps *int32   `json:"bandwidthInMbps,omitempty"`
	Billing         *Billing `json:"billing,omitempty"`
}
