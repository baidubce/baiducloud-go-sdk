package cloudassistant

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateActionResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	Code      *string `json:"code,omitempty"`
	Message   *string `json:"message,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}
