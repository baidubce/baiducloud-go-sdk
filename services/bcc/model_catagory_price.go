package bcc

type CatagoryPrice struct {
	PrePayCategoryPrice  *float64 `json:"prePayCategoryPrice,omitempty"`
	PostPayCategoryPrice *float64 `json:"postPayCategoryPrice,omitempty"`
}
