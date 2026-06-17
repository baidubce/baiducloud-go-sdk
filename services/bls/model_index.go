package bls

type Index struct {
	Fulltext       *bool              `json:"fulltext,omitempty"`
	CaseSensitive  *bool              `json:"caseSensitive,omitempty"`
	IncludeChinese *bool              `json:"includeChinese,omitempty"`
	Separators     *string            `json:"separators,omitempty"`
	Fields         *map[string]*Field `json:"fields,omitempty"`
}
