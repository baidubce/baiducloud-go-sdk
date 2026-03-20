package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateUserGatewayResponse struct {
	bce.BaseResponse
	CgwId *string `json:"cgwId,omitempty"`
}
