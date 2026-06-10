package bcc

type VolumeMultiAttachInfo struct {
	VolumeId   *string `json:"volumeId,omitempty"`
	Serial     *string `json:"serial,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	Status     *string `json:"status,omitempty"`
}
