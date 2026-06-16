package bcc

type AttachAspRequest struct {
	AspId              *string   `json:"-"`
	VolumeIds          []*string `json:"volumeIds,omitempty"`
	DeleteAutoSnapshot *bool     `json:"deleteAutoSnapshot,omitempty"`
}
