package rds

type QueryBackupDetailsRequest struct {
	InstanceId *string `json:"-"`
	BackupId   *string `json:"-"`
}
