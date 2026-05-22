package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeBlbResponse struct {
	bce.BaseResponse
	BlbId                        *string          `json:"blbId,omitempty"`
	Status                       *string          `json:"status,omitempty"`
	Name                         *string          `json:"name,omitempty"`
	Desc                         *string          `json:"desc,omitempty"`
	Address                      *string          `json:"address,omitempty"`
	PublicIp                     *string          `json:"publicIp,omitempty"`
	Cidr                         *string          `json:"cidr,omitempty"`
	VpcName                      *string          `json:"vpcName,omitempty"`
	VpcId                        *string          `json:"vpcId,omitempty"`
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
}
