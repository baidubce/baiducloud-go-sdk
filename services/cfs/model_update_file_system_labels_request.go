package cfs

type UpdateFileSystemLabelsRequest struct {
	Tag  *string   `json:"-"`
	FsId []*string `json:"fsId,omitempty"`
	Tags []*Tag    `json:"tags,omitempty"`
}
