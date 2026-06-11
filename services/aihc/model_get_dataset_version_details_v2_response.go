package aihc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDatasetVersionDetailsV2Response struct {
	bce.BaseResponse
	Id              *string              `json:"id,omitempty"`
	Name            *string              `json:"name,omitempty"`
	StorageType     *string              `json:"storageType,omitempty"`
	StorageInstance *string              `json:"storageInstance,omitempty"`
	ImportFormat    *string              `json:"importFormat,omitempty"`
	Description     *string              `json:"description,omitempty"`
	Owner           *string              `json:"owner,omitempty"`
	OwnerName       *string              `json:"ownerName,omitempty"`
	VisibilityScope *string              `json:"visibilityScope,omitempty"`
	VisibilityUser  []*PermissionEntry   `json:"visibilityUser,omitempty"`
	VisibilityGroup []*PermissionEntry   `json:"visibilityGroup,omitempty"`
	Permission      *string              `json:"permission,omitempty"`
	VersionEntry    *DatasetVersionEntry `json:"versionEntry,omitempty"`
	CreatedAt       *string              `json:"createdAt,omitempty"`
	UpdatedAt       *string              `json:"updatedAt,omitempty"`
}
