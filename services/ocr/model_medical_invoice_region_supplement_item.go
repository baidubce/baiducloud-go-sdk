package ocr

type MedicalInvoiceRegionSupplementItem struct {
	Name        *string                    `json:"name,omitempty"`
	Word        *string                    `json:"word,omitempty"`
	Probability *MedicalInvoiceProbability `json:"probability,omitempty"`
	Position    *MedicalInvoicePosition    `json:"position,omitempty"`
}
