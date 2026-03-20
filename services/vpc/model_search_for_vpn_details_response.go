package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SearchForVpnDetailsResponse struct {
	bce.BaseResponse
	VpnId           *string       `json:"vpnId,omitempty"`
	VpnName         *string       `json:"vpnName,omitempty"`
	CreateTime      *string       `json:"createTime,omitempty"`
	Description     *string       `json:"description,omitempty"`
	Type            *string       `json:"type,omitempty"`
	Status          *string       `json:"status,omitempty"`
	ExpiredTime     *string       `json:"expiredTime,omitempty"`
	PaymentTiming   *string       `json:"paymentTiming,omitempty"`
	Eip             *string       `json:"eip,omitempty"`
	BandwidthInMbps *int32        `json:"bandwidthInMbps,omitempty"`
	VpcId           *string       `json:"vpcId,omitempty"`
	VpnConnNum      *int32        `json:"vpnConnNum,omitempty"`
	MaxConnection   *int32        `json:"maxConnection,omitempty"`
	VpnConns        []*VpnConn    `json:"vpnConns,omitempty"`
	SslVpnServer    *SslVpnServer `json:"sslVpnServer,omitempty"`
	Tags            []*TagModel   `json:"tags,omitempty"`
	DeleteProtect   *bool         `json:"deleteProtect,omitempty"`
}
