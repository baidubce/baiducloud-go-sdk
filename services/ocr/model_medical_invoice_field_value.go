package ocr

type MedicalInvoiceFieldValue struct {
	Word        *string                    `json:"word,omitempty"`
	Location    *MedicalInvoiceLocation    `json:"location,omitempty"`
	Probability *MedicalInvoiceProbability `json:"probability,omitempty"`
}
