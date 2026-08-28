package ocr

type DrivingLicenseQualityPropobility struct {
	IsClearPropobility    *float32 `json:"is_clear_propobility,omitempty"`
	IsCompletePropobility *float32 `json:"is_complete_propobility,omitempty"`
	IsNoshieldPropobility *string  `json:"is_noshield_propobility,omitempty"`
}
