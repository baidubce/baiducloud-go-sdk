package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateNodesClusterExpansionV2Response struct {
	bce.BaseResponse
	CceInstanceIDs []*string `json:"cceInstanceIDs,omitempty"`
	RequestID      *string   `json:"requestID,omitempty"`
}
