package bcc

type AttachVolumeRequest struct {
	VolumeId           *string `json:"-"`
	InstanceId         *string `json:"instanceId,omitempty"`
	DeleteWithInstance *bool   `json:"deleteWithInstance,omitempty"`
	DeleteAutoSnapshot *bool   `json:"deleteAutoSnapshot,omitempty"`
}
