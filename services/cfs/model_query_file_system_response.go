package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryFileSystemResponse struct {
	bce.BaseResponse
	FileSystemList []*FileSystemModel `json:"fileSystemList,omitempty"`
	Marker         *string            `json:"marker,omitempty"`
	IsTruncated    *bool              `json:"isTruncated,omitempty"`
	NextMarker     *string            `json:"nextMarker,omitempty"`
	MaxKeys        *int32             `json:"maxKeys,omitempty"`
}
