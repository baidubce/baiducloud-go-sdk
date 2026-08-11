package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListL3MountTargetResponse struct {
	bce.BaseResponse
	RequestId   *string            `json:"requestId,omitempty"`
	IsTruncated *string            `json:"isTruncated,omitempty"`
	Marker      *string            `json:"marker,omitempty"`
	MaxKeys     *string            `json:"maxKeys,omitempty"`
	NextMarker  *string            `json:"nextMarker,omitempty"`
	Result      []*MountTargetInfo `json:"result,omitempty"`
}
