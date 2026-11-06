package rds

type RecoveryToSourceInstanceTableModel struct {
	TableName    *string `json:"tableName,omitempty"`
	NewTablename *string `json:"newTablename,omitempty"`
}
