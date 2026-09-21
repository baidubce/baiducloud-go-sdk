package vdb

type UpdateInstanceNameUsingPOSTRequest struct {
	InstanceId   *string `json:"-"`
	EngineType   *string `json:"-"`
	InstanceName *string `json:"instanceName,omitempty"`
}
