package as

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ScalingDownV2Response struct {
	bce.BaseResponse
	Success *bool     `json:"success,omitempty"`
	Msg     *string   `json:"msg,omitempty"`
	Result  *DagModel `json:"result,omitempty"`
}
