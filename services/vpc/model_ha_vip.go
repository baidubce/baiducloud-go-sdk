package vpc

type HaVip struct {
	HaVipId          *string `json:"haVipId,omitempty"`
	Name             *string `json:"name,omitempty"`
	Description      *string `json:"description,omitempty"`
	VpcId            *string `json:"vpcId,omitempty"`
	SubnetId         *string `json:"subnetId,omitempty"`
	Status           *string `json:"status,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
	CreatedTime      *string `json:"createdTime,omitempty"`
}
