package ocr

type CostDetailItem struct {
	WordName    *string                   `json:"word_name,omitempty"`
	Word        *string                   `json:"word,omitempty"`
	Location    *MedicalLocation          `json:"location,omitempty"`
	Probability *MedicalDetailProbability `json:"probability,omitempty"`
}
