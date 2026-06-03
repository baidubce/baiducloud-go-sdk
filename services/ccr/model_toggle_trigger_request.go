package ccr

type ToggleTriggerRequest struct {
	InstanceId *string `json:"-"`
	PolicyId   *string `json:"-"`
	Enabled    *string `json:"-"`
}
