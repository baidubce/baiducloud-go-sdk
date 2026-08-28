package ocr

type FourFactorsVerificationRequest struct {
	Name    *string `json:"name,omitempty"`
	Idcard  *string `json:"idcard,omitempty"`
	Company *string `json:"company,omitempty"`
	Regnum  *string `json:"regnum,omitempty"`
}
