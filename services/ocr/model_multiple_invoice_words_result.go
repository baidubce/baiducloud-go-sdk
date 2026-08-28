package ocr

type MultipleInvoiceWordsResult struct {
	OcrType     *string                `json:"type,omitempty"`
	Top         *int32                 `json:"top,omitempty"`
	Left        *int32                 `json:"left,omitempty"`
	Width       *int32                 `json:"width,omitempty"`
	Height      *int32                 `json:"height,omitempty"`
	Probability *float32               `json:"probability,omitempty"`
	Result      *MultipleInvoiceResult `json:"result,omitempty"`
}
