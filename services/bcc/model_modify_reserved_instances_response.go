package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ModifyReservedInstancesResponse struct {
	bce.BaseResponse
	ModifyReservedInstanceOrders []*ModifyReservedInstanceOrder `json:"modifyReservedInstanceOrders,omitempty"`
}
