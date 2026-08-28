package ocr

type BankcardResult struct {
	BankCardNumber         *string                 `json:"bank_card_number,omitempty"`
	ValidDate              *string                 `json:"valid_date,omitempty"`
	BankCardType           *int32                  `json:"bank_card_type,omitempty"`
	BankName               *string                 `json:"bank_name,omitempty"`
	HolderName             *string                 `json:"holder_name,omitempty"`
	BankCardNumberLocation *BankCardNumberLocation `json:"bank_card_number_location,omitempty"`
	CardQuality            *CardQuality            `json:"card_quality,omitempty"`
}
