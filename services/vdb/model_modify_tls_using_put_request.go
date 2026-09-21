package vdb

type ModifyTLSUsingPUTRequest struct {
	EngineType *string `json:"-"`
	Action     *string `json:"action,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
}
