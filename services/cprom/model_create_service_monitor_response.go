package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateServiceMonitorResponse struct {
	bce.BaseResponse
	ServiceMonitorName *string `json:"serviceMonitorName,omitempty"`
}
