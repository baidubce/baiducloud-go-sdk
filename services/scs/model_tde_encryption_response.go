package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TdeEncryptionResponse struct {
	bce.BaseResponse
	OperateStatus *string `json:"operateStatus,omitempty"`
	ErrorMessage  *string `json:"errorMessage,omitempty"`
}
