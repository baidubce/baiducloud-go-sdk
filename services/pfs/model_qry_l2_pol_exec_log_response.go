package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QryL2PolExecLogResponse struct {
	bce.BaseResponse
	RequestId    *string        `json:"requestId,omitempty"`
	PolicyId     *string        `json:"policyId,omitempty"`
	InstanceId   *string        `json:"instanceId,omitempty"`
	ExecuteInfos []*ExecuteInfo `json:"executeInfos,omitempty"`
}
