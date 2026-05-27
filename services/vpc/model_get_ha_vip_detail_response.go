package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetHaVipDetailResponse struct {
	bce.BaseResponse
	HaVipId          *string                `json:"haVipId,omitempty"`
	Name             *string                `json:"name,omitempty"`
	Description      *string                `json:"description,omitempty"`
	VpcId            *string                `json:"vpcId,omitempty"`
	SubnetId         *string                `json:"subnetId,omitempty"`
	Status           *string                `json:"status,omitempty"`
	PrivateIpAddress *string                `json:"privateIpAddress,omitempty"`
	PublicIpAddress  *string                `json:"publicIpAddress,omitempty"`
	BindedInstances  []*HaVipBindedInstance `json:"bindedInstances,omitempty"`
	CreatedTime      *string                `json:"createdTime,omitempty"`
}
