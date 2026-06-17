package bls

type CreateIndexRequest struct {
	LogStoreName   *string                 `json:"-"`
	Project        *string                 `json:"-"`
	Fulltext       *bool                   `json:"fulltext,omitempty"`
	CaseSensitive  *bool                   `json:"caseSensitive,omitempty"`
	IncludeChinese *bool                   `json:"includeChinese,omitempty"`
	Separators     *string                 `json:"separators,omitempty"`
	Fields         *map[string]*IndexField `json:"fields,omitempty"`
}
