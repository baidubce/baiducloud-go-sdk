package bcm

type AlarmTarget struct {
	BcmType             *string           `json:"type,omitempty"`
	Instances           []*TargetInstance `json:"instances,omitempty"`
	Region              *string           `json:"region,omitempty"`
	Tags                []*Dimension      `json:"tags,omitempty"`
	InstanceGroups      []*string         `json:"instanceGroups,omitempty"`
	IncludingDimensions []*string         `json:"includingDimensions,omitempty"`
	ExcludingDimensions []*string         `json:"excludingDimensions,omitempty"`
}
