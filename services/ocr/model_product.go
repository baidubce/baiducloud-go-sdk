package ocr

type Product struct {
	Word        *string                     `json:"word,omitempty"`
	Location    *ShoppingReceiptLocation    `json:"location,omitempty"`
	Probability *ShoppingReceiptProbability `json:"probability,omitempty"`
}
