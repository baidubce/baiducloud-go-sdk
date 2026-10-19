package rds

type GetNewPurchasePriceRequest struct {
	Instance    *Instance `json:"instance,omitempty"`
	Duration    *int32    `json:"duration,omitempty"`
	Number      *int32    `json:"number,omitempty"`
	ProductType *string   `json:"productType,omitempty"`
}
