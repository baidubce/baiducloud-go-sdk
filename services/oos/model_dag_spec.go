package oos

type DagSpec struct {
	Ref              *string                `json:"ref,omitempty"`
	Namespace        *string                `json:"namespace,omitempty"`
	Name             *string                `json:"name,omitempty"`
	Names            []*string              `json:"names,omitempty"`
	Description      *string                `json:"description,omitempty"`
	Tags             *interface{}           `json:"tags,omitempty"`
	Operators        []*TaskOperatorSummary `json:"operators,omitempty"`
	Linear           *bool                  `json:"linear,omitempty"`
	Links            []*LinkModel           `json:"links,omitempty"`
	Inputs           []*InputModel          `json:"inputs,omitempty"`
	Outputs          []*OutputModel         `json:"outputs,omitempty"`
	Parallelism      *int32                 `json:"parallelism,omitempty"`
	Extra            *interface{}           `json:"extra,omitempty"`
	CreatedTimestamp *int64                 `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp *int64                 `json:"updatedTimestamp,omitempty"`
}
