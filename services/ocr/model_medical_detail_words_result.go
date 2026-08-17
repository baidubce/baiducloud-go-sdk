package ocr

type MedicalDetailWordsResult struct {
	Name        *FieldValue         `json:"Name,omitempty"`
	Date        *FieldValue         `json:"Date,omitempty"`
	PatientID   *FieldValue         `json:"PatientID,omitempty"`
	TotalAmount *FieldValue         `json:"TotalAmount,omitempty"`
	CostDetail  [][]*CostDetailItem `json:"CostDetail,omitempty"`
}
