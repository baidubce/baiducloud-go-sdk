

package vpc


import "github.com/baidubce/baiducloud-go-sdk/bce"



type CreateVpnTunnelResponse struct {
	bce.BaseResponse
	VpnConnId *string `json:"vpnConnId,omitempty"`
}

