package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetDnsResolverDetailResponse struct {
	bce.BaseResponse
	Id              *string    `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Status          *string    `json:"status,omitempty"`
	Description     *string    `json:"description,omitempty"`
	VpcId           *string    `json:"vpcId,omitempty"`
	VpcRegion       *string    `json:"vpcRegion,omitempty"`
	PrivatezoneType *string    `json:"type,omitempty"`
	IpModels        []*IpModel `json:"ipModels,omitempty"`
	CreateTime      *string    `json:"createTime,omitempty"`
	UpdateTime      *string    `json:"updateTime,omitempty"`
}
