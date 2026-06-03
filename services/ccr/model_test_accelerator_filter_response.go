package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TestAcceleratorFilterResponse struct {
	bce.BaseResponse
	Matched *bool `json:"matched,omitempty"`
}
