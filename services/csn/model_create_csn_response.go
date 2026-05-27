package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateCsnResponse struct {
	bce.BaseResponse
	CsnId *string `json:"csnId,omitempty"`
}
