package dbsc

type TableColumnItem struct {
	SchemaName   *string `json:"schemaName,omitempty"`
	TableName    *string `json:"tableName,omitempty"`
	ColumnName   *string `json:"columnName,omitempty"`
	Position     *int64  `json:"position,omitempty"`
	DbscType     *string `json:"type,omitempty"`
	Charset      *string `json:"charset,omitempty"`
	Collation    *string `json:"collation,omitempty"`
	Length       *int64  `json:"length,omitempty"`
	Nullable     *string `json:"nullable,omitempty"`
	Comment      *string `json:"comment,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
}
