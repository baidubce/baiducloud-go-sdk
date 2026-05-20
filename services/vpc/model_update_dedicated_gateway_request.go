package vpc

type UpdateDedicatedGatewayRequest struct {
	EtGatewayId    *string   `json:"-"`
	ClientToken    *string   `json:"-"`
	Name           *string   `json:"name,omitempty"`
	Speed          *int32    `json:"speed,omitempty"`
	Description    *string   `json:"description,omitempty"`
	LocalCidrs     []*string `json:"localCidrs,omitempty"`
	EnableIpv6     *int32    `json:"enableIpv6,omitempty"`
	Ipv6LocalCidrs []*string `json:"ipv6LocalCidrs,omitempty"`
}
