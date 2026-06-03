package rapidfs

type StockInfo struct {
	Zone             *string `json:"zone,omitempty"`
	StockCapacityTiB *int32  `json:"stockCapacityTiB,omitempty"`
	StockQuantity    *int32  `json:"stockQuantity,omitempty"`
}
