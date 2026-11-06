package rds

type RestoreTheDatabaseTableByTimePointRequest struct {
	InstanceId       *string                               `json:"-"`
	SourceInstanceId *string                               `json:"sourceInstanceId,omitempty"`
	TargetInstanceId *string                               `json:"targetInstanceId,omitempty"`
	Datetime         *string                               `json:"datetime,omitempty"`
	Data             []*RecoveryToSourceInstanceModel      `json:"data,omitempty"`
	DbName           *string                               `json:"dbName,omitempty"`
	NewDbname        *string                               `json:"newDbname,omitempty"`
	RestoreMode      *string                               `json:"restoreMode,omitempty"`
	Tables           []*RecoveryToSourceInstanceTableModel `json:"tables,omitempty"`
	TableName        *string                               `json:"tableName,omitempty"`
	NewTablename     *string                               `json:"newTablename,omitempty"`
}
