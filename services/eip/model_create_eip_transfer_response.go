package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEipTransferResponse struct {
	bce.BaseResponse
	Transfers []*TransferInfo `json:"transfers,omitempty"`
}
