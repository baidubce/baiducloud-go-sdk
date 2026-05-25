package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheSpecifiedIpAddressFamilyResponse struct {
	bce.BaseResponse
	IpGroupId       *string                       `json:"ipGroupId,omitempty"`
	Name            *string                       `json:"name,omitempty"`
	Description     *string                       `json:"description,omitempty"`
	IpVersion       *string                       `json:"ipVersion,omitempty"`
	IpSetIds        []*string                     `json:"ipSetIds,omitempty"`
	BindedInstances []*IpCollectionBindedInstance `json:"bindedInstances,omitempty"`
}
