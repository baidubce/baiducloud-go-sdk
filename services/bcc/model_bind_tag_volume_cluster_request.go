package bcc

type BindTagVolumeClusterRequest struct {
	ClusterId  *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
