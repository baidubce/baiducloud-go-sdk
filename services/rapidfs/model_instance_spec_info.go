package rapidfs

type InstanceSpecInfo struct {
	ManagedMode     *string      `json:"managedMode,omitempty"`
	MetaSpec        *string      `json:"metaSpec,omitempty"`
	DataSpec        *string      `json:"dataSpec,omitempty"`
	MinCapacityTiB  *int32       `json:"minCapacityTiB,omitempty"`
	StepCapacityTiB *int32       `json:"stepCapacityTiB,omitempty"`
	MaxCapacityTiB  *int32       `json:"maxCapacityTiB,omitempty"`
	StockInfos      []*StockInfo `json:"stockInfos,omitempty"`
}
