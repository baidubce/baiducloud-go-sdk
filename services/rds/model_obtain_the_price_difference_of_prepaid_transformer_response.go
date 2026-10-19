package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ObtainThePriceDifferenceOfPrepaidTransformerResponse struct {
	bce.BaseResponse
	Price            *map[string]interface{} `json:"price,omitempty"`
	CatalogPrice     *map[string]interface{} `json:"catalogPrice,omitempty"`
	TrafficInGB      *map[string]interface{} `json:"trafficInGB,omitempty"`
	Discount         *map[string]interface{} `json:"discount,omitempty"`
	RealCatalogPrice *map[string]interface{} `json:"realCatalogPrice,omitempty"`
	BackupInGB       *map[string]interface{} `json:"backupInGB,omitempty"`
	DiscountRate     *int32                  `json:"discountRate,omitempty"`
}
