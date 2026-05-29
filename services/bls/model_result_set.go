package bls

type ResultSet struct {
	QueryType   *string         `json:"queryType,omitempty"`
	Columns     []*string       `json:"columns,omitempty"`
	ColumnTypes []*string       `json:"columnTypes,omitempty"`
	Rows        [][]interface{} `json:"rows,omitempty"`
}
