package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbSslListenerResponse struct {
	bce.BaseResponse
	ListenerList []*SSLListenerModel `json:"listenerList,omitempty"`
	Marker       *string             `json:"marker,omitempty"`
	IsTruncated  *bool               `json:"isTruncated,omitempty"`
	NextMarker   *string             `json:"nextMarker,omitempty"`
	MaxKeys      *int32              `json:"maxKeys,omitempty"`
}
