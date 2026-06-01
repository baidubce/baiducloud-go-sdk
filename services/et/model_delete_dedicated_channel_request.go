package et

type DeleteDedicatedChannelRequest struct {
	EtId        *string `json:"-"`
	EtChannelId *string `json:"-"`
	ClientToken *string `json:"-"`
}
