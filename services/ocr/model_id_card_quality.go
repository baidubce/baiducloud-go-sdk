package ocr

type IdCardQuality struct {
	IsClear               *int32   `json:"IsClear,omitempty"`
	IsClearPropobility    *float64 `json:"IsClear_propobility,omitempty"`
	IsComplete            *int32   `json:"IsComplete,omitempty"`
	IsCompletePropobility *float64 `json:"IsComplete_propobility,omitempty"`
	IsNoCover             *int32   `json:"IsNoCover,omitempty"`
	IsNoCoverPropobility  *float64 `json:"IsNoCover_propobility,omitempty"`
}
