package privatezone

type DnsResolver struct {
	Id              *string    `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Status          *string    `json:"status,omitempty"`
	Description     *string    `json:"description,omitempty"`
	VpcId           *string    `json:"vpcId,omitempty"`
	PrivatezoneType *string    `json:"type,omitempty"`
	VpcRegion       *string    `json:"vpcRegion,omitempty"`
	IpModels        []*IpModel `json:"ipModels,omitempty"`
	CreateTime      *string    `json:"createTime,omitempty"`
	UpdateTime      *string    `json:"updateTime,omitempty"`
}
