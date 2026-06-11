package bcc

type RemoteCopyImageRequest struct {
	ImageId    *string   `json:"-"`
	Name       *string   `json:"name,omitempty"`
	DestRegion []*string `json:"destRegion,omitempty"`
}
