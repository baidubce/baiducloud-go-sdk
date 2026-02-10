package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAReservedNetworkSegmentResponse struct {
	bce.BaseResponse
	IpReserveId *string `json:"ipReserveId,omitempty"`
}
