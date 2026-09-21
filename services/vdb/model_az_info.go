package vdb

type AzInfo struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	Count            *int32  `json:"count,omitempty"`
	SubnetId         *string `json:"subnetId,omitempty"`
}
