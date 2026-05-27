package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryCsnBpDetailResponse struct {
	bce.BaseResponse
	CsnBpId         *string     `json:"csnBpId,omitempty"`
	Name            *string     `json:"name,omitempty"`
	Bandwidth       *string     `json:"bandwidth,omitempty"`
	UsedBandwidth   *string     `json:"usedBandwidth,omitempty"`
	CsnId           *string     `json:"csnId,omitempty"`
	InterworkType   *string     `json:"interworkType,omitempty"`
	InterworkRegion *string     `json:"interworkRegion,omitempty"`
	Status          *string     `json:"status,omitempty"`
	PaymentTiming   *string     `json:"paymentTiming,omitempty"`
	ExpireTime      *string     `json:"expireTime,omitempty"`
	CreatedTime     *string     `json:"createdTime,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
}
