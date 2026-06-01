package et

type UpdateDedicatedChannelBfdRequest struct {
	EtId             *string `json:"-"`
	EtChannelId      *string `json:"-"`
	ClientToken      *string `json:"-"`
	SendInterval     *int32  `json:"sendInterval,omitempty"`
	ReceivInterval   *int32  `json:"receivInterval,omitempty"`
	DetectMultiplier *int32  `json:"detectMultiplier,omitempty"`
}
