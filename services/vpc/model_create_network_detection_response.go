package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateNetworkDetectionResponse struct {
	bce.BaseResponse
	ProbeId *string `json:"probeId,omitempty"`
}
