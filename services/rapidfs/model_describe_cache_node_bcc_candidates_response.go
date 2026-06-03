package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeCacheNodeBccCandidatesResponse struct {
	bce.BaseResponse
	BccInstanceInfos []*BccInstanceInfo `json:"bccInstanceInfos,omitempty"`
	Marker           *string            `json:"marker,omitempty"`
	IsTruncated      *bool              `json:"isTruncated,omitempty"`
	NextMarker       *string            `json:"nextMarker,omitempty"`
	MaxKeys          *int32             `json:"maxKeys,omitempty"`
}
