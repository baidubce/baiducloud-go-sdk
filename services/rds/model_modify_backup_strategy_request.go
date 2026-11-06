package rds

type ModifyBackupStrategyRequest struct {
	InstanceId                    *string   `json:"-"`
	BackupDays                    *string   `json:"backupDays,omitempty"`
	BackupTime                    *string   `json:"backupTime,omitempty"`
	Persistent                    *bool     `json:"persistent,omitempty"`
	ExpireInDays                  *int32    `json:"expireInDays,omitempty"`
	LogBackupRetainDays           *int32    `json:"logBackupRetainDays,omitempty"`
	DataBackupType                *string   `json:"dataBackupType,omitempty"`
	IncrementalDataBackupEnable   *bool     `json:"incrementalDataBackupEnable,omitempty"`
	IncrementalDataBackupInterval *int32    `json:"incrementalDataBackupInterval,omitempty"`
	LatestDataBackupRetainDays    *int32    `json:"latestDataBackupRetainDays,omitempty"`
	RetentionPeriod               *int32    `json:"retentionPeriod,omitempty"`
	TimeUnit                      *string   `json:"timeUnit,omitempty"`
	DataBackupCopyEnable          *bool     `json:"dataBackupCopyEnable,omitempty"`
	DataBackupCopyStoragesRegions []*string `json:"dataBackupCopyStoragesRegions,omitempty"`
	DataBackupCopyRetainDays      *int32    `json:"dataBackupCopyRetainDays,omitempty"`
}
