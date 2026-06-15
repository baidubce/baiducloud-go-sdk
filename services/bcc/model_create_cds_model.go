package bcc

type CreateCdsModel struct {
	StorageType            *string `json:"storageType,omitempty"`
	CdsSizeInGB            *int32  `json:"cdsSizeInGB,omitempty"`
	CdsNum                 *int32  `json:"cdsNum,omitempty"`
	CdsExtraIo             *int32  `json:"cdsExtraIo,omitempty"`
	SnapshotId             *string `json:"snapshotId,omitempty"`
	EncryptKey             *string `json:"encryptKey,omitempty"`
	EnableDeleteProtection *string `json:"enableDeleteProtection,omitempty"`
	DeleteWithInstance     *bool   `json:"deleteWithInstance,omitempty"`
	DeleteAutoSnapshot     *bool   `json:"deleteAutoSnapshot,omitempty"`
	Name                   *string `json:"name,omitempty"`
}
