package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheSpecifiedIpAddressGroupResponse struct {
	bce.BaseResponse
	IpSetId         *string                       `json:"ipSetId,omitempty"`
	Name            *string                       `json:"name,omitempty"`
	Description     *string                       `json:"description,omitempty"`
	IpVersion       *string                       `json:"ipVersion,omitempty"`
	IpAddressInfo   []*TemplateIpAddressInfo      `json:"ipAddressInfo,omitempty"`
	BindedInstances []*IpCollectionBindedInstance `json:"bindedInstances,omitempty"`
}
