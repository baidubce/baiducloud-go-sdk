package ccr

type AddVpcLinkRequest struct {
	InstanceId     *string `json:"-"`
	VpcID          *string `json:"vpcID,omitempty"`
	SubnetID       *string `json:"subnetID,omitempty"`
	IpType         *string `json:"ipType,omitempty"`
	IpAddress      *string `json:"ipAddress,omitempty"`
	AutoDnsResolve *bool   `json:"autoDnsResolve,omitempty"`
}
