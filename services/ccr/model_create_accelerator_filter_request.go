package ccr

type CreateAcceleratorFilterRequest struct {
	InstanceId  *string              `json:"-"`
	Description *string              `json:"description,omitempty"`
	Filters     []*AcceleratorFilter `json:"filters,omitempty"`
	Name        *string              `json:"name,omitempty"`
}
