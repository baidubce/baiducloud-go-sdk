package apm

type SampleConfig struct {
	Enabled    *bool              `json:"enabled,omitempty"`
	Processors []*SampleProcessor `json:"processors,omitempty"`
}
