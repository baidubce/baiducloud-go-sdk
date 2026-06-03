package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateInstanceNameResponse struct {
	bce.BaseResponse
	Id           *string     `json:"id,omitempty"`
	Name         *string     `json:"name,omitempty"`
	InstanceType *string     `json:"instanceType,omitempty"`
	PublicURL    *string     `json:"publicURL,omitempty"`
	Region       *string     `json:"region,omitempty"`
	Status       *string     `json:"status,omitempty"`
	CreateTime   *string     `json:"createTime,omitempty"`
	Tags         *LogicalTag `json:"tags,omitempty"`
}
