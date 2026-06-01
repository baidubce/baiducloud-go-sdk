package apm

type SampleProcessor struct {
	Name       *string         `json:"name,omitempty"`
	Enabled    *bool           `json:"enabled,omitempty"`
	Filters    []*SampleFilter `json:"filters,omitempty"`
	SampleRate *float32        `json:"sampleRate,omitempty"`
}
