package bcc

type VolumeModel struct {
	Id                      *string                   `json:"id,omitempty"`
	DiskCategory            *string                   `json:"diskCategory,omitempty"`
	ProductCategory         *string                   `json:"productCategory,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	DiskSizeInGB            *int32                    `json:"diskSizeInGB,omitempty"`
	CdsExtraIo              *int32                    `json:"cdsExtraIo,omitempty"`
	FailureStatus           *string                   `json:"failureStatus,omitempty"`
	CreateTime              *string                   `json:"createTime,omitempty"`
	ExpireTime              *string                   `json:"expireTime,omitempty"`
	Status                  *string                   `json:"status,omitempty"`
	ShareSnapshotId         *string                   `json:"shareSnapshotId,omitempty"`
	EnableDeleteProtection  *bool                     `json:"enableDeleteProtection,omitempty"`
	EbcDiskSize             *int32                    `json:"ebcDiskSize,omitempty"`
	EnableAutoRenew         *bool                     `json:"enableAutoRenew,omitempty"`
	AutoRenewTime           *int32                    `json:"autoRenewTime,omitempty"`
	Tags                    []*TagModel               `json:"tags,omitempty"`
	BccType                 *string                   `json:"type,omitempty"`
	StorageType             *string                   `json:"storageType,omitempty"`
	IsSystemVolume          *bool                     `json:"isSystemVolume,omitempty"`
	Description             *string                   `json:"description,omitempty"`
	Desc                    *string                   `json:"desc,omitempty"`
	PaymentTiming           *string                   `json:"paymentTiming,omitempty"`
	ZoneName                *string                   `json:"zoneName,omitempty"`
	RegionId                *string                   `json:"regionId,omitempty"`
	SourceSnapshotId        *string                   `json:"sourceSnapshotId,omitempty"`
	SnapshotNum             *string                   `json:"snapshotNum,omitempty"`
	ClusterId               *string                   `json:"clusterId,omitempty"`
	ResGroupInfos           []*GroupInfo              `json:"resGroupInfos,omitempty"`
	AutoSnapshotPolicy      *AutoSnapshotPolicyModel  `json:"autoSnapshotPolicy,omitempty"`
	AutoSnapshotPolicyInfos []*AutoSnapshotPolicyInfo `json:"autoSnapshotPolicyInfos,omitempty"`
	EncryptKey              *string                   `json:"encryptKey,omitempty"`
	EncryptKeySpec          *string                   `json:"encryptKeySpec,omitempty"`
	Encrypted               *bool                     `json:"encrypted,omitempty"`
	DeleteWithInstance      *bool                     `json:"deleteWithInstance,omitempty"`
	DeleteAutoSnapshot      *bool                     `json:"deleteAutoSnapshot,omitempty"`
	Attachments             []*VolumeAttachmentModel  `json:"attachments,omitempty"`
	MultiAttachInfos        []*VolumeMultiAttachInfo  `json:"multiAttachInfos,omitempty"`
	MultiAttach             *bool                     `json:"multiAttach,omitempty"`
	VolumeId                *string                   `json:"volumeId,omitempty"`
}
