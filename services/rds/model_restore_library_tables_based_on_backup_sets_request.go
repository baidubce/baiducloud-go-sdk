package rds

type RestoreLibraryTablesBasedOnBackupSetsRequest struct {
	InstanceId       *string                               `json:"-"`
	SourceInstanceId *string                               `json:"sourceInstanceId,omitempty"`
	TargetInstanceId *string                               `json:"targetInstanceId,omitempty"`
	SnapshotId       *string                               `json:"snapshotId,omitempty"`
	Data             []*RecoveryToSourceInstanceModel      `json:"data,omitempty"`
	DbName           *string                               `json:"dbName,omitempty"`
	NewDbname        *string                               `json:"newDbname,omitempty"`
	RestoreMode      *string                               `json:"restoreMode,omitempty"`
	Tables           []*RecoveryToSourceInstanceTableModel `json:"tables,omitempty"`
	TableName        *string                               `json:"tableName,omitempty"`
	NewTablename     *string                               `json:"newTablename,omitempty"`
}
