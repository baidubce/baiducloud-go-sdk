package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AccountListResponse struct {
	bce.BaseResponse
	Success *bool       `json:"success,omitempty"`
	Result  []*ListItem `json:"result,omitempty"`
}
