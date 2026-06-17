package bls

type IndexField struct {
	BlsType        *string `json:"type,omitempty"`
	CaseSensitive  *bool   `json:"caseSensitive,omitempty"`
	IncludeChinese *bool   `json:"includeChinese,omitempty"`
	Separators     *string `json:"separators,omitempty"`
	DynamicMapping *bool   `json:"dynamicMapping,omitempty"`
}
