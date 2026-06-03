package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRepositoryResponse struct {
	bce.BaseResponse
	ProjectName           *string `json:"projectName,omitempty"`
	RepositoryName        *string `json:"repositoryName,omitempty"`
	Description           *string `json:"description,omitempty"`
	RepositoryPath        *string `json:"repositoryPath,omitempty"`
	PrivateRepositoryPath *string `json:"privateRepositoryPath,omitempty"`
	TagCount              *int32  `json:"tagCount,omitempty"`
	PullCount             *int32  `json:"pullCount,omitempty"`
	CreationTime          *string `json:"creationTime,omitempty"`
	UpdateTime            *string `json:"updateTime,omitempty"`
}
