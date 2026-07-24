package bcc

type DetachAspRequest struct {
	AspId     *string   `json:"-"`
	VolumeIds []*string `json:"volumeIds,omitempty"`
}
