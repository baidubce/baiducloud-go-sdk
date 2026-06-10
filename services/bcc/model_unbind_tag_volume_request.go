package bcc

type UnbindTagVolumeRequest struct {
	VolumeId   *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
