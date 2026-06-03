package ccr

type TriggerFilter struct {
	CcrType *string `json:"type,omitempty"`
	Value   *string `json:"value,omitempty"`
}
