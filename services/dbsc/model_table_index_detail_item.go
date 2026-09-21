package dbsc

type TableIndexDetailItem struct {
	SchemaName *string                 `json:"schemaName,omitempty"`
	TableName  *string                 `json:"tableName,omitempty"`
	IndexName  *string                 `json:"indexName,omitempty"`
	ColumnName *string                 `json:"columnName,omitempty"`
	NonUnique  *int32                  `json:"nonUnique,omitempty"`
	IndexType  *string                 `json:"indexType,omitempty"`
	Comment    *string                 `json:"comment,omitempty"`
	Columns    []*TableIndexColumnItem `json:"columns,omitempty"`
}
