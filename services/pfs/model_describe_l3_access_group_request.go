package pfs

type DescribeL3AccessGroupRequest struct {
	Action          *string `json:"-"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	ZoneName        *string `json:"zoneName,omitempty"`
}
