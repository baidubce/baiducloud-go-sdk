package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryIpWhitelistResponse struct {
	bce.BaseResponse
	SecurityIps []*string `json:"securityIps,omitempty"`
}
