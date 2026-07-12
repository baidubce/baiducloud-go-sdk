package oos

type CreateExecutionV2Request struct {
	Locale      *string      `json:"-"`
	Description *string      `json:"description,omitempty"`
	Template    *Template    `json:"template,omitempty"`
	Parallelism *int32       `json:"parallelism,omitempty"`
	Manually    *bool        `json:"manually,omitempty"`
	Properties  *interface{} `json:"properties,omitempty"`
	Tags        []*Tag       `json:"tags,omitempty"`
}
