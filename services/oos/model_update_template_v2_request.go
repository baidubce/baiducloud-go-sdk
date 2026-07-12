package oos

type UpdateTemplateV2Request struct {
	Namespace   *string         `json:"namespace,omitempty"`
	Id          *string         `json:"id,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Tags        []*KeyValuePair `json:"tags,omitempty"`
	Linear      *bool           `json:"linear,omitempty"`
	Parallelism *int32          `json:"parallelism,omitempty"`
	Operators   []*Operator     `json:"operators,omitempty"`
	Links       []*LinkModel    `json:"links,omitempty"`
	Properties  []*Property     `json:"properties,omitempty"`
}
