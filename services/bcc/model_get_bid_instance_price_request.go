package bcc

type GetBidInstancePriceRequest struct {
	Spec                  *string           `json:"spec,omitempty"`
	RootDiskSizeInGb      *int32            `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   *string           `json:"rootDiskStorageType,omitempty"`
	CreateCdsList         []*CreateCdsModel `json:"createCdsList,omitempty"`
	NetworkCapacityInMbps *int32            `json:"networkCapacityInMbps,omitempty"`
	InternetChargeType    *string           `json:"internetChargeType,omitempty"`
	PurchaseCount         *int32            `json:"purchaseCount,omitempty"`
	ZoneName              *string           `json:"zoneName,omitempty"`
}
