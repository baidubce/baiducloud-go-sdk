package vpc

type CreateGatewayLimitRulesRequest struct {
	ClientToken    *string `json:"-"`
	IpVersion      *string `json:"ipVersion,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	ServiceType    *string `json:"serviceType,omitempty"`
	SubServiceType *string `json:"subServiceType,omitempty"`
	PeerRegion     *string `json:"peerRegion,omitempty"`
	ResourceId     *string `json:"resourceId,omitempty"`
	Direction      *string `json:"direction,omitempty"`
	Cidr           *string `json:"cidr,omitempty"`
	Bandwidth      *int32  `json:"bandwidth,omitempty"`
}
