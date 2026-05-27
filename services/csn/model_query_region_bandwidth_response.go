package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryRegionBandwidthResponse struct {
	bce.BaseResponse
	BpLimits []*CsnBpLimit `json:"bpLimits,omitempty"`
}
