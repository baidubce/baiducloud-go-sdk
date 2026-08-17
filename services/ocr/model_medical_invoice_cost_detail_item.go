package ocr

type MedicalInvoiceCostDetailItem struct {
	Name     *string   `json:"name,omitempty"`
	Word     *string   `json:"word,omitempty"`
	MediInfo *MediInfo `json:"medi_info,omitempty"`
}
