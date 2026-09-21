package dbsc

type TableIndexColumnItem struct {
	ColumnName  *string `json:"columnName,omitempty"`
	Sequence    *int32  `json:"sequence,omitempty"`
	Collation   *string `json:"collation,omitempty"`
	Cardinality *int64  `json:"cardinality,omitempty"`
	Nullable    *string `json:"nullable,omitempty"`
}
