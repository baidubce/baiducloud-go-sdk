package cfs

type UpdateFileSystemLabelsRequest struct {
	FsId []*string `json:"fsId,omitempty"`
	Tags []*Tag    `json:"tags,omitempty"`
}
