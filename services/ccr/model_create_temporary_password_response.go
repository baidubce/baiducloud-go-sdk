package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateTemporaryPasswordResponse struct {
	bce.BaseResponse
	Password *string `json:"password,omitempty"`
}
