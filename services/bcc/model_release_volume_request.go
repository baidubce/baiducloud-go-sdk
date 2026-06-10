package bcc

type ReleaseVolumeRequest struct {
	VolumeId           *string `json:"-"`
	AutoSnapshot       *string `json:"autoSnapshot,omitempty"`
	ManualSnapshot     *string `json:"manualSnapshot,omitempty"`
	CdsAttributeActive *bool   `json:"cdsAttributeActive,omitempty"`
	Recycle            *string `json:"recycle,omitempty"`
}
