package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListTbspAreaBlockingResponse struct {
	bce.BaseResponse
	AreaBlockingList []*TbspAreaBlockingModel `json:"areaBlockingList,omitempty"`
	Id               *string                  `json:"id,omitempty"`
}
