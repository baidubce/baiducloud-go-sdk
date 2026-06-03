package rapidfs

type DescribePriceRequest struct {
	Action      *string `json:"-"`
	Currency    *string `json:"currency,omitempty"`
	ManagedMode *string `json:"managedMode,omitempty"`
	MetaSpec    *string `json:"metaSpec,omitempty"`
	DataSpec    *string `json:"dataSpec,omitempty"`
	CapacityTiB *int32  `json:"capacityTiB,omitempty"`
}
