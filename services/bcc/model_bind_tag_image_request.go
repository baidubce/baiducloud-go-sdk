package bcc

type BindTagImageRequest struct {
	ImageId    *string     `json:"-"`
	ChangeTags []*TagModel `json:"changeTags,omitempty"`
}
