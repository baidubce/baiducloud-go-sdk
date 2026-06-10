package bcc

type VolumeAttachmentModel struct {
	VolumeId   *string `json:"volumeId,omitempty"`
	Serial     *string `json:"serial,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	Device     *string `json:"device,omitempty"`
}
