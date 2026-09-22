package scs

type ReplicationMap struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	SubnetId         *string `json:"subnetId,omitempty"`
	IsMaster         *int32  `json:"isMaster,omitempty"`
}
