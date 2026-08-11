package pfs

type CreateL3MountTargetRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	VpcId      *string `json:"vpcId,omitempty"`
	SubnetId   *string `json:"subnetId,omitempty"`
}
