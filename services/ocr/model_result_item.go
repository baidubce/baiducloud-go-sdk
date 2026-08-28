package ocr

type ResultItem struct {
	Word        *string                     `json:"word,omitempty"`
	Probability *MultipleInvoiceProbability `json:"probability,omitempty"`
	Location    *MultipleInvoiceLocation    `json:"location,omitempty"`
}
