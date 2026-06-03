package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CheckBeforeCreateInstanceResponse struct {
	bce.BaseResponse
	Pass    *bool    `json:"pass,omitempty"`
	ErrInfo *ErrInfo `json:"errInfo,omitempty"`
}
