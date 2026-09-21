package vdb

type Subnet struct {
	Az       *string `json:"az,omitempty"`
	Cidr     *string `json:"cidr,omitempty"`
	Name     *string `json:"name,omitempty"`
	ShortId  *string `json:"shortId,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`
}
