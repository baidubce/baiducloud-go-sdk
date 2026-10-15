package rds

type RecoveryToSourceInstanceModel struct {
	RestoreMode *string  `json:"restoreMode,omitempty"`
	DbName      *string  `json:"dbName,omitempty"`
	NewDbname   *string  `json:"newDbname,omitempty"`
	Tables      []*Table `json:"tables,omitempty"`
}
