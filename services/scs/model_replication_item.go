package scs

type ReplicationItem struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	SubnetId         *string `json:"subnetId,omitempty"`
	IsMaster         *int32  `json:"isMaster,omitempty"`
	Weight           *int32  `json:"weight,omitempty"`
}
