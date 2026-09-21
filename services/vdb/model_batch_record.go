package vdb

type BatchRecord struct {
	BackupId     *int64  `json:"backupId,omitempty"`
	BackupStatus *string `json:"backupStatus,omitempty"`
	BackupType   *string `json:"backupType,omitempty"`
	Comment      *string `json:"comment,omitempty"`
	Duration     *int64  `json:"duration,omitempty"`
	NodeInfo     *string `json:"nodeInfo,omitempty"`
	ObjectSize   *int64  `json:"objectSize,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
}
