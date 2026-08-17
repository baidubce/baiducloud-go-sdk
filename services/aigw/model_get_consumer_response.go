package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetConsumerResponse struct {
	bce.BaseResponse
	Success  *bool               `json:"success,omitempty"`
	Status   *int32              `json:"status,omitempty"`
	Consumer *ConsumerDetailInfo `json:"consumer,omitempty"`
}
