package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreatePeerConnResponse struct {
	bce.BaseResponse
	PeerConnId *string `json:"peerConnId,omitempty"`
}
