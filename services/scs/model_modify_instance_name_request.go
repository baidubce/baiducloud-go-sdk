package scs

type ModifyInstanceNameRequest struct {
	InstanceId   *string `json:"-"`
	InstanceName *string `json:"instanceName,omitempty"`
}
