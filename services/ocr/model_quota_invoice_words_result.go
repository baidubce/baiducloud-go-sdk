package ocr

type QuotaInvoiceWordsResult struct {
	InvoiceCode          *string `json:"invoice_code,omitempty"`
	InvoiceNumber        *string `json:"invoice_number,omitempty"`
	InvoiceRate          *string `json:"invoice_rate,omitempty"`
	Location             *string `json:"location,omitempty"`
	InvoiceRateLowercase *string `json:"invoice_rate_lowercase,omitempty"`
	Province             *string `json:"province,omitempty"`
	City                 *string `json:"city,omitempty"`
}
