package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ImportKeypairResponse struct {
	bce.BaseResponse
	Keypair *KeypairModel `json:"keypair,omitempty"`
}
