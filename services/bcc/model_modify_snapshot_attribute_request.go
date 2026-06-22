package bcc

type ModifySnapshotAttributeRequest struct {
	SnapshotId      *string `json:"-"`
	SnapshotName    *string `json:"snapshotName,omitempty"`
	RetentionInDays *int32  `json:"retentionInDays,omitempty"`
	Desc            *string `json:"desc,omitempty"`
}
