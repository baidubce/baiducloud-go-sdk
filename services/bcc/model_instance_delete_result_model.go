package bcc

type InstanceDeleteResultModel struct {
	InstanceId  *string   `json:"instanceId,omitempty"`
	Eip         *string   `json:"eip,omitempty"`
	InsnapIds   []*string `json:"insnapIds,omitempty"`
	SnapshotIds []*string `json:"snapshotIds,omitempty"`
	VolumeIds   []*string `json:"volumeIds,omitempty"`
}
