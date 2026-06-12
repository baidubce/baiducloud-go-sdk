package aihc

type CreateDatasetRequest struct {
	Name             *string              `json:"name,omitempty"`
	StorageType      *string              `json:"storageType,omitempty"`
	StorageInstance  *string              `json:"storageInstance,omitempty"`
	ImportFormat     *string              `json:"importFormat,omitempty"`
	Description      *string              `json:"description,omitempty"`
	Owner            *string              `json:"owner,omitempty"`
	VisibilityScope  *string              `json:"visibilityScope,omitempty"`
	VisibilityUser   []*PermissionEntry   `json:"visibilityUser,omitempty"`
	VisibilityGroup  []*PermissionEntry   `json:"visibilityGroup,omitempty"`
	InitVersionEntry *DatasetVersionEntry `json:"initVersionEntry,omitempty"`
}
