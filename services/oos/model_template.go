package oos

type Template struct {
	Id                     *string         `json:"id,omitempty"`
	Ref                    *string         `json:"ref,omitempty"`
	Name                   *string         `json:"name,omitempty"`
	OosType                *string         `json:"type,omitempty"`
	Description            *string         `json:"description,omitempty"`
	Tags                   []*KeyValuePair `json:"tags,omitempty"`
	Linear                 *bool           `json:"linear,omitempty"`
	Parallelism            *int32          `json:"parallelism,omitempty"`
	Operators              []*Operator     `json:"operators,omitempty"`
	Links                  []*LinkModel    `json:"links,omitempty"`
	Properties             []*Property     `json:"properties,omitempty"`
	UpdatedTime            *string         `json:"updatedTime,omitempty"`
	SupportedInstanceTypes []*string       `json:"supportedInstanceTypes,omitempty"`
}
