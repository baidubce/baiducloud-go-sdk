package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateCsnBpResponse struct {
	bce.BaseResponse
	CsnBpId *string `json:"csnBpId,omitempty"`
}
