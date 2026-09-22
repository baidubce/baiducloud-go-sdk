package scs

type ModifyEntranceRequest struct {
	InstanceId *string `json:"-"`
	IsDefer    *bool   `json:"isDefer,omitempty"`
}
