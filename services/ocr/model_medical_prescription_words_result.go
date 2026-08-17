package ocr

type MedicalPrescriptionWordsResult struct {
	CommonData []*MedicalPrescriptionCommonDataItem   `json:"CommonData,omitempty"`
	CostDetail [][]*MedicalPrescriptionCostDetailItem `json:"CostDetail,omitempty"`
}
