package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type KeypairDetailResponse struct {
	bce.BaseResponse
	Keypair *KeypairModel `json:"keypair,omitempty"`
}
