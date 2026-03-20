package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateVpnResponse struct {
	bce.BaseResponse
	VpnId *string `json:"vpnId,omitempty"`
}
