package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateConsumerResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Status  *int32  `json:"status,omitempty"`
	Result  *string `json:"result,omitempty"`
}
