package rds

type UpdateInstanceSynchronizationModeRequest struct {
	InstanceId *string `json:"-"`
	SyncMode   *string `json:"syncMode,omitempty"`
}
