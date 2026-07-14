package as

type AsEip struct {
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
	Address         *string `json:"address,omitempty"`
	EipId           *string `json:"eipId,omitempty"`
	EipStatus       *string `json:"eipStatus,omitempty"`
	EipAllocationId *string `json:"eipAllocationId,omitempty"`
}
