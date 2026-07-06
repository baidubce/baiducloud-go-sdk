package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeJobNodesResponse struct {
	bce.BaseResponse
	RequestId *string   `json:"requestId,omitempty"`
	JobId     *string   `json:"jobId,omitempty"`
	Nodes     []*string `json:"nodes,omitempty"`
}
