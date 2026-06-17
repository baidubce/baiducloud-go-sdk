package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteAlarmMaskingsResponse struct {
	bce.BaseResponse
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}
