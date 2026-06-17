package bls

type RawLog struct {
	Project      *string                   `json:"project,omitempty"`
	LogStoreName *string                   `json:"logStoreName,omitempty"`
	Query        *string                   `json:"query,omitempty"`
	Columns      []*string                 `json:"columns,omitempty"`
	Limit        *int32                    `json:"limit,omitempty"`
	Logs         []*map[string]interface{} `json:"logs,omitempty"`
}
