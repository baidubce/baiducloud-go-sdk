package scs

type Subnet struct {
	Name     *string `json:"name,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`
	ZoneName *string `json:"zoneName,omitempty"`
	Cidr     *string `json:"cidr,omitempty"`
	VpcId    *string `json:"vpcId,omitempty"`
}
