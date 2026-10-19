package rds

type UpdateInstanceNameRequest struct {
	InstanceId   *string `json:"-"`
	InstanceName *string `json:"instanceName,omitempty"`
}
