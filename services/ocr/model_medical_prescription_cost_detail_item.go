package ocr

type MedicalPrescriptionCostDetailItem struct {
	WordName    *string                         `json:"word_name,omitempty"`
	Word        *string                         `json:"word,omitempty"`
	Location    *MedicalPrescriptionLocation    `json:"location,omitempty"`
	Probability *MedicalPrescriptionProbability `json:"probability,omitempty"`
}
