package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAPeerToPeerConnectionResponse struct {
	bce.BaseResponse
	PeerConnId *string `json:"peerConnId,omitempty"`
}
