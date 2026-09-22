package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetParameterListResponse struct {
	bce.BaseResponse
	Parameters []*Parameter `json:"parameters,omitempty"`
}
