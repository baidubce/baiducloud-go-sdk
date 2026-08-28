package ocr

type TollInvoiceWordsResult struct {
	InvoiceCode *string `json:"InvoiceCode,omitempty"`
	InvoiceNum  *string `json:"InvoiceNum,omitempty"`
	Entrance    *string `json:"Entrance,omitempty"`
	Exit        *string `json:"Exit,omitempty"`
	Date        *string `json:"Date,omitempty"`
	Time        *string `json:"Time,omitempty"`
	Fare        *string `json:"Fare,omitempty"`
	Province    *string `json:"Province,omitempty"`
	City        *string `json:"City,omitempty"`
}
