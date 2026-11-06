package rds

type OpenApiBackup struct {
	BackupId        *string `json:"backupId,omitempty"`
	BackupSize      *int64  `json:"backupSize,omitempty"`
	BackupType      *string `json:"backupType,omitempty"`
	BackupStatus    *string `json:"backupStatus,omitempty"`
	BackupStartTime *string `json:"backupStartTime,omitempty"`
	BackupEndTime   *string `json:"backupEndTime,omitempty"`
	DownloadUrl     *string `json:"downloadUrl,omitempty"`
	DownloadExpires *string `json:"downloadExpires,omitempty"`
}
