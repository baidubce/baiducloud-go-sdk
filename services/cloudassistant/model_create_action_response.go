package cloudassistant

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateActionResponse struct {
	bce.BaseResponse
	RequestId *string       `json:"requestId,omitempty"`
	Code      *string       `json:"code,omitempty"`
	Message   *string       `json:"message,omitempty"`
	Success   *bool         `json:"success,omitempty"`
	Result    *ActionResult `json:"result,omitempty"`
}
