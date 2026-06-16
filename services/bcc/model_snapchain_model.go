package bcc

type SnapchainModel struct {
	Status          *string `json:"status,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty"`
	UserId          *string `json:"userId,omitempty"`
	ChainSize       *string `json:"chainSize,omitempty"`
	ChainId         *string `json:"chainId,omitempty"`
	VolumeId        *string `json:"volumeId,omitempty"`
	ManualSnapCount *int32  `json:"manualSnapCount,omitempty"`
	AutoSnapCount   *int32  `json:"autoSnapCount,omitempty"`
	VolumeSize      *int32  `json:"volumeSize,omitempty"`
	CreateTime      *string `json:"createTime,omitempty"`
}
