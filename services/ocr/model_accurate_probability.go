package ocr

type AccurateProbability struct {
	Average  *float64 `json:"average,omitempty"`
	Variance *float64 `json:"variance,omitempty"`
	Min      *float64 `json:"min,omitempty"`
}
