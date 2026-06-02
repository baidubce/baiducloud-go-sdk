package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLbdcBlbResponse struct {
	bce.BaseResponse
	BlbList []*AssociateBlbModel `json:"blbList,omitempty"`
}
