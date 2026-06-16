package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetAspResponse struct {
	bce.BaseResponse
	AutoSnapshotPolicy *AutoSnapshotPolicyModel `json:"autoSnapshotPolicy,omitempty"`
}
