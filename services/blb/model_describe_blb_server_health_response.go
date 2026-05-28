package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbServerHealthResponse struct {
	bce.BaseResponse
	BackendServerList []*BackendServerStatus `json:"backendServerList,omitempty"`
	BlbType           *string                `json:"type,omitempty"`
	ListenerPort      *int32                 `json:"listenerPort,omitempty"`
	BackendPort       *int32                 `json:"backendPort,omitempty"`
	Marker            *string                `json:"marker,omitempty"`
	IsTruncated       *bool                  `json:"isTruncated,omitempty"`
	NextMarker        *string                `json:"nextMarker,omitempty"`
	MaxKeys           *int32                 `json:"maxKeys,omitempty"`
}
