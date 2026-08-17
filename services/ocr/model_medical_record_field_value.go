package ocr

type MedicalRecordFieldValue struct {
	Word        *string                   `json:"word,omitempty"`
	Location    *MedicalRecordLocation    `json:"location,omitempty"`
	Probability *MedicalRecordProbability `json:"probability,omitempty"`
}
