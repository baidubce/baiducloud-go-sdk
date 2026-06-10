package bcc

type GetCdsPriceRequest struct {
	StorageType    *string `json:"storageType,omitempty"`
	CdsSizeInGB    *int32  `json:"cdsSizeInGB,omitempty"`
	PaymentTiming  *string `json:"paymentTiming,omitempty"`
	ZoneName       *string `json:"zoneName,omitempty"`
	PurchaseCount  *int32  `json:"purchaseCount,omitempty"`
	PurchaseLength *int32  `json:"purchaseLength,omitempty"`
	CdsExtraIo     *int32  `json:"cdsExtraIo,omitempty"`
}
