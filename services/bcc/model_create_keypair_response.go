package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateKeypairResponse struct {
	bce.BaseResponse
	Keypair *KeypairModel `json:"keypair,omitempty"`
}
