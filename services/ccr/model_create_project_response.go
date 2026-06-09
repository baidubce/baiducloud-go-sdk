package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateProjectResponse struct {
	bce.BaseResponse
	AutoScan     *string `json:"autoScan,omitempty"`
	ChartCount   *int32  `json:"chartCount,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
	ProjectId    *int32  `json:"projectId,omitempty"`
	ProjectName  *string `json:"projectName,omitempty"`
	Public       *string `json:"public,omitempty"`
	RepoCount    *int32  `json:"repoCount,omitempty"`
	UpdateTime   *string `json:"updateTime,omitempty"`
}
