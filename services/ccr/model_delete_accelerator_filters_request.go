package ccr

type DeleteAcceleratorFiltersRequest struct {
	InstanceId *string  `json:"-"`
	Items      []*int32 `json:"items,omitempty"`
}
