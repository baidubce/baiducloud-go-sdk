package ccr

type DeleteVpcLinkRequest struct {
	InstanceId *string `json:"-"`
	VpcID      *string `json:"vpcID,omitempty"`
	SubnetID   *string `json:"subnetID,omitempty"`
}
