package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SearchVpnTunnelResponse struct {
	bce.BaseResponse
	VpnConns []*VpnConn `json:"vpnConns,omitempty"`
}
