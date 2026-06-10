package bcc

type CreateVolumeRequest struct {
	ZoneName               *string                  `json:"zoneName,omitempty"`
	StorageType            *string                  `json:"storageType,omitempty"`
	CdsSizeInGB            *int32                   `json:"cdsSizeInGB,omitempty"`
	CdsExtraIo             *int32                   `json:"cdsExtraIo,omitempty"`
	SnapshotId             *string                  `json:"snapshotId,omitempty"`
	ShareSnapshotId        *string                  `json:"shareSnapshotId,omitempty"`
	EnableDeleteProtection *string                  `json:"enableDeleteProtection,omitempty"`
	InstanceId             *string                  `json:"instanceId,omitempty"`
	EncryptKey             *string                  `json:"encryptKey,omitempty"`
	Name                   *string                  `json:"name,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	RenewTimeUnit          *string                  `json:"renewTimeUnit,omitempty"`
	RenewTime              *int32                   `json:"renewTime,omitempty"`
	RelationTag            *bool                    `json:"relationTag,omitempty"`
	Tags                   []*TagModel              `json:"tags,omitempty"`
	ResGroupId             *string                  `json:"resGroupId,omitempty"`
	Billing                *Billing                 `json:"billing,omitempty"`
	ClusterId              *string                  `json:"clusterId,omitempty"`
	ChargeType             *string                  `json:"chargeType,omitempty"`
	AutoSnapshotPolicy     *AutoSnapshotPolicyModel `json:"autoSnapshotPolicy,omitempty"`
	DeleteWithInstance     *bool                    `json:"deleteWithInstance,omitempty"`
	DeleteAutoSnapshot     *bool                    `json:"deleteAutoSnapshot,omitempty"`
	PurchaseCount          *int32                   `json:"purchaseCount,omitempty"`
}
