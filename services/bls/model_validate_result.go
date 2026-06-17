package bls

type ValidateResult struct {
	LogStore    *LogStore `json:"logStore,omitempty"`
	Valid       *bool     `json:"valid,omitempty"`
	Reason      *string   `json:"reason,omitempty"`
	Columns     []*string `json:"columns,omitempty"`
	ColumnTypes []*string `json:"columnTypes,omitempty"`
}
