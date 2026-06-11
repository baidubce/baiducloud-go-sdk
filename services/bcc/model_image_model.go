package bcc

type ImageModel struct {
	Id             *string          `json:"id,omitempty"`
	Name           *string          `json:"name,omitempty"`
	BccType        *string          `json:"type,omitempty"`
	OsType         *string          `json:"osType,omitempty"`
	OsVersion      *string          `json:"osVersion,omitempty"`
	OsArch         *string          `json:"osArch,omitempty"`
	OsName         *string          `json:"osName,omitempty"`
	OsBuild        *string          `json:"osBuild,omitempty"`
	OsLang         *string          `json:"osLang,omitempty"`
	SpecialVersion *string          `json:"specialVersion,omitempty"`
	CreateTime     *string          `json:"createTime,omitempty"`
	Status         *string          `json:"status,omitempty"`
	Encrypted      *bool            `json:"encrypted,omitempty"`
	BccPackage     *bool            `json:"package,omitempty"`
	Desc           *string          `json:"desc,omitempty"`
	DiskSize       *int32           `json:"diskSize,omitempty"`
	MinDiskGb      *int32           `json:"minDiskGb,omitempty"`
	EphemeralSize  *int32           `json:"ephemeralSize,omitempty"`
	Snapshots      []*SnapshotModel `json:"snapshots,omitempty"`
}
