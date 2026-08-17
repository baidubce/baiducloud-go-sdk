package ocr

type MedicalReportDetectionWordsResult struct {
	CommonData []*MedicalReportDetectionCommonDataItem `json:"CommonData,omitempty"`
	Item       [][]*MedicalReportDetectionItemField    `json:"Item,omitempty"`
}
