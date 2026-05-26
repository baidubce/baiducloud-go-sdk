package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetNatResponse struct {
	bce.BaseResponse
	Id            *string        `json:"id,omitempty"`
	Name          *string        `json:"name,omitempty"`
	NatType       *string        `json:"natType,omitempty"`
	VpcId         *string        `json:"vpcId,omitempty"`
	Spec          *string        `json:"spec,omitempty"`
	CuNum         *int32         `json:"cuNum,omitempty"`
	BindEips      []*string      `json:"bindEips,omitempty"`
	Status        *string        `json:"status,omitempty"`
	IpVersion     *string        `json:"ipVersion,omitempty"`
	SessionConfig *SessionConfig `json:"sessionConfig,omitempty"`
	PaymentTiming *string        `json:"paymentTiming,omitempty"`
	ExpiredTime   *string        `json:"expiredTime,omitempty"`
	CreateTime    *string        `json:"createTime,omitempty"`
	Tags          []*TagModel    `json:"tags,omitempty"`
	DeleteProtect *bool          `json:"deleteProtect,omitempty"`
}
