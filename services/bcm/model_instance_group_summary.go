package bcm

type InstanceGroupSummary struct {
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Scope         *string `json:"scope,omitempty"`
	ResourceType  *string `json:"resourceType,omitempty"`
	InstanceCount *int32  `json:"instanceCount,omitempty"`
}
