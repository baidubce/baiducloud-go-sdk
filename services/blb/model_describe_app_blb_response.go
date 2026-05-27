package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAppBlbResponse struct {
	bce.BaseResponse
	BlbId                        *string          `json:"blbId,omitempty"`
	Name                         *string          `json:"name,omitempty"`
	Status                       *string          `json:"status,omitempty"`
	Desc                         *string          `json:"desc,omitempty"`
	Address                      *string          `json:"address,omitempty"`
	PublicIp                     *string          `json:"publicIp,omitempty"`
	Cidr                         *string          `json:"cidr,omitempty"`
	VpcId                        *string          `json:"vpcId,omitempty"`
	VpcName                      *string          `json:"vpcName,omitempty"`
	SubnetCider                  *string          `json:"subnetCider,omitempty"`
	SubnetName                   *string          `json:"subnetName,omitempty"`
	CreateTime                   *string          `json:"createTime,omitempty"`
	ReleaseTime                  *string          `json:"releaseTime,omitempty"`
	Listener                     []*ListenerModel `json:"listener,omitempty"`
	Tags                         []*TagModel      `json:"tags,omitempty"`
	AllowDelete                  *bool            `json:"allowDelete,omitempty"`
	AllowModify                  *bool            `json:"allowModify,omitempty"`
	ModificationProtectionReason *string          `json:"modificationProtectionReason,omitempty"`
	PaymentTiming                *string          `json:"paymentTiming,omitempty"`
	BillingMethod                *string          `json:"billingMethod,omitempty"`
	PerformanceLevel             *string          `json:"performanceLevel,omitempty"`
	ExpireTime                   *string          `json:"expireTime,omitempty"`
	EipRouteType                 *string          `json:"eipRouteType,omitempty"`
	PublicIpv6                   *string          `json:"publicIpv6,omitempty"`
	EipV6RouteType               *string          `json:"eipV6RouteType,omitempty"`
	Ipv6                         *string          `json:"ipv6,omitempty"`
	SupportAcl                   *bool            `json:"supportAcl,omitempty"`
}
