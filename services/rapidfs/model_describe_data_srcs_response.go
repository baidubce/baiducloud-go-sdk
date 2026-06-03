package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeDataSrcsResponse struct {
	bce.BaseResponse
	DataSrcInfos []*DataSrcInfo `json:"dataSrcInfos,omitempty"`
	Marker       *string        `json:"marker,omitempty"`
	IsTruncated  *bool          `json:"isTruncated,omitempty"`
	NextMarker   *string        `json:"nextMarker,omitempty"`
	MaxKeys      *int32         `json:"maxKeys,omitempty"`
}
