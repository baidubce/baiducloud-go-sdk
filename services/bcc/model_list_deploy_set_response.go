package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListDeploySetResponse struct {
	bce.BaseResponse
	PageNo     *int32            `json:"pageNo,omitempty"`
	PageSize   *int32            `json:"pageSize,omitempty"`
	TotalCount *int32            `json:"totalCount,omitempty"`
	DeploySets []*DeploySetModel `json:"deploySets,omitempty"`
}
