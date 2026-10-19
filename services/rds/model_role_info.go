package rds

type RoleInfo struct {
	Id       *string `json:"id,omitempty"`
	Azone    *string `json:"azone,omitempty"`
	SubnetId *string `json:"subnetId,omitempty"`
	Name     *string `json:"name,omitempty"`
	VpcCidr  *string `json:"vpcCidr,omitempty"`
	ShortId  *string `json:"shortId,omitempty"`
}
