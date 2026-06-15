package bcm

type CreateInstanceGroupRequest struct {
	Scope        *string                  `json:"scope,omitempty"`
	ResourceType *string                  `json:"resourceType,omitempty"`
	Name         *string                  `json:"name,omitempty"`
	Instances    []*InstanceGroupInstance `json:"instances,omitempty"`
}
