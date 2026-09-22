package scs

type Record struct {
	BackupRecordId *string `json:"backupRecordId,omitempty"`
	BackupId       *string `json:"backupId,omitempty"`
	StartTime      *string `json:"startTime,omitempty"`
	Duration       *int64  `json:"duration,omitempty"`
	ObjectSize     *int64  `json:"objectSize,omitempty"`
	BackupType     *string `json:"backupType,omitempty"`
	BackupStatus   *string `json:"backupStatus,omitempty"`
	ShardName      *string `json:"shardName,omitempty"`
	Comment        *string `json:"comment,omitempty"`
}
