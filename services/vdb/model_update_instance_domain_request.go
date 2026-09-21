package vdb

type UpdateInstanceDomainRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	Domain     *string `json:"domain,omitempty"`
}
