package apm

type SampleFilter struct {
	Key   *string `json:"key,omitempty"`
	Op    *string `json:"op,omitempty"`
	Value *string `json:"value,omitempty"`
}
