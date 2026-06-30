package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdateSubUserIdpStatusResponse struct {
	bce.BaseResponse
	Idp *Idp `json:"idp,omitempty"`
}
