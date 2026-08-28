package ocr

type CardQuality struct {
	IsClear               *int32   `json:"IsClear,omitempty"`
	IsClearProbability    *float64 `json:"IsClear_probability,omitempty"`
	IsComplete            *int32   `json:"IsComplete,omitempty"`
	IsCompleteProbability *float64 `json:"IsComplete_probability,omitempty"`
}
