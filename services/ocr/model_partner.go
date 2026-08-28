package ocr

type Partner struct {
	Stockname        *string `json:"stockname,omitempty"`
	Stocktype        *string `json:"stocktype,omitempty"`
	Stockpercent     *string `json:"stockpercent,omitempty"`
	Stockcapital     *string `json:"stockcapital,omitempty"`
	Shouddate        *string `json:"shouddate,omitempty"`
	Investtype       *string `json:"investtype,omitempty"`
	Stockrealcapital *string `json:"stockrealcapital,omitempty"`
	Capidate         *string `json:"capidate,omitempty"`
	Investname       *string `json:"investname,omitempty"`
	Concur           *string `json:"concur,omitempty"`
}
