package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEhcClusterResponse struct {
	bce.BaseResponse
	EhcClusterId *string `json:"ehcClusterId,omitempty"`
}
