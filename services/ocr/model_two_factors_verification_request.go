package ocr

type TwoFactorsVerificationRequest struct {
	Company *string `json:"company,omitempty"`
	Regnum  *string `json:"regnum,omitempty"`
}
