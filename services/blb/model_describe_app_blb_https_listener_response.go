package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAppBlbHttpsListenerResponse struct {
	bce.BaseResponse
	ListenerList []*AppHTTPSListenerModel `json:"listenerList,omitempty"`
	Marker       *string                  `json:"marker,omitempty"`
	IsTruncated  *bool                    `json:"isTruncated,omitempty"`
	NextMarker   *string                  `json:"nextMarker,omitempty"`
	MaxKeys      *int32                   `json:"maxKeys,omitempty"`
}
