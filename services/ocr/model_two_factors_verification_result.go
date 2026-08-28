package ocr

type TwoFactorsVerificationResult struct {
	Verifyresult *string `json:"verifyresult,omitempty"`
	Companymatch *string `json:"companymatch,omitempty"`
	Regnummatch  *string `json:"regnummatch,omitempty"`
}
