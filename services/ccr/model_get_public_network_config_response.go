package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPublicNetworkConfigResponse struct {
	bce.BaseResponse
	Domain    *string      `json:"domain,omitempty"`
	Status    *string      `json:"status,omitempty"`
	Whitelist []*Whitelist `json:"whitelist,omitempty"`
}
