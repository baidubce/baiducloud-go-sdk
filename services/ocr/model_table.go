package ocr

type Table struct {
	Product        *Product `json:"product,omitempty"`
	Quantity       *string  `json:"quantity,omitempty"`
	UnitPrice      *string  `json:"unit_price,omitempty"`
	SubtotalAmount *string  `json:"subtotal_amount,omitempty"`
}
