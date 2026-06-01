package et

type CreateDedicatedChannelRouteRulesRequest struct {
	EtId        *string `json:"-"`
	EtChannelId *string `json:"-"`
	ClientToken *string `json:"-"`
	IpVersion   *int32  `json:"ipVersion,omitempty"`
	DestAddress *string `json:"destAddress,omitempty"`
	NexthopType *string `json:"nexthopType,omitempty"`
	NexthopId   *string `json:"nexthopId,omitempty"`
	Description *string `json:"description,omitempty"`
}
