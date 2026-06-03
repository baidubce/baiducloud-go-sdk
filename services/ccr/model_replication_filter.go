package ccr

type ReplicationFilter struct {
	CcrType *string `json:"type,omitempty"`
	Value   *string `json:"value,omitempty"`
}
