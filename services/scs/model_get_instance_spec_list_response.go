package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceSpecListResponse struct {
	bce.BaseResponse
	DefaultNodeTypeList     []*NodeTypeItem `json:"defaultNodeTypeList,omitempty"`
	ClusterNodeTypeList     []*NodeTypeItem `json:"clusterNodeTypeList,omitempty"`
	PegaClusterNodeTypeList []*NodeTypeItem `json:"pegaClusterNodeTypeList,omitempty"`
}
