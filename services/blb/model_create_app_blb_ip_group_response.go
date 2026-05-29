package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAppBlbIpGroupResponse struct {
	bce.BaseResponse
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Desc *string `json:"desc,omitempty"`
}
