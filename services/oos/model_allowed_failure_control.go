package oos

type AllowedFailureControl struct {
	Ratio *float64 `json:"ratio,omitempty"`
	Count *int32   `json:"count,omitempty"`
}
