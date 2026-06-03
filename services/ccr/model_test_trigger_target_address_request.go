package ccr

type TestTriggerTargetAddressRequest struct {
	InstanceId *string            `json:"-"`
	Address    *string            `json:"address,omitempty"`
	Headers    *map[string]string `json:"headers,omitempty"`
}
