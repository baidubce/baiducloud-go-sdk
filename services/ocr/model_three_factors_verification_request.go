package ocr

type ThreeFactorsVerificationRequest struct {
	Name    *string `json:"name,omitempty"`
	Company *string `json:"company,omitempty"`
	Regnum  *string `json:"regnum,omitempty"`
}
