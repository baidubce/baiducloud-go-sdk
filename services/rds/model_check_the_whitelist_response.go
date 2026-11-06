package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CheckTheWhitelistResponse struct {
	bce.BaseResponse
	Etag            *string   `json:"etag,omitempty"`
	SecurityIps     []*string `json:"securityIps,omitempty"`
	SecurityIpGroup *string   `json:"securityIpGroup,omitempty"`
	Tag             *string   `json:"tag,omitempty"`
}
