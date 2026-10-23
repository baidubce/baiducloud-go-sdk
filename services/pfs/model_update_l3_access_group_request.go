package pfs

type UpdateL3AccessGroupRequest struct {
	Action          *string `json:"-"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
	Description     *string `json:"description,omitempty"`
}
