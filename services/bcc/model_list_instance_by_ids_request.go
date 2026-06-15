package bcc

type ListInstanceByIdsRequest struct {
	Marker      *string   `json:"-"`
	MaxKeys     *int32    `json:"-"`
	InstanceIds []*string `json:"instanceIds,omitempty"`
}
