package bcc

type UnbindTagVolumeClusterRequest struct {
	ClusterId  *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
