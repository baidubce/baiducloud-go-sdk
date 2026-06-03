package ccr

type ToggleAcceleratorFilterRequest struct {
	InstanceId *string `json:"-"`
	PolicyId   *string `json:"-"`
	Enabled    *string `json:"-"`
}
