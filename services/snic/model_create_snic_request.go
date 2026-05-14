package snic

type CreateSnicRequest struct {
	ClientToken     *string     `json:"-"`
	VpcId           *string     `json:"vpcId,omitempty"`
	Name            *string     `json:"name,omitempty"`
	SubnetId        *string     `json:"subnetId,omitempty"`
	Service         *string     `json:"service,omitempty"`
	Description     *string     `json:"description,omitempty"`
	IpAddress       *string     `json:"ipAddress,omitempty"`
	Bandwidth       *int32      `json:"bandwidth,omitempty"`
	Billing         *Billing    `json:"billing,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	ResourceGroupId *string     `json:"resourceGroupId,omitempty"`
}
