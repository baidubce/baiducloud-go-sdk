package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryRegionBandwidthByCsnResponse struct {
	bce.BaseResponse
	BpLimits []*CsnBpLimit `json:"bpLimits,omitempty"`
}
