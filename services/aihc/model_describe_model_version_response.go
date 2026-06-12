package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeModelVersionResponse struct {
	bce.BaseResponse
	Name            *string            `json:"name,omitempty"`
	Id              *string            `json:"id,omitempty"`
	InitSource      *string            `json:"initSource,omitempty"`
	ModelFormat     *string            `json:"modelFormat,omitempty"`
	Description     *string            `json:"description,omitempty"`
	CreatedAt       *string            `json:"createdAt,omitempty"`
	UpdatedAt       *string            `json:"updatedAt,omitempty"`
	Owner           *string            `json:"owner,omitempty"`
	OwnerName       *string            `json:"ownerName,omitempty"`
	VisibilityScope *string            `json:"visibilityScope,omitempty"`
	VersionEntry    *ModelVersionEntry `json:"versionEntry,omitempty"`
}
