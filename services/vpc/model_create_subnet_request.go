package vpc

type CreateSubnetRequest struct {
	ClientToken      *string     `json:"-"`
	Name             *string     `json:"name,omitempty"`
	EnableIpv6       *bool       `json:"enableIpv6,omitempty"`
	ZoneName         *string     `json:"zoneName,omitempty"`
	Cidr             *string     `json:"cidr,omitempty"`
	VpcId            *string     `json:"vpcId,omitempty"`
	VpcSecondaryCidr *string     `json:"vpcSecondaryCidr,omitempty"`
	SubnetType       *string     `json:"subnetType,omitempty"`
	Description      *string     `json:"description,omitempty"`
	Tags             []*TagModel `json:"tags,omitempty"`
}
