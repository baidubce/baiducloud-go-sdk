package bcc

type AutoRenewVolumeClusterRequest struct {
	ClusterId     *string `json:"clusterId,omitempty"`
	RenewTimeUnit *string `json:"renewTimeUnit,omitempty"`
	RenewTime     *int32  `json:"renewTime,omitempty"`
}
