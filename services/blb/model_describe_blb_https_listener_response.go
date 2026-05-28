package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbHttpsListenerResponse struct {
	bce.BaseResponse
	ListenerList []*HTTPSListenerModel `json:"listenerList,omitempty"`
	Marker       *string               `json:"marker,omitempty"`
	IsTruncated  *bool                 `json:"isTruncated,omitempty"`
	NextMarker   *string               `json:"nextMarker,omitempty"`
	MaxKeys      *int32                `json:"maxKeys,omitempty"`
}
