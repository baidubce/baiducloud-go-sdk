package ocr

type MedicalDetailProbability struct {
	Average *float64 `json:"average,omitempty"`
	Min     *float64 `json:"min,omitempty"`
}
