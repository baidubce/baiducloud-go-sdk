package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateVolumeClusterResponse struct {
	bce.BaseResponse
	ClusterIds []*string `json:"clusterIds,omitempty"`
}
