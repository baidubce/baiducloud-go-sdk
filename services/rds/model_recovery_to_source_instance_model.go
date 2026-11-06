package rds

type RecoveryToSourceInstanceModel struct {
	DbName      *string  `json:"dbName,omitempty"`
	NewDbname   *string  `json:"newDbname,omitempty"`
	RestoreMode *string  `json:"restoreMode,omitempty"`
	Tables      []*Table `json:"tables,omitempty"`
}
