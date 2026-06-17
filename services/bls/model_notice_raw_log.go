package bls

type NoticeRawLog struct {
	BlsType      *string       `json:"type,omitempty"`
	RefTarget    *int32        `json:"refTarget,omitempty"`
	CustomTarget *CustomTarget `json:"customTarget,omitempty"`
	Columns      []*string     `json:"columns,omitempty"`
	Limit        *int32        `json:"limit,omitempty"`
}
