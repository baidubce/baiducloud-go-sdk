package blb

type Service struct {
	ServiceId     *string            `json:"serviceId,omitempty"`
	Name          *string            `json:"name,omitempty"`
	Description   *string            `json:"description,omitempty"`
	ServiceName   *string            `json:"serviceName,omitempty"`
	BindType      *string            `json:"bindType,omitempty"`
	InstanceId    *string            `json:"instanceId,omitempty"`
	Status        *string            `json:"status,omitempty"`
	Service       *string            `json:"service,omitempty"`
	CreateTime    *string            `json:"createTime,omitempty"`
	EndpointCount *int32             `json:"endpointCount,omitempty"`
	EndpointList  []*RelatedEndpoint `json:"endpointList,omitempty"`
	AuthList      []*Auth            `json:"authList,omitempty"`
}
