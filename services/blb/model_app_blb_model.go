package blb

type AppBLBModel struct {
	BlbId                        *string     `json:"blbId,omitempty"`
	Name                         *string     `json:"name,omitempty"`
	Desc                         *string     `json:"desc,omitempty"`
	Address                      *string     `json:"address,omitempty"`
	Status                       *string     `json:"status,omitempty"`
	SubnetId                     *string     `json:"subnetId,omitempty"`
	VpcId                        *string     `json:"vpcId,omitempty"`
	PublicIp                     *string     `json:"publicIp,omitempty"`
	Tags                         []*TagModel `json:"tags,omitempty"`
	AllowDelete                  *bool       `json:"allowDelete,omitempty"`
	AllowModify                  *bool       `json:"allowModify,omitempty"`
	ModificationProtectionReason *string     `json:"modificationProtectionReason,omitempty"`
	EipRouteType                 *string     `json:"eipRouteType,omitempty"`
	PublicIpv6                   *string     `json:"publicIpv6,omitempty"`
	EipV6RouteType               *string     `json:"eipV6RouteType,omitempty"`
	PaymentTiming                *string     `json:"paymentTiming,omitempty"`
	BillingMethod                *string     `json:"billingMethod,omitempty"`
}
