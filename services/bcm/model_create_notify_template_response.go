package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateNotifyTemplateResponse struct {
	bce.BaseResponse
	Success  *bool   `json:"success,omitempty"`
	Code     *string `json:"code,omitempty"`
	Message  *string `json:"message,omitempty"`
	NotifyId *string `json:"notifyId,omitempty"`
}
