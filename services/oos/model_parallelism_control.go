package oos

type ParallelismControl struct {
	Ratio *float64 `json:"ratio,omitempty"`
	Count *int32   `json:"count,omitempty"`
}
