package blb

type AssociateBlbModel struct {
	BlbId        *string `json:"blbId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Status       *string `json:"status,omitempty"`
	BlbType      *string `json:"blbType,omitempty"`
	PublicIp     *string `json:"publicIp,omitempty"`
	EipRouteType *string `json:"eipRouteType,omitempty"`
	Bandwidth    *int32  `json:"bandwidth,omitempty"`
	Address      *string `json:"address,omitempty"`
	Ipv6         *string `json:"ipv6,omitempty"`
	VpcId        *string `json:"vpcId,omitempty"`
	SubnetId     *string `json:"subnetId,omitempty"`
}
