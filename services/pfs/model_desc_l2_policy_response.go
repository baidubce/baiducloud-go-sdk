package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescL2PolicyResponse struct {
	bce.BaseResponse
	RequestId   *string `json:"requestId,omitempty"`
	PolicyId    *string `json:"policyId,omitempty"`
	PolicyName  *string `json:"policyName,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty"`
	Path        *string `json:"path,omitempty"`
	ExpiredTime *int32  `json:"expiredTime,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	ExecuteTime *int32  `json:"executeTime,omitempty"`
	PfsType     *int32  `json:"type,omitempty"`
	BosPath     *string `json:"bosPath,omitempty"`
	Status      *int32  `json:"status,omitempty"`
}
