package apm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeTopologyResponse struct {
	bce.BaseResponse
	Success *bool       `json:"success,omitempty"`
	Code    *string     `json:"code,omitempty"`
	Message *string     `json:"message,omitempty"`
	Nodes   []*TopoNode `json:"nodes,omitempty"`
	Edges   []*TopoEdge `json:"edges,omitempty"`
}
