package bci

type SubnetModel struct {
	SubnetId    *string `json:"subnetId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Cidr        *string `json:"cidr,omitempty"`
	VpcId       *string `json:"vpcId,omitempty"`
	SubnetType  *string `json:"subnetType,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
}
