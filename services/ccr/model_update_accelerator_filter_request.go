package ccr

type UpdateAcceleratorFilterRequest struct {
	InstanceId  *string              `json:"-"`
	PolicyId    *string              `json:"-"`
	Description *string              `json:"description,omitempty"`
	Filters     []*AcceleratorFilter `json:"filters,omitempty"`
	Name        *string              `json:"name,omitempty"`
}
