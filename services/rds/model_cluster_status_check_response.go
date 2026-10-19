package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ClusterStatusCheckResponse struct {
	bce.BaseResponse
	ClusterStatus *string              `json:"clusterStatus,omitempty"`
	CheckList     *CheckStatusResponse `json:"checkList,omitempty"`
}
