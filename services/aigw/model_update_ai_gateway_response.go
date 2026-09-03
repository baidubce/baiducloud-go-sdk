package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateAIGatewayResponse struct {
	bce.BaseResponse
	InstanceId       *string `json:"instanceId,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	DeleteProtection *bool   `json:"deleteProtection,omitempty"`
	PublicAccessible *bool   `json:"publicAccessible,omitempty"`
	Replicas         *int32  `json:"replicas,omitempty"`
	UpdateTime       *string `json:"updateTime,omitempty"`
}
