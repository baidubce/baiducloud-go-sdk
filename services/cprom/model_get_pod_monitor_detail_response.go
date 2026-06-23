package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPodMonitorDetailResponse struct {
	bce.BaseResponse
	ApiVersion *string         `json:"apiVersion,omitempty"`
	Kind       *string         `json:"kind,omitempty"`
	Metadata   *ObjectMeta     `json:"metadata,omitempty"`
	Spec       *PodMonitorSpec `json:"spec,omitempty"`
}
