package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListL2PolicyResponse struct {
	bce.BaseResponse
	RequestId   *string       `json:"requestId,omitempty"`
	IsTruncated *bool         `json:"isTruncated,omitempty"`
	MaxKeys     *int32        `json:"maxKeys,omitempty"`
	Marker      *string       `json:"marker,omitempty"`
	NextMarker  *string       `json:"nextMarker,omitempty"`
	Result      []*PolicyInfo `json:"result,omitempty"`
}
