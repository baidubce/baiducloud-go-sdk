package ocr

type MedicalReportDetectionCommonDataItem struct {
	WordName    *string                            `json:"word_name,omitempty"`
	Word        *string                            `json:"word,omitempty"`
	Location    *MedicalReportDetectionLocation    `json:"location,omitempty"`
	Probability *MedicalReportDetectionProbability `json:"probability,omitempty"`
}
