package scs

type BatchBackupRecord struct {
	BatchId     *string   `json:"batchId,omitempty"`
	BackupType  *string   `json:"backupType,omitempty"`
	Comment     *string   `json:"comment,omitempty"`
	StartTime   *string   `json:"startTime,omitempty"`
	Recoverable *string   `json:"recoverable,omitempty"`
	Records     []*Record `json:"records,omitempty"`
}
