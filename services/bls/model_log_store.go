package bls

type LogStore struct {
	Project      *string   `json:"project,omitempty"`
	LogStoreName *string   `json:"logStoreName,omitempty"`
	LogStoreId   *string   `json:"logStoreId,omitempty"`
	Region       *string   `json:"region,omitempty"`
	Name         *string   `json:"name,omitempty"`
	Extends      []*Extend `json:"extends,omitempty"`
}
