package cfs

type CreateFileSystemRequest struct {
	FsName        *string `json:"fsName,omitempty"`
	Zone          *string `json:"zone,omitempty"`
	CfsType       *string `json:"type,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	Tags          []*Tag  `json:"tags,omitempty"`
	CapacityQuota *int64  `json:"capacityQuota,omitempty"`
}
