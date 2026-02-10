package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListTbspProtocolBlockingResponse struct {
	bce.BaseResponse
	ProtocolBlockingList []*TbspProtocolBlockingModel `json:"protocolBlockingList,omitempty"`
	Id                   *string                      `json:"id,omitempty"`
}
