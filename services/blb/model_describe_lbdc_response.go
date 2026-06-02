package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLbdcResponse struct {
	bce.BaseResponse
	Id                 *string     `json:"id,omitempty"`
	Name               *string     `json:"name,omitempty"`
	BlbType            *string     `json:"type,omitempty"`
	Status             *string     `json:"status,omitempty"`
	CcuCount           *int32      `json:"ccuCount,omitempty"`
	CreateTime         *string     `json:"createTime,omitempty"`
	ExpireTime         *string     `json:"expireTime,omitempty"`
	TotalConnectCount  *int64      `json:"totalConnectCount,omitempty"`
	NewConnectCps      *int64      `json:"newConnectCps,omitempty"`
	NetworkInBps       *int64      `json:"networkInBps,omitempty"`
	NetworkOutBps      *int64      `json:"networkOutBps,omitempty"`
	HttpsQps           *int64      `json:"httpsQps,omitempty"`
	HttpQps            *int64      `json:"httpQps,omitempty"`
	HttpNewConnectCps  *int64      `json:"httpNewConnectCps,omitempty"`
	HttpsNewConnectCps *int64      `json:"httpsNewConnectCps,omitempty"`
	SslNewConnectCps   *int64      `json:"sslNewConnectCps,omitempty"`
	Tags               []*TagModel `json:"tags,omitempty"`
}
