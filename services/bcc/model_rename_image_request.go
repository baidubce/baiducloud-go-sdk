package bcc

type RenameImageRequest struct {
	ImageIds []*string `json:"imageIds,omitempty"`
	Name     *string   `json:"name,omitempty"`
}
