package ocr

type HealthReportProbability struct {
	Average *float64 `json:"average,omitempty"`
	Min     *float64 `json:"min,omitempty"`
}
