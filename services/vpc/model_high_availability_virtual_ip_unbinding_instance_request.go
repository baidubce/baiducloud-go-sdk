package vpc

type HighAvailabilityVirtualIpUnbindingInstanceRequest struct {
	HaVipId      *string   `json:"-"`
	ClientToken  *string   `json:"-"`
	InstanceIds  []*string `json:"instanceIds,omitempty"`
	InstanceType *string   `json:"instanceType,omitempty"`
}
