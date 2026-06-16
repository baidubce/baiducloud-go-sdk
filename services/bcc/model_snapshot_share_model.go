package bcc

type SnapshotShareModel struct {
	SourceSnapshotId   *string `json:"sourceSnapshotId,omitempty"`
	SourceSnapshotUuid *string `json:"sourceSnapshotUuid,omitempty"`
	SnapshotId         *string `json:"snapshotId,omitempty"`
	SourceAccountId    *string `json:"sourceAccountId,omitempty"`
	AccountId          *string `json:"accountId,omitempty"`
	SnapshotType       *string `json:"snapshotType,omitempty"`
	Name               *string `json:"name,omitempty"`
	SizeInGB           *int32  `json:"sizeInGB,omitempty"`
	ShareTime          *string `json:"shareTime,omitempty"`
	Desc               *string `json:"desc,omitempty"`
	ShareStatus        *string `json:"shareStatus,omitempty"`
	EncryptKey         *string `json:"encryptKey,omitempty"`
	IsSourceDeleted    *bool   `json:"isSourceDeleted,omitempty"`
}
