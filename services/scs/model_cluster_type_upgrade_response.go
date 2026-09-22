package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ClusterTypeUpgradeResponse struct {
	bce.BaseResponse
	OrderId *string `json:"orderId,omitempty"`
}
