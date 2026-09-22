package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryMemoryScalingConfigResponse struct {
	bce.BaseResponse
	MemSpec *MemSpec `json:"memSpec,omitempty"`
}
