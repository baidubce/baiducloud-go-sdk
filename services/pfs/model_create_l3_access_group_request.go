package pfs

type CreateL3AccessGroupRequest struct {
	Action                *string `json:"-"`
	AccessGroupName       *string `json:"accessGroupName,omitempty"`
	ZoneName              *string `json:"zoneName,omitempty"`
	Description           *string `json:"description,omitempty"`
	SourceZoneName        *string `json:"sourceZoneName,omitempty"`
	SourceAccessGroupName *string `json:"sourceAccessGroupName,omitempty"`
}
