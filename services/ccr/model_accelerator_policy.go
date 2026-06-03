package ccr

type AcceleratorPolicy struct {
	CreationTime *string              `json:"creationTime,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Enabled      *bool                `json:"enabled,omitempty"`
	Filters      []*AcceleratorFilter `json:"filters,omitempty"`
	Id           *int32               `json:"id,omitempty"`
	Name         *string              `json:"name,omitempty"`
	UpdateTime   *string              `json:"updateTime,omitempty"`
}
