package bcc

type RenameImageRequest struct {
	ImageId  *string   `json:"imageId,omitempty"`
	ImageIds []*string `json:"imageIds,omitempty"`
	Name     *string   `json:"name,omitempty"`
}
