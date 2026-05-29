package cfs

type UpdatePermissionGroupRequest struct {
	AgName      *string `json:"-"`
	Description *string `json:"description,omitempty"`
}
