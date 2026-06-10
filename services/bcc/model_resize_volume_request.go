package bcc

type ResizeVolumeRequest struct {
	VolumeId       *string `json:"-"`
	NewCdsSizeInGB *int32  `json:"newCdsSizeInGB,omitempty"`
	NewExtraIO     *int32  `json:"newExtraIO,omitempty"`
	NewVolumeType  *string `json:"newVolumeType,omitempty"`
}
