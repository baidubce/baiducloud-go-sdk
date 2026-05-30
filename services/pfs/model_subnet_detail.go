package pfs

type SubnetDetail struct {
	Cidr         *string `json:"cidr,omitempty"`
	PhysicalZone *string `json:"physicalZone,omitempty"`
	SubnetId     *string `json:"subnetId,omitempty"`
	ZoneName     *string `json:"zoneName,omitempty"`
}
