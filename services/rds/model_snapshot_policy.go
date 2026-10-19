package rds

type SnapshotPolicy struct {
	BackupTime                    *string                     `json:"backupTime,omitempty"`
	BackupDays                    *string                     `json:"backupDays,omitempty"`
	Persistent                    *bool                       `json:"persistent,omitempty"`
	ExpireInDays                  *int32                      `json:"expireInDays,omitempty"`
	LogBackupRetainDays           *int32                      `json:"logBackupRetainDays,omitempty"`
	FreeSpaceInGB                 *int32                      `json:"freeSpaceInGB,omitempty"`
	DataBackupType                *string                     `json:"dataBackupType,omitempty"`
	FreeSpace                     *int32                      `json:"freeSpace,omitempty"`
	IncrementalDataBackupEnable   *bool                       `json:"incrementalDataBackupEnable,omitempty"`
	TimeUnit                      *string                     `json:"timeUnit,omitempty"`
	DataBackupCopyEnable          *bool                       `json:"dataBackupCopyEnable,omitempty"`
	IncrementalDataBackupInterval *int32                      `json:"incrementalDataBackupInterval,omitempty"`
	LatestDataBackupRetainDays    *int32                      `json:"latestDataBackupRetainDays,omitempty"`
	RetentionPeriod               *int32                      `json:"retentionPeriod,omitempty"`
	DataBackupRetainStrategys     []*DataBackupRetainStrategy `json:"dataBackupRetainStrategys,omitempty"`
	EncryptStrategy               *EncryptStrategy            `json:"encryptStrategy,omitempty"`
}
