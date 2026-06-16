package bcc

type ListSnapshotShareRequest struct {
	Marker  *string `json:"marker,omitempty"`
	MaxKeys *int32  `json:"maxKeys,omitempty"`
}
