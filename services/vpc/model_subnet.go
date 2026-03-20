package vpc

type Subnet struct {
	SubnetId              *string     `json:"subnetId,omitempty"`
	Name                  *string     `json:"name,omitempty"`
	ZoneName              *string     `json:"zoneName,omitempty"`
	Cidr                  *string     `json:"cidr,omitempty"`
	Ipv6Cidr              *string     `json:"ipv6Cidr,omitempty"`
	VpcId                 *string     `json:"vpcId,omitempty"`
	SubnetType            *string     `json:"subnetType,omitempty"`
	Description           *string     `json:"description,omitempty"`
	AvailableIp           *int32      `json:"availableIp,omitempty"`
	AvailableUnreservedIp *int32      `json:"availableUnreservedIp,omitempty"`
	CreatedTime           *string     `json:"createdTime,omitempty"`
	Tags                  []*TagModel `json:"tags,omitempty"`
}
