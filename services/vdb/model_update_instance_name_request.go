package vdb

type UpdateInstanceNameRequest struct {
	InstanceId   *string `json:"-"`
	EngineType   *string `json:"-"`
	InstanceName *string `json:"instanceName,omitempty"`
}
