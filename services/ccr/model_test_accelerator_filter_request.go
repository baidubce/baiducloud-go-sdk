package ccr

type TestAcceleratorFilterRequest struct {
	Filters    []*AcceleratorFilter `json:"filters,omitempty"`
	Repository *string              `json:"repository,omitempty"`
}
