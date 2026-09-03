package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAIGatewayResponse struct {
	bce.BaseResponse
	InstanceId      *string `json:"instanceId,omitempty"`
	RequestId       *string `json:"requestId,omitempty"`
	TaskId          *string `json:"taskId,omitempty"`
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
}
