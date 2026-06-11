package bcc

type SnapshotModel struct {
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	SizeInGB     *int32  `json:"sizeInGB,omitempty"`
	CreateTime   *string `json:"createTime,omitempty"`
	Status       *string `json:"status,omitempty"`
	CreateMethod *string `json:"createMethod,omitempty"`
	VolumeId     *string `json:"volumeId,omitempty"`
	Desc         *string `json:"desc,omitempty"`
	ExpireTime   *string `json:"expireTime,omitempty"`
	InsnapId     *string `json:"insnapId,omitempty"`
	BccPackage   *bool   `json:"package,omitempty"`
	TemplateId   *string `json:"templateId,omitempty"`
	Encrypted    *bool   `json:"encrypted,omitempty"`
	Progress     *string `json:"progress,omitempty"`
	Tags         []*Tag  `json:"tags,omitempty"`
}
