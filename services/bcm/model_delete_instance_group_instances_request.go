package bcm

type DeleteInstanceGroupInstancesRequest struct {
	Id        *string                  `json:"id,omitempty"`
	Instances []*InstanceGroupInstance `json:"instances,omitempty"`
}
