package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetServiceListResponse struct {
	bce.BaseResponse
	ServiceName *string `json:"serviceName,omitempty"`
	Namespace   *string `json:"namespace,omitempty"`
	ClusterId   *string `json:"clusterId,omitempty"`
}
