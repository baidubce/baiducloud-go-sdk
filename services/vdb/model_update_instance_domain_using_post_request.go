package vdb

type UpdateInstanceDomainUsingPOSTRequest struct {
	InstanceId *string `json:"-"`
	EngineType *string `json:"-"`
	Domain     *string `json:"domain,omitempty"`
}
