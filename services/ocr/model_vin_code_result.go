package ocr

type VinCodeResult struct {
	Location *VinCodeLocation `json:"location,omitempty"`
	Words    *string          `json:"words,omitempty"`
}
