package rds

type InstanceSubnet struct {
	Name     *string `json:"name,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`
	Az       *string `json:"az,omitempty"`
	ZoneName *string `json:"zoneName,omitempty"`
	Cidr     *string `json:"cidr,omitempty"`
	ShortId  *string `json:"shortId,omitempty"`
}
