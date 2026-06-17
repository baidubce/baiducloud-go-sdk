package bcm

type DataPoint struct {
	Timestamp *int64   `json:"timestamp,omitempty"`
	Avg       *float32 `json:"avg,omitempty"`
	Max       *float32 `json:"max,omitempty"`
	Min       *float32 `json:"min,omitempty"`
	Sum       *float32 `json:"sum,omitempty"`
	Count     *int32   `json:"count,omitempty"`
}
