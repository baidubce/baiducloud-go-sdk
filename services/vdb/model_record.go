package vdb

type Record struct {
	BackupType     *string        `json:"backupType,omitempty"`
	BatchId        *string        `json:"batchId,omitempty"`
	BatchRecords   []*BatchRecord `json:"batchRecords,omitempty"`
	Comment        *string        `json:"comment,omitempty"`
	EndTime        *string        `json:"endTime,omitempty"`
	Recoverable    *string        `json:"recoverable,omitempty"`
	StartTime      *string        `json:"startTime,omitempty"`
	Status         *string        `json:"status,omitempty"`
	StorageType    *string        `json:"storageType,omitempty"`
	TotalSizeBytes *int64         `json:"totalSizeBytes,omitempty"`
}
