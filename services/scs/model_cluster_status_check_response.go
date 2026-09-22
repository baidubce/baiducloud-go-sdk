package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ClusterStatusCheckResponse struct {
	bce.BaseResponse
	ClusterStatus *string    `json:"clusterStatus,omitempty"`
	CheckList     *CheckList `json:"checkList,omitempty"`
}
