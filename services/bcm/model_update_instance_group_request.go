package bcm

type UpdateInstanceGroupRequest struct {
	Id           *string                  `json:"id,omitempty"`
	Scope        *string                  `json:"scope,omitempty"`
	ResourceType *string                  `json:"resourceType,omitempty"`
	Name         *string                  `json:"name,omitempty"`
	Instances    []*InstanceGroupInstance `json:"instances,omitempty"`
}
