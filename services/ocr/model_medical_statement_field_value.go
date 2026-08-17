package ocr

type MedicalStatementFieldValue struct {
	Word        *string                      `json:"word,omitempty"`
	Location    *MedicalStatementLocation    `json:"location,omitempty"`
	Probability *MedicalStatementProbability `json:"probability,omitempty"`
}
