package et

type QueryDedicatedChannelRequest struct {
	EtId        *string `json:"-"`
	ClientToken *string `json:"-"`
	EtChannelId *string `json:"-"`
}
