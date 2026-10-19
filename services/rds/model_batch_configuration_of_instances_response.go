package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchConfigurationOfInstancesResponse struct {
	bce.BaseResponse
	OrderList []*string `json:"orderList,omitempty"`
}
