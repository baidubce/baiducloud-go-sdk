package vpc

type NAT struct {
	Id            *string        `json:"id,omitempty"`
	Name          *string        `json:"name,omitempty"`
	NatType       *string        `json:"natType,omitempty"`
	VpcId         *string        `json:"vpcId,omitempty"`
	CuNum         *int32         `json:"cuNum,omitempty"`
	Spec          *string        `json:"spec,omitempty"`
	BindEips      []*string      `json:"bindEips,omitempty"`
	Status        *string        `json:"status,omitempty"`
	PaymentTiming *string        `json:"paymentTiming,omitempty"`
	SessionConfig *SessionConfig `json:"sessionConfig,omitempty"`
	IpVersion     *string        `json:"ipVersion,omitempty"`
	ExpiredTime   *string        `json:"expiredTime,omitempty"`
	CreateTime    *string        `json:"createTime,omitempty"`
}
