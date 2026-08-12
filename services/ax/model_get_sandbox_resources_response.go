package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSandboxResourcesResponse struct {
	bce.BaseResponse
	SandboxID   *string                           `json:"sandboxID,omitempty"`
	RuntimeType *string                           `json:"runtimeType,omitempty"`
	Status      *string                           `json:"status,omitempty"`
	Containers  []*SandboxContainerResourceStatus `json:"containers,omitempty"`
	Conditions  []*SandboxResourceCondition       `json:"conditions,omitempty"`
	Message     *string                           `json:"message,omitempty"`
}
