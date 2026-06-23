package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetServiceMonitorDetailResponse struct {
	bce.BaseResponse
	ApiVersion *string             `json:"apiVersion,omitempty"`
	Kind       *string             `json:"kind,omitempty"`
	Metadata   *ObjectMeta         `json:"metadata,omitempty"`
	Spec       *ServiceMonitorSpec `json:"spec,omitempty"`
}
