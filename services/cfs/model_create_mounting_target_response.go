package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateMountingTargetResponse struct {
	bce.BaseResponse
	Domain  *string `json:"domain,omitempty"`
	MountId *string `json:"mountId,omitempty"`
}
