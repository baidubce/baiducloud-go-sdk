package bls

type DataStreams struct {
	Name           *string   `json:"name,omitempty"`
	BackingIndices []*string `json:"backing_indices,omitempty"`
	TimestampField *string   `json:"timestamp_field,omitempty"`
}
