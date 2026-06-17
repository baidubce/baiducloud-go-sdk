package bls

type Hit struct {
	Index   *string                   `json:"_index,omitempty"`
	Id      *string                   `json:"_id,omitempty"`
	Score   *float64                  `json:"_score,omitempty"`
	Sort    []*string                 `json:"sort,omitempty"`
	Version *int32                    `json:"_version,omitempty"`
	Fields  *map[string][]interface{} `json:"fields,omitempty"`
}
