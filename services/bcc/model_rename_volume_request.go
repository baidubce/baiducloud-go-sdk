package bcc

type RenameVolumeRequest struct {
	VolumeId *string `json:"-"`
	Name     *string `json:"name,omitempty"`
}
