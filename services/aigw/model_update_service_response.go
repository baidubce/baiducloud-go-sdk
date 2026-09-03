package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateServiceResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
