package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreatePodmonitorResponse struct {
	bce.BaseResponse
	PodMonitorName *string `json:"podMonitorName,omitempty"`
}
