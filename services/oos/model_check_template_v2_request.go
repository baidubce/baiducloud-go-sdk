package oos

type CheckTemplateV2Request struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Tags        []*KeyValuePair `json:"tags,omitempty"`
	Linear      *bool           `json:"linear,omitempty"`
	Parallelism *int32          `json:"parallelism,omitempty"`
	Operators   []*Operator     `json:"operators,omitempty"`
	Links       []*LinkModel    `json:"links,omitempty"`
	Properties  []*Property     `json:"properties,omitempty"`
}
