package scs

type SubnetInfo struct {
	SubnetId *string `json:"subnetId,omitempty"`
	Name     *string `json:"name,omitempty"`
	Cidr     *string `json:"cidr,omitempty"`
	Az       *string `json:"az,omitempty"`
}
