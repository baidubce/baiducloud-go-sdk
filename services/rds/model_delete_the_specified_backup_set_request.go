package rds

type DeleteTheSpecifiedBackupSetRequest struct {
	InstanceId *string `json:"-"`
	SnapshotId *string `json:"-"`
}
