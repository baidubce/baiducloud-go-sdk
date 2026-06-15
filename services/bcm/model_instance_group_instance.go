package bcm

type InstanceGroupInstance struct {
	Dimensions []*Dimension `json:"dimensions,omitempty"`
	Region     *string      `json:"region,omitempty"`
}
