package ocr

type FourFactorsVerificationResult struct {
	Verifyresult *string `json:"verifyresult,omitempty"`
	Namematch    *string `json:"namematch,omitempty"`
	Idnummatch   *string `json:"idnummatch,omitempty"`
	Companymatch *string `json:"companymatch,omitempty"`
	Regnummatch  *string `json:"regnummatch,omitempty"`
}
