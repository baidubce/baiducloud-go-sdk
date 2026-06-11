package bcc

type UnbindTagImageRequest struct {
	ImageId    *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
