package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRecordResponse struct {
	bce.BaseResponse
	Marker      *string          `json:"marker,omitempty"`
	IsTruncated *bool            `json:"isTruncated,omitempty"`
	NextMarker  *string          `json:"nextMarker,omitempty"`
	MaxKeys     *int32           `json:"maxKeys,omitempty"`
	Records     []*PrivateRecord `json:"records,omitempty"`
}
