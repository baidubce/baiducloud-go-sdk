package ccr

type ReplicationFilterRequest struct {
	CcrType *string `json:"type,omitempty"`
	Value   *string `json:"value,omitempty"`
}
