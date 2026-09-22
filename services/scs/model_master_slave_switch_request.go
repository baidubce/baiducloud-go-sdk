package scs

type MasterSlaveSwitchRequest struct {
	InstanceId *string                   `json:"-"`
	Shards     []*SwitchMasterSlaveShard `json:"shards,omitempty"`
}
