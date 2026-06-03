package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListVpcLinksResponse struct {
	bce.BaseResponse
	Domain *string         `json:"domain,omitempty"`
	Items  []*Privatelinks `json:"items,omitempty"`
}
