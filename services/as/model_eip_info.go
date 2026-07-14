package as

type EipInfo struct {
	IfBindEip       *bool   `json:"ifBindEip,omitempty"`
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
	EipProductType  *string `json:"eipProductType,omitempty"`
	PurchaseType    *string `json:"purchaseType,omitempty"`
}
