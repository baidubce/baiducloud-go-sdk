package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetClusterBindStatusResponse struct {
	bce.BaseResponse
	BindingStatus *string `json:"bindingStatus,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty"`
	AgentId       *string `json:"agentId,omitempty"`
}
