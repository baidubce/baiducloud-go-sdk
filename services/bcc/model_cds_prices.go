package bcc

type CdsPrices struct {
	StorageType *string  `json:"storageType,omitempty"`
	CdsSizeInGB *int32   `json:"cdsSizeInGB,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	SpecPrice   *float64 `json:"specPrice,omitempty"`
	Unit        *string  `json:"unit,omitempty"`
}
