package et

type UnrelatedDedicatedLineChannelRequest struct {
	EtId           *string `json:"-"`
	EtChannelId    *string `json:"-"`
	ClientToken    *string `json:"-"`
	ExtraChannelId *string `json:"extraChannelId,omitempty"`
}
