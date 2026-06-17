package bls

type Field struct {
	BlsType        *string `json:"type,omitempty"`
	CaseSensitive  *bool   `json:"caseSensitive,omitempty"`
	IncludeChinese *bool   `json:"includeChinese,omitempty"`
	Separators     *string `json:"separators,omitempty"`
	DynamicMapping *bool   `json:"dynamicMapping,omitempty"`
	Searchable     *bool   `json:"searchable,omitempty"`
	Aggregatable   *bool   `json:"aggregatable,omitempty"`
	MetadataField  *bool   `json:"metadata_field,omitempty"`
}
