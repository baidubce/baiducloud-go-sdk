package bcc

type GetReservedInstancePriceRequest struct {
	SpecId                *string `json:"specId,omitempty"`
	Spec                  *string `json:"spec,omitempty"`
	OfferingType          *string `json:"offeringType,omitempty"`
	Scope                 *string `json:"scope,omitempty"`
	ZoneName              *string `json:"zoneName,omitempty"`
	ReservedInstanceCount *int32  `json:"reservedInstanceCount,omitempty"`
	PriceTimeUnit         *string `json:"priceTimeUnit,omitempty"`
	ReservedInstanceTime  *int32  `json:"reservedInstanceTime,omitempty"`
	PurchaseNum           *int32  `json:"purchaseNum,omitempty"`
}
