package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAlertResponse struct {
	bce.BaseResponse
	AlertId *string `json:"alertId,omitempty"`
}
