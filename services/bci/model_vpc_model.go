package bci

type VpcModel struct {
	VpcId       *string `json:"vpcId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Cidr        *string `json:"cidr,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	Description *string `json:"description,omitempty"`
	IsDefault   *bool   `json:"isDefault,omitempty"`
}
