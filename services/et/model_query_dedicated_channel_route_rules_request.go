package et

type QueryDedicatedChannelRouteRulesRequest struct {
	EtId        *string `json:"-"`
	EtChannelId *string `json:"-"`
	Marker      *string `json:"-"`
	MaxKeys     *int32  `json:"-"`
	DestAddress *string `json:"-"`
}
