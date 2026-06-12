package aihc

type ModifyDatasetRequest struct {
	DatasetId       *string            `json:"-"`
	Name            *string            `json:"name,omitempty"`
	Description     *string            `json:"description,omitempty"`
	VisibilityScope *string            `json:"visibilityScope,omitempty"`
	VisibilityUser  []*PermissionEntry `json:"visibilityUser,omitempty"`
	VisibilityGroup []*PermissionEntry `json:"visibilityGroup,omitempty"`
}
