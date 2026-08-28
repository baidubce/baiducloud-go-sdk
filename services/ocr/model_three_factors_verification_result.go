package ocr

type ThreeFactorsVerificationResult struct {
	Verifyresult *string `json:"verifyresult,omitempty"`
	Namematch    *string `json:"namematch,omitempty"`
	Companymatch *string `json:"companymatch,omitempty"`
	Regnummatch  *string `json:"regnummatch,omitempty"`
}
