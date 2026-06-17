package bls

type UpdateLogStoreViewRequest struct {
	Project   *string     `json:"project,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Logstores []*LogStore `json:"logstores,omitempty"`
}
