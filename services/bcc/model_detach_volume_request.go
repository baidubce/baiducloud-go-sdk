package bcc

type DetachVolumeRequest struct {
	VolumeId   *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
}
