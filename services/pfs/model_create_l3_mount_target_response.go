package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateL3MountTargetResponse struct {
	bce.BaseResponse
	RequestId     *string `json:"requestId,omitempty"`
	Domain        *string `json:"domain,omitempty"`
	MountTargetId *string `json:"mountTargetId,omitempty"`
}
