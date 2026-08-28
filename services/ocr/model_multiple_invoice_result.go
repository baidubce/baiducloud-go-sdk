package ocr

type MultipleInvoiceResult struct {
	InvoiceCode          []*WordOnlyItem `json:"invoice_code,omitempty"`
	InvoiceNum           []*ResultItem   `json:"invoice_num,omitempty"`
	InvoiceDate          []*ResultItem   `json:"invoice_date,omitempty"`
	TotalAmount          []*ResultItem   `json:"total_amount,omitempty"`
	InvoiceType          []*WordOnlyItem `json:"invoice_type,omitempty"`
	CheckCode            []*WordOnlyItem `json:"check_code,omitempty"`
	SellerName           []*ResultItem   `json:"seller_name,omitempty"`
	SellerRegisterNum    []*ResultItem   `json:"seller_register_num,omitempty"`
	PurchaserName        []*ResultItem   `json:"purchaser_name,omitempty"`
	PurchaserRegisterNum []*ResultItem   `json:"purchaser_register_num,omitempty"`
}
