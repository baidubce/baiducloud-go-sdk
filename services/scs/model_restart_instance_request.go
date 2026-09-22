package scs

type RestartInstanceRequest struct {
	InstanceId *string `json:"-"`
	IsDefer    *bool   `json:"isDefer,omitempty"`
}
