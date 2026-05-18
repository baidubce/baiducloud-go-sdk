package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetEipGroupResponse struct {
	bce.BaseResponse
	Id                        *string     `json:"id,omitempty"`
	Name                      *string     `json:"name,omitempty"`
	Eips                      []*EipModel `json:"eips,omitempty"`
	Eipv6s                    []*EipModel `json:"eipv6s,omitempty"`
	Status                    *string     `json:"status,omitempty"`
	BandwidthInMbps           *int32      `json:"bandwidthInMbps,omitempty"`
	DefaultDomesticBandwidth  *int32      `json:"defaultDomesticBandwidth,omitempty"`
	BwShortId                 *string     `json:"bwShortId,omitempty"`
	BwBandwidthInMbps         *int32      `json:"bwBandwidthInMbps,omitempty"`
	DomesticBwShortId         *string     `json:"domesticBwShortId,omitempty"`
	DomesticBwBandwidthInMbps *int32      `json:"domesticBwBandwidthInMbps,omitempty"`
	PaymentTiming             *string     `json:"paymentTiming,omitempty"`
	BillingMethod             *string     `json:"billingMethod,omitempty"`
	CreateTime                *string     `json:"createTime,omitempty"`
	ExpireTime                *string     `json:"expireTime,omitempty"`
	Region                    *string     `json:"region,omitempty"`
	RouteType                 *string     `json:"routeType,omitempty"`
	Tags                      []*TagModel `json:"tags,omitempty"`
}
