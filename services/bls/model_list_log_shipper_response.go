package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListLogShipperResponse struct {
	bce.BaseResponse
	Success    *bool             `json:"success,omitempty"`
	Code       *string           `json:"code,omitempty"`
	TotalCount *int32            `json:"totalCount,omitempty"`
	Result     []*ShipperSummary `json:"result,omitempty"`
}
