package bcc

type BindTagVolumeRequest struct {
	VolumeId   *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
