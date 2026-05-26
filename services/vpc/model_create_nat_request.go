package vpc

type CreateNatRequest struct {
	ClientToken     *string        `json:"-"`
	Name            *string        `json:"name,omitempty"`
	VpcId           *string        `json:"vpcId,omitempty"`
	CuNum           *int32         `json:"cuNum,omitempty"`
	IpVersion       *string        `json:"ipVersion,omitempty"`
	BindEips        []*string      `json:"bindEips,omitempty"`
	Billing         *Billing       `json:"billing,omitempty"`
	SessionConfig   *SessionConfig `json:"sessionConfig,omitempty"`
	Tags            []*TagModel    `json:"tags,omitempty"`
	ResourceGroupId *string        `json:"resourceGroupId,omitempty"`
	DeleteProtect   *bool          `json:"deleteProtect,omitempty"`
}
