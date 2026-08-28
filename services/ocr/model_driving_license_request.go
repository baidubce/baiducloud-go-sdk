package ocr

type DrivingLicenseRequest struct {
	Image              *string `json:"image,omitempty"`
	Url                *string `json:"url,omitempty"`
	DetectDirection    *bool   `json:"detect_direction,omitempty"`
	DrivingLicenseSide *string `json:"driving_license_side,omitempty"`
	UnifiedValidPeriod *bool   `json:"unified_valid_period,omitempty"`
	QualityWarn        *bool   `json:"quality_warn,omitempty"`
	RiskWarn           *bool   `json:"risk_warn,omitempty"`
}
