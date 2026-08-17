package ocr

type MedicalInvoiceRequest struct {
	Image       *string `json:"image,omitempty"`
	Url         *string `json:"url,omitempty"`
	Location    *bool   `json:"location,omitempty"`
	Probability *bool   `json:"probability,omitempty"`
	MediQuery   *string `json:"medi_query,omitempty"`
}
