package ocr

type WordResult struct {
	WordName    *string                  `json:"word_name,omitempty"`
	Word        *string                  `json:"word,omitempty"`
	Location    *HealthReportLocation    `json:"location,omitempty"`
	Probability *HealthReportProbability `json:"probability,omitempty"`
}
