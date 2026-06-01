package blb

type CreateServiceRequest struct {
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty"`
	AuthList    []*Auth `json:"authList,omitempty"`
}
