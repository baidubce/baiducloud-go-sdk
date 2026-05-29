package cfs

type CreatePermissionGroupRequest struct {
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	Description     *string `json:"description,omitempty"`
}
