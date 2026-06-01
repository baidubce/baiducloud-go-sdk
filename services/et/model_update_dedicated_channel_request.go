package et

type UpdateDedicatedChannelRequest struct {
	EtId          *string `json:"-"`
	EtChannelId   *string `json:"-"`
	ClientToken   *string `json:"-"`
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	BgpRouteLimit *int32  `json:"bgpRouteLimit,omitempty"`
}
