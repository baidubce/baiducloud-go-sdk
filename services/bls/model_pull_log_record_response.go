package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type PullLogRecordResponse struct {
	bce.BaseResponse
	Result      []*LogRecord `json:"result,omitempty"`
	Marker      *string      `json:"marker,omitempty"`
	IsTruncated *bool        `json:"isTruncated,omitempty"`
	NextMarker  *string      `json:"nextMarker,omitempty"`
}
