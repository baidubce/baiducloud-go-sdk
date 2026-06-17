package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type EhcClusterListResponse struct {
	bce.BaseResponse
	TotalCount  *int32        `json:"totalCount,omitempty"`
	EhcClusters []*EhcCluster `json:"ehcClusters,omitempty"`
}
