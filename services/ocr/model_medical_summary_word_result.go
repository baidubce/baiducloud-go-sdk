package ocr

type MedicalSummaryWordResult struct {
	WordName    *string                    `json:"word_name,omitempty"`
	Word        *string                    `json:"word,omitempty"`
	Location    *MedicalSummaryLocation    `json:"location,omitempty"`
	Probability *MedicalSummaryProbability `json:"probability,omitempty"`
}
