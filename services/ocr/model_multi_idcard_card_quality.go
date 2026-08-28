package ocr

type MultiIdcardCardQuality struct {
	IsClear               *int32   `json:"IsClear,omitempty"`
	IsComplete            *int32   `json:"IsComplete,omitempty"`
	IsNoCover             *int32   `json:"IsNoCover,omitempty"`
	IsClearPropobility    *float32 `json:"IsClear_propobility,omitempty"`
	IsCompletePropobility *float32 `json:"IsComplete_propobility,omitempty"`
	IsNoCoverPropobility  *float32 `json:"IsNoCover_propobility,omitempty"`
}
