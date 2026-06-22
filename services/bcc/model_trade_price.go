package bcc

type TradePrice struct {
	PrePayTradePrice  *float64 `json:"prePayTradePrice,omitempty"`
	PostPayTradePrice *float64 `json:"postPayTradePrice,omitempty"`
}
