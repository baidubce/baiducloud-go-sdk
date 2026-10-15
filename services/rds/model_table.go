package rds

type Table struct {
	TableName    *string `json:"tableName,omitempty"`
	NewTablename *string `json:"newTablename,omitempty"`
}
