package et

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryDedicatedChannelResponse struct {
	bce.BaseResponse
	EtChannels []*EtChannel `json:"etChannels,omitempty"`
}
