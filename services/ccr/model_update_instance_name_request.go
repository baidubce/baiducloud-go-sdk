package ccr

type UpdateInstanceNameRequest struct {
	InstanceId *string `json:"-"`
	Name       *string `json:"name,omitempty"`
}
