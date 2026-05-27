package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateProbeResponse struct {
	bce.BaseResponse
	ProbeId *string `json:"probeId,omitempty"`
}
