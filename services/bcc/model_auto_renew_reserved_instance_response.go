package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AutoRenewReservedInstanceResponse struct {
	bce.BaseResponse
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
}
