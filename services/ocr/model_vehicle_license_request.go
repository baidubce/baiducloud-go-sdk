package ocr

type VehicleLicenseRequest struct {
	Image              *string `json:"image,omitempty"`
	Url                *string `json:"url,omitempty"`
	DetectDirection    *bool   `json:"detect_direction,omitempty"`
	VehicleLicenseSide *string `json:"vehicle_license_side,omitempty"`
	Unified            *bool   `json:"unified,omitempty"`
	QualityWarn        *bool   `json:"quality_warn,omitempty"`
	RiskWarn           *bool   `json:"risk_warn,omitempty"`
}
