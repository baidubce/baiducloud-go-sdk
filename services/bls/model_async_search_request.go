package bls

type AsyncSearchRequest struct {
	Name        *string                       `json:"-"`
	Query       *interface{}                  `json:"query,omitempty"`
	Aggs        *interface{}                  `json:"aggs,omitempty"`
	Fields      []*string                     `json:"fields,omitempty"`
	Sort        *map[string]map[string]string `json:"sort,omitempty"`
	SearchAfter []*string                     `json:"searchAfter,omitempty"`
	Highlight   *Highlight                    `json:"highlight,omitempty"`
	Size        *int32                        `json:"size,omitempty"`
}
