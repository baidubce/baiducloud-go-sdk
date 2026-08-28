package ocr

type BusinessLicenseQuality struct {
	IsClear               *int32   `json:"is_clear,omitempty"`
	IsClearPropobility    *float64 `json:"is_clear_propobility,omitempty"`
	IsComplete            *int32   `json:"is_complete,omitempty"`
	IsCompletePropobility *float64 `json:"is_complete_propobility,omitempty"`
}
