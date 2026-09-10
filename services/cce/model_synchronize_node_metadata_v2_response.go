package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SynchronizeNodeMetadataV2Response struct {
	bce.BaseResponse
	ClusterID *string `json:"clusterID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
}
